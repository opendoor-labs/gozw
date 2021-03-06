// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package versionv3

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandCapabilitiesReport cc.CommandID = 0x16

func init() {
	gob.Register(CapabilitiesReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x86),
		Command:      cc.CommandID(0x16),
		Version:      3,
	}, NewCapabilitiesReport)
}

func NewCapabilitiesReport() cc.Command {
	return &CapabilitiesReport{}
}

// <no value>
type CapabilitiesReport struct {
	Properties1 struct {
		Version bool

		CommandClass bool

		ZWaveSoftware bool
	}
}

func (cmd CapabilitiesReport) CommandClassID() cc.CommandClassID {
	return 0x86
}

func (cmd CapabilitiesReport) CommandID() cc.CommandID {
	return CommandCapabilitiesReport
}

func (cmd CapabilitiesReport) CommandIDString() string {
	return "VERSION_CAPABILITIES_REPORT"
}

func (cmd *CapabilitiesReport) UnmarshalBinary(data []byte) error {
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

	cmd.Properties1.Version = payload[i]&0x01 == 0x01

	cmd.Properties1.CommandClass = payload[i]&0x02 == 0x02

	cmd.Properties1.ZWaveSoftware = payload[i]&0x04 == 0x04

	i += 1

	return nil
}

func (cmd *CapabilitiesReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	{
		var val byte

		if cmd.Properties1.Version {
			val |= byte(0x01) // flip bits on
		} else {
			val &= ^byte(0x01) // flip bits off
		}

		if cmd.Properties1.CommandClass {
			val |= byte(0x02) // flip bits on
		} else {
			val &= ^byte(0x02) // flip bits off
		}

		if cmd.Properties1.ZWaveSoftware {
			val |= byte(0x04) // flip bits on
		} else {
			val &= ^byte(0x04) // flip bits off
		}

		payload = append(payload, val)
	}

	return
}
