// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package thermostatheating

import "errors"

// <no value>

type ThermostatHeatingRelayStatusReport struct {
	RelayStatus byte
}

func (cmd *ThermostatHeatingRelayStatusReport) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.RelayStatus = payload[i]
	i++

	return nil
}

func (cmd *ThermostatHeatingRelayStatusReport) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.RelayStatus)

	return
}