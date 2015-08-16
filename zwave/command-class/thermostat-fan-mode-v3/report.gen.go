// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package thermostatfanmodev3

import "errors"

// <no value>

type ThermostatFanModeReport struct {
	Properties1 struct {
		Off bool

		FanMode byte
	}
}

func (cmd *ThermostatFanModeReport) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.FanMode = (payload[i] & 0x0F)

	if payload[i]&0x80 == 0x80 {
		cmd.Properties1.Off = true
	} else {
		cmd.Properties1.Off = false
	}

	i += 1

	return nil
}

func (cmd *ThermostatFanModeReport) MarshalBinary() (payload []byte, err error) {

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
