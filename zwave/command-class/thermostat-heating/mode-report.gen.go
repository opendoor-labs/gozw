// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package thermostatheating

import "errors"

// <no value>

type ThermostatHeatingModeReport struct {
	Mode byte
}

func (cmd *ThermostatHeatingModeReport) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Mode = payload[i]
	i++

	return nil
}

func (cmd *ThermostatHeatingModeReport) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.Mode)

	return
}
