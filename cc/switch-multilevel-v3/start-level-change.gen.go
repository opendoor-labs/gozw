// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package switchmultilevelv3

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandStartLevelChange cc.CommandID = 0x04

func init() {
	gob.Register(StartLevelChange{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x26),
		Command:      cc.CommandID(0x04),
		Version:      3,
	}, NewStartLevelChange)
}

func NewStartLevelChange() cc.Command {
	return &StartLevelChange{}
}

// <no value>
type StartLevelChange struct {
	Properties1 struct {
		IgnoreStartLevel bool

		IncDec byte

		UpDown byte
	}

	StartLevel byte

	DimmingDuration byte

	StepSize byte
}

func (cmd StartLevelChange) CommandClassID() cc.CommandClassID {
	return 0x26
}

func (cmd StartLevelChange) CommandID() cc.CommandID {
	return CommandStartLevelChange
}

func (cmd StartLevelChange) CommandIDString() string {
	return "SWITCH_MULTILEVEL_START_LEVEL_CHANGE"
}

func (cmd *StartLevelChange) UnmarshalBinary(data []byte) error {
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

	cmd.Properties1.IncDec = (payload[i] & 0x18) >> 3

	cmd.Properties1.UpDown = (payload[i] & 0xC0) >> 6

	cmd.Properties1.IgnoreStartLevel = payload[i]&0x20 == 0x20

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StartLevel = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.DimmingDuration = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StepSize = payload[i]
	i++

	return nil
}

func (cmd *StartLevelChange) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	{
		var val byte

		val |= (cmd.Properties1.IncDec << byte(3)) & byte(0x18)

		val |= (cmd.Properties1.UpDown << byte(6)) & byte(0xC0)

		if cmd.Properties1.IgnoreStartLevel {
			val |= byte(0x20) // flip bits on
		} else {
			val &= ^byte(0x20) // flip bits off
		}

		payload = append(payload, val)
	}

	payload = append(payload, cmd.StartLevel)

	payload = append(payload, cmd.DimmingDuration)

	payload = append(payload, cmd.StepSize)

	return
}
