package serialapi

import (
	"time"

	"github.com/gozwave/gozw/frame"
	"github.com/gozwave/gozw/protocol"
	"github.com/gozwave/gozw/session"
)

// FactoryReset will clear the network configuration and Soft Reset the controller
// WARNING: This can (and often will) cause the device to get a new USB address,
// rendering the serial port's file descriptor invalid.
func (s *Layer) FactoryReset() {
	request := &session.Request{
		FunctionID:       protocol.FnSetDefault,
		HasReturn:        false,
		ReceivesCallback: true,
		Callback: func(f frame.Frame) {
			return
		},
	}

	s.sessionLayer.MakeRequest(request)

	time.Sleep(1500 * time.Millisecond)
}
