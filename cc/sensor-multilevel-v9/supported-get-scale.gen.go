// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package sensormultilevelv9

import (
	"encoding/gob"
	"errors"

	"github.com/opendoor-labs/gozw/cc"
)

const CommandSupportedGetScale cc.CommandID = 0x03

func init() {
	gob.Register(SupportedGetScale{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x31),
		Command:      cc.CommandID(0x03),
		Version:      9,
	}, NewSupportedGetScale)
}

func NewSupportedGetScale() cc.Command {
	return &SupportedGetScale{}
}

// <no value>
type SupportedGetScale struct {
	SensorType byte
}

func (cmd SupportedGetScale) CommandClassID() cc.CommandClassID {
	return 0x31
}

func (cmd SupportedGetScale) CommandID() cc.CommandID {
	return CommandSupportedGetScale
}

func (cmd SupportedGetScale) CommandIDString() string {
	return "SENSOR_MULTILEVEL_SUPPORTED_GET_SCALE"
}

func (cmd *SupportedGetScale) UnmarshalBinary(data []byte) error {
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

	cmd.SensorType = payload[i]
	i++

	return nil
}

func (cmd *SupportedGetScale) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.SensorType)

	return
}
