// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package doorlockv3

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandConfigurationSet cc.CommandID = 0x04

func init() {
	gob.Register(ConfigurationSet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x62),
		Command:      cc.CommandID(0x04),
		Version:      3,
	}, NewConfigurationSet)
}

func NewConfigurationSet() cc.Command {
	return &ConfigurationSet{}
}

// <no value>
type ConfigurationSet struct {
	OperationType byte

	Properties1 struct {
		InsideDoorHandlesState byte

		OutsideDoorHandlesState byte
	}

	LockTimeoutMinutes byte

	LockTimeoutSeconds byte
}

func (cmd ConfigurationSet) CommandClassID() cc.CommandClassID {
	return 0x62
}

func (cmd ConfigurationSet) CommandID() cc.CommandID {
	return CommandConfigurationSet
}

func (cmd ConfigurationSet) CommandIDString() string {
	return "DOOR_LOCK_CONFIGURATION_SET"
}

func (cmd *ConfigurationSet) UnmarshalBinary(data []byte) error {
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

	cmd.OperationType = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.InsideDoorHandlesState = (payload[i] & 0x0F)

	cmd.Properties1.OutsideDoorHandlesState = (payload[i] & 0xF0) >> 4

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.LockTimeoutMinutes = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.LockTimeoutSeconds = payload[i]
	i++

	return nil
}

func (cmd *ConfigurationSet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.OperationType)

	{
		var val byte

		val |= (cmd.Properties1.InsideDoorHandlesState) & byte(0x0F)

		val |= (cmd.Properties1.OutsideDoorHandlesState << byte(4)) & byte(0xF0)

		payload = append(payload, val)
	}

	payload = append(payload, cmd.LockTimeoutMinutes)

	payload = append(payload, cmd.LockTimeoutSeconds)

	return
}