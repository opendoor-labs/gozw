// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package thermostatsetpoint

import "errors"

// <no value>

type ThermostatSetpointSupportedReport struct {
	BitMask byte
}

func (cmd *ThermostatSetpointSupportedReport) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.BitMask = payload[i]
	i++

	return nil
}

func (cmd *ThermostatSetpointSupportedReport) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.BitMask)

	return
}
