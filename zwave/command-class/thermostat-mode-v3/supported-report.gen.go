// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package thermostatmodev3

import "errors"

// <no value>

type ThermostatModeSupportedReport struct {
	BitMask byte
}

func (cmd *ThermostatModeSupportedReport) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.BitMask = payload[i]
	i++

	return nil
}

func (cmd *ThermostatModeSupportedReport) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.BitMask)

	return
}