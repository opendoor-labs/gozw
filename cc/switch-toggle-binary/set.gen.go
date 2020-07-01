// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package switchtogglebinary

import (
	"encoding/gob"

	"github.com/opendoor-labs/gozw/cc"
)

const CommandSet cc.CommandID = 0x01

func init() {
	gob.Register(Set{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x28),
		Command:      cc.CommandID(0x01),
		Version:      1,
	}, NewSet)
}

func NewSet() cc.Command {
	return &Set{}
}

// <no value>
type Set struct {
}

func (cmd Set) CommandClassID() cc.CommandClassID {
	return 0x28
}

func (cmd Set) CommandID() cc.CommandID {
	return CommandSet
}

func (cmd Set) CommandIDString() string {
	return "SWITCH_TOGGLE_BINARY_SET"
}

func (cmd *Set) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	return nil
}

func (cmd *Set) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	return
}

func Noop() {
	// does nothing, just here to allow
	// consumers to invoke the init function
}
