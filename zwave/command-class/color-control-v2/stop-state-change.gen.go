// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package colorcontrolv2

import "errors"

// <no value>

type StopStateChange struct {
	CapabilityId byte
}

func (cmd *StopStateChange) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.CapabilityId = payload[i]
	i++

	return nil
}

func (cmd *StopStateChange) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.CapabilityId)

	return
}
