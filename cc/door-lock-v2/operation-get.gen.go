// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package doorlockv2

import (
	"encoding/gob"

	"github.com/opendoor-labs/gozw/cc"
)

const CommandOperationGet cc.CommandID = 0x02

func init() {
	gob.Register(OperationGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x62),
		Command:      cc.CommandID(0x02),
		Version:      2,
	}, NewOperationGet)
}

func NewOperationGet() cc.Command {
	return &OperationGet{}
}

// <no value>
type OperationGet struct {
}

func (cmd OperationGet) CommandClassID() cc.CommandClassID {
	return 0x62
}

func (cmd OperationGet) CommandID() cc.CommandID {
	return CommandOperationGet
}

func (cmd OperationGet) CommandIDString() string {
	return "DOOR_LOCK_OPERATION_GET"
}

func (cmd *OperationGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	return nil
}

func (cmd *OperationGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	return
}
