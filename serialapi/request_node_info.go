package serialapi

import (
	"errors"

	"github.com/davecgh/go-spew/spew"
	"go.uber.org/zap"

	"github.com/opendoor-labs/gozw/frame"
	"github.com/opendoor-labs/gozw/protocol"
	"github.com/opendoor-labs/gozw/session"
)

// RequestNodeInfo will request info for a node.
func (s *Layer) RequestNodeInfo(nodeID byte) (*NodeInfoFrame, error) {
	var nodeInfo NodeInfoFrame

	done := make(chan *frame.Frame)

	request := &session.Request{
		FunctionID: protocol.FnRequestNodeInfo,
		Payload:    []byte{nodeID},
		HasReturn:  true,
		ReturnCallback: func(err error, ret *frame.Frame) bool {
			done <- ret
			if err != nil {
				s.l.Error("Request Node Failed:", zap.Error(err))
			}
			return false
		},
	}

	s.sessionLayer.MakeRequest(request)
	ret := <-done

	if ret == nil {
		return nil, errors.New("Error requesting node information frame")
	}

	status := ret.Payload[1]

	if status == 0 {
		return nil, errors.New("Failed putting node info request in transmit queue")
	}

	spew.Dump(ret)

	return &nodeInfo, nil
}

//  NodeInfoFrame contains a node info frame.
type NodeInfoFrame struct {
}
