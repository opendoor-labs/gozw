// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package thermostatfanmodev2

import "errors"

// <no value>

type ThermostatFanModeSupportedReport struct {
	BitMask byte
}

func (cmd *ThermostatFanModeSupportedReport) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.BitMask = payload[i]
	i++

	return nil
}

func (cmd *ThermostatFanModeSupportedReport) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.BitMask)

	return
}
