// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package switchcolorv2

import (
	"encoding/gob"
	"errors"

	"github.com/opendoor-labs/gozw/cc"
)

const CommandStopLevelChange cc.CommandID = 0x07

func init() {
	gob.Register(StopLevelChange{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x33),
		Command:      cc.CommandID(0x07),
		Version:      2,
	}, NewStopLevelChange)
}

func NewStopLevelChange() cc.Command {
	return &StopLevelChange{}
}

// <no value>
type StopLevelChange struct {
	ColorComponentId byte
}

func (cmd StopLevelChange) CommandClassID() cc.CommandClassID {
	return 0x33
}

func (cmd StopLevelChange) CommandID() cc.CommandID {
	return CommandStopLevelChange
}

func (cmd StopLevelChange) CommandIDString() string {
	return "SWITCH_COLOR_STOP_LEVEL_CHANGE"
}

func (cmd *StopLevelChange) UnmarshalBinary(data []byte) error {
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

	cmd.ColorComponentId = payload[i]
	i++

	return nil
}

func (cmd *StopLevelChange) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.ColorComponentId)

	return
}
