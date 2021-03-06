// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package doorlockv4

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandOperationSet cc.CommandID = 0x01

func init() {
	gob.Register(OperationSet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x62),
		Command:      cc.CommandID(0x01),
		Version:      4,
	}, NewOperationSet)
}

func NewOperationSet() cc.Command {
	return &OperationSet{}
}

// <no value>
type OperationSet struct {
	DoorLockMode byte
}

func (cmd OperationSet) CommandClassID() cc.CommandClassID {
	return 0x62
}

func (cmd OperationSet) CommandID() cc.CommandID {
	return CommandOperationSet
}

func (cmd OperationSet) CommandIDString() string {
	return "DOOR_LOCK_OPERATION_SET"
}

func (cmd *OperationSet) UnmarshalBinary(data []byte) error {
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

	cmd.DoorLockMode = payload[i]
	i++

	return nil
}

func (cmd *OperationSet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.DoorLockMode)

	return
}
