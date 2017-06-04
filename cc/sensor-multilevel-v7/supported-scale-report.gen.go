// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package sensormultilevelv7

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandSupportedScaleReport cc.CommandID = 0x06

func init() {
	gob.Register(SupportedScaleReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x31),
		Command:      cc.CommandID(0x06),
		Version:      7,
	}, NewSupportedScaleReport)
}

func NewSupportedScaleReport() cc.Command {
	return &SupportedScaleReport{}
}

// <no value>
type SupportedScaleReport struct {
	SensorType byte

	Properties1 struct {
		ScaleBitMask byte
	}
}

func (cmd SupportedScaleReport) CommandClassID() cc.CommandClassID {
	return 0x31
}

func (cmd SupportedScaleReport) CommandID() cc.CommandID {
	return CommandSupportedScaleReport
}

func (cmd SupportedScaleReport) CommandIDString() string {
	return "SENSOR_MULTILEVEL_SUPPORTED_SCALE_REPORT"
}

func (cmd *SupportedScaleReport) UnmarshalBinary(data []byte) error {
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

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.ScaleBitMask = (payload[i] & 0x0F)

	i += 1

	return nil
}

func (cmd *SupportedScaleReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.SensorType)

	{
		var val byte

		val |= (cmd.Properties1.ScaleBitMask) & byte(0x0F)

		payload = append(payload, val)
	}

	return
}
