package frame

import (
	"context"
	"io"
	"time"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// ILayer is an interface for a frame layer.
type ILayer interface {
	Write(frame *Frame)
	GetOutputChannel() <-chan Frame
}

// Layer contains a frame layer.
type Layer struct {
	transportLayer io.ReadWriter

	frameParser      *Parser
	parserInput      chan<- byte
	parserOutput     <-chan *ParseEvent
	acks, naks, cans <-chan bool

	l *zap.Logger

	pendingWrites chan *Frame
	frameOutput   chan Frame

	ctx context.Context
}

// NewFrameLayer will return a new frame layer.
func NewFrameLayer(ctx context.Context, transportLayer io.ReadWriter, logger *zap.Logger) (*Layer, error) {
	if _, ok := transportLayer.(io.ByteReader); !ok {
		return nil, errors.New("transport layer does not implement io.ByteReader")
	}

	parserInput := make(chan byte)
	parserOutput := make(chan *ParseEvent, 1)
	acks := make(chan bool, 1)
	naks := make(chan bool, 1)
	cans := make(chan bool, 1)

	parser := NewParser(ctx, parserInput, parserOutput, acks, naks, cans, logger)

	frameLayer := Layer{
		transportLayer: transportLayer,
		frameParser:    parser,
		parserInput:    parserInput,
		parserOutput:   parserOutput,
		acks:           acks,
		naks:           naks,
		cans:           cans,
		l:              logger,
		pendingWrites:  make(chan *Frame),
		frameOutput:    make(chan Frame, 5),
		ctx:            ctx,
	}

	go frameLayer.bgWork()
	go func() {
		err := frameLayer.bgRead()
		if err != nil {
			logger.Fatal("background read job failed", zap.Error(err))
		}
	}()

	return &frameLayer, nil

}

func (l *Layer) bgWork() {

	for {
		select {
		case frameIn := <-l.parserOutput:
			l.l.Debug("parser output received")

			if frameIn.status == ParseOk {
				err := l.sendAck()
				if err != nil {
					l.l.Error("send ack failed", zap.Error(err))
				}
				l.l.Debug("received frame successfully, writing output")
				l.frameOutput <- frameIn.frame
			} else if frameIn.status == ParseNotOk {
				l.l.Warn("received frame, parse not ok")
				err := l.sendNak()
				if err != nil {
					l.l.Error("send-nak-failed", zap.Error(err))
				}
			} else if frameIn.status == ParseTimeout {
				// Does NOT transmit a ACK or NAK
				// Relevant reference believed to be INS12350-14 § 6.2.1
				l.l.Warn("receive frame timeout")
			}

		case <-l.acks:
			l.l.Debug("rx ack")
		case <-l.naks:
			l.l.Debug("rx nak")
		case <-l.cans:
			l.l.Debug("rx can")

		case frameToWrite := <-l.pendingWrites:
			l.l.Debug("frame received, writing to transport")
			// this method never returns an error, so ignore it
			buf, _ := frameToWrite.MarshalBinary()

			_, err := l.writeToTransport(buf)
			if err != nil {
				l.l.Error("write-transport-failed", zap.Error(err))
			}
			// TODO: this needs to time out

			// <-l.acks
			select {
			case <-l.acks:
				l.l.Debug("received ack")
			case <-time.After(1 * time.Second):
				l.l.Error("ack timed out")
			}
		case <-l.ctx.Done():
			l.l.Info("closing frame layer bg work")
			return
		}
	}
}

func (l *Layer) Write(frame *Frame) {
	go func() {
		l.pendingWrites <- frame
	}()
}

// GetOutputChannel will return the output channel.
func (l *Layer) GetOutputChannel() <-chan Frame {
	return l.frameOutput
}

func (l *Layer) bgRead() error {
	for {
		byt, err := l.transportLayer.(io.ByteReader).ReadByte()
		if err != nil {
			return err
		}
		l.parserInput <- byt
	}
}

func (l *Layer) writeToTransport(buf []byte) (int, error) {
	return l.transportLayer.Write(buf)
}

func (l *Layer) sendAck() error {
	_, err := l.transportLayer.Write([]byte{HeaderAck})
	return err
}

func (l *Layer) sendNak() error {
	_, err := l.transportLayer.Write([]byte{HeaderNak})
	return err
}
