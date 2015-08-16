// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package thermostatsetback

import "errors"

// <no value>

type ThermostatSetbackReport struct {
	Properties1 struct {
		SetbackType byte
	}

	SetbackState byte
}

func (cmd *ThermostatSetbackReport) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.SetbackType = (payload[i] & 0x03)

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.SetbackState = payload[i]
	i++

	return nil
}

func (cmd *ThermostatSetbackReport) MarshalBinary() (payload []byte, err error) {

	{
		var val byte

		val |= (cmd.Properties1.SetbackType) & byte(0x03)

		payload = append(payload, val)
	}

	payload = append(payload, cmd.SetbackState)

	return
}