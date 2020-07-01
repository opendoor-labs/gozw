// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package thermostatfanmodev3

import (
	"encoding/gob"
	"errors"

	"github.com/opendoor-labs/gozw/cc"
)

const CommandReport cc.CommandID = 0x03

func init() {
	gob.Register(Report{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x44),
		Command:      cc.CommandID(0x03),
		Version:      3,
	}, NewReport)
}

func NewReport() cc.Command {
	return &Report{}
}

// <no value>
type Report struct {
	Properties1 struct {
		Off bool

		FanMode byte
	}
}

func (cmd Report) CommandClassID() cc.CommandClassID {
	return 0x44
}

func (cmd Report) CommandID() cc.CommandID {
	return CommandReport
}

func (cmd Report) CommandIDString() string {
	return "THERMOSTAT_FAN_MODE_REPORT"
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

	cmd.Properties1.FanMode = (payload[i] & 0x0F)

	cmd.Properties1.Off = payload[i]&0x80 == 0x80

	i += 1

	return nil
}

func (cmd *Report) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	{
		var val byte

		val |= (cmd.Properties1.FanMode) & byte(0x0F)

		if cmd.Properties1.Off {
			val |= byte(0x80) // flip bits on
		} else {
			val &= ^byte(0x80) // flip bits off
		}

		payload = append(payload, val)
	}

	return
}
