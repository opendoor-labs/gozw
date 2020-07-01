// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package versionv3

import (
	"encoding/gob"
	"errors"

	"github.com/opendoor-labs/gozw/cc"
)

const CommandCommandClassGet cc.CommandID = 0x13

func init() {
	gob.Register(CommandClassGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x86),
		Command:      cc.CommandID(0x13),
		Version:      3,
	}, NewCommandClassGet)
}

func NewCommandClassGet() cc.Command {
	return &CommandClassGet{}
}

// <no value>
type CommandClassGet struct {
	RequestedCommandClass byte
}

func (cmd CommandClassGet) CommandClassID() cc.CommandClassID {
	return 0x86
}

func (cmd CommandClassGet) CommandID() cc.CommandID {
	return CommandCommandClassGet
}

func (cmd CommandClassGet) CommandIDString() string {
	return "VERSION_COMMAND_CLASS_GET"
}

func (cmd *CommandClassGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	payload := make([]byte, len(data))
	copy(payload, data)

	if len(payload) < 2 {
		return errors.New("Payload length underflow")
	}

	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.RequestedCommandClass = payload[i]
	i++

	return nil
}

func (cmd *CommandClassGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.RequestedCommandClass)

	return
}

func Noop() {
	// does nothing, just here to allow
	// consumers to invoke the init function
}
