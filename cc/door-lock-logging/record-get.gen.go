// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package doorlocklogging

import (
	"encoding/gob"
	"errors"
	"fmt"

	"github.com/gozwave/gozw/cc"
)

const CommandRecordGet cc.CommandID = 0x03

func init() {
	gob.Register(RecordGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x4C),
		Command:      cc.CommandID(0x03),
		Version:      1,
	}, NewRecordGet)
}

func NewRecordGet() cc.Command {
	return &RecordGet{}
}

// <no value>
type RecordGet struct {
	RecordNumber byte
}

func (cmd RecordGet) CommandClassID() cc.CommandClassID {
	return 0x4C
}

func (cmd RecordGet) CommandID() cc.CommandID {
	return CommandRecordGet
}

func (cmd RecordGet) CommandIDString() string {
	return "RECORD_GET"
}

func (cmd *RecordGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	payload := make([]byte, len(data))
	copy(payload, data)

	if len(payload) < 2 {
		return errors.New("Payload length underflow")
	}

	i := 2

	if len(payload) <= i {
		return fmt.Errorf("slice index out of bounds (.RecordNumber) %d<=%d", len(payload), i)
	}

	cmd.RecordNumber = payload[i]
	i++

	return nil
}

func (cmd *RecordGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.RecordNumber)

	return
}
