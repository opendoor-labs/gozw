// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package sensorconfiguration

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandSensorTriggerLevelReport cc.CommandID = 0x03

func init() {
	gob.Register(SensorTriggerLevelReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x9E),
		Command:      cc.CommandID(0x03),
		Version:      1,
	}, NewSensorTriggerLevelReport)
}

func NewSensorTriggerLevelReport() cc.Command {
	return &SensorTriggerLevelReport{}
}

// <no value>
type SensorTriggerLevelReport struct {
	SensorType byte

	Properties1 struct {
		Size byte

		Scale byte

		Precision byte
	}

	TriggerValue []byte
}

func (cmd SensorTriggerLevelReport) CommandClassID() cc.CommandClassID {
	return 0x9E
}

func (cmd SensorTriggerLevelReport) CommandID() cc.CommandID {
	return CommandSensorTriggerLevelReport
}

func (cmd SensorTriggerLevelReport) CommandIDString() string {
	return "SENSOR_TRIGGER_LEVEL_REPORT"
}

func (cmd *SensorTriggerLevelReport) UnmarshalBinary(data []byte) error {
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

	cmd.Properties1.Size = (payload[i] & 0x07)

	cmd.Properties1.Scale = (payload[i] & 0x18) >> 3

	cmd.Properties1.Precision = (payload[i] & 0xE0) >> 5

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	{
		length := (payload[1+2]) & 0x07
		cmd.TriggerValue = payload[i : i+int(length)]
		i += int(length)
	}

	return nil
}

func (cmd *SensorTriggerLevelReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.SensorType)

	{
		var val byte

		val |= (cmd.Properties1.Size) & byte(0x07)

		val |= (cmd.Properties1.Scale << byte(3)) & byte(0x18)

		val |= (cmd.Properties1.Precision << byte(5)) & byte(0xE0)

		payload = append(payload, val)
	}

	if cmd.TriggerValue != nil && len(cmd.TriggerValue) > 0 {
		payload = append(payload, cmd.TriggerValue...)
	}

	return
}
