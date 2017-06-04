// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package sensormultilevelv10

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandReport cc.CommandID = 0x05

func init() {
	gob.Register(Report{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x31),
		Command:      cc.CommandID(0x05),
		Version:      10,
	}, NewReport)
}

func NewReport() cc.Command {
	return &Report{}
}

// <no value>
type Report struct {
	SensorType byte

	Level struct {
		Size byte

		Scale byte

		Precision byte
	}

	SensorValue []byte
}

func (cmd Report) CommandClassID() cc.CommandClassID {
	return 0x31
}

func (cmd Report) CommandID() cc.CommandID {
	return CommandReport
}

func (cmd Report) CommandIDString() string {
	return "SENSOR_MULTILEVEL_REPORT"
}

func (cmd *Report) UnmarshalBinary(data []byte) error {
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

	cmd.Level.Size = (payload[i] & 0x07)

	cmd.Level.Scale = (payload[i] & 0x18) >> 3

	cmd.Level.Precision = (payload[i] & 0xE0) >> 5

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	{
		length := (payload[1+2] >> 0) & 0x07
		cmd.SensorValue = payload[i : i+int(length)]
		i += int(length)
	}

	return nil
}

func (cmd *Report) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.SensorType)

	{
		var val byte

		val |= (cmd.Level.Size) & byte(0x07)

		val |= (cmd.Level.Scale << byte(3)) & byte(0x18)

		val |= (cmd.Level.Precision << byte(5)) & byte(0xE0)

		payload = append(payload, val)
	}

	if cmd.SensorValue != nil && len(cmd.SensorValue) > 0 {
		payload = append(payload, cmd.SensorValue...)
	}

	return
}
