// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package hrvcontrol

import "errors"

// <no value>

type HrvControlVentilationRateSet struct {
	VentilationRate byte
}

func (cmd *HrvControlVentilationRateSet) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.VentilationRate = payload[i]
	i++

	return nil
}

func (cmd *HrvControlVentilationRateSet) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.VentilationRate)

	return
}