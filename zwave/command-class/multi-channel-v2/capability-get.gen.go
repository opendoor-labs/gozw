// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package multichannelv2

import "errors"

// <no value>

type MultiChannelCapabilityGet struct {
	Properties1 struct {
		EndPoint byte
	}
}

func (cmd *MultiChannelCapabilityGet) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.EndPoint = (payload[i] & 0x7F)

	i += 1

	return nil
}

func (cmd *MultiChannelCapabilityGet) MarshalBinary() (payload []byte, err error) {

	{
		var val byte

		val |= (cmd.Properties1.EndPoint) & byte(0x7F)

		payload = append(payload, val)
	}

	return
}