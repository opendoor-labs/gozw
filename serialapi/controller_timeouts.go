package serialapi

import (
	"errors"
	"github.com/gozwave/gozw/frame"
	"github.com/gozwave/gozw/protocol"
	"github.com/gozwave/gozw/session"
)

type ControllerTimeouts struct {
	AckTimeout  byte
	ByteTimeout byte
}

const (
	// See INS12350-14 section 7.5
	ackTimeout  = 0x96 // 0x95 = 150 -- multiply by 10 to get 1500ms
	byteTimeout = 0x0F // 0x0F = 15 -- multiply by 10 to get 150ms

	errUnableToSetControllerTimeouts = "unable to set controller timeouts"
)

func (s *Layer) SetControllerTimeouts() (*ControllerTimeouts, error) {
	done := make(chan *frame.Frame)

	request := &session.Request{
		FunctionID: protocol.FnSerialAPISetTimeouts,
		HasReturn:  true,
		ReturnCallback: func(err error, ret *frame.Frame) bool {
			done <- ret
			return false
		},
		Payload: []byte{
			ackTimeout,
			byteTimeout,
		},
	}

	s.sessionLayer.MakeRequest(request)
	ret := <-done

	if ret == nil {
		return nil, errors.New(errUnableToSetControllerTimeouts)
	}

	val := ControllerTimeouts{
		AckTimeout:  ret.Payload[1],
		ByteTimeout: ret.Payload[2],
	}

	return &val, nil
}
