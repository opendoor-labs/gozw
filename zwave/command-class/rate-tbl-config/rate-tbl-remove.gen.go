// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package ratetblconfig

import "errors"

// <no value>

type RateTblRemove struct {
	Properties1 struct {
		RateParameterSetIds byte
	}

	RateParameterSetId []byte
}

func (cmd *RateTblRemove) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.RateParameterSetIds = (payload[i] & 0x3F)

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.RateParameterSetId = payload[i : i+0]
	i += 0

	return nil
}

func (cmd *RateTblRemove) MarshalBinary() (payload []byte, err error) {

	{
		var val byte

		val |= (cmd.Properties1.RateParameterSetIds) & byte(0x3F)

		payload = append(payload, val)
	}

	if cmd.RateParameterSetId != nil && len(cmd.RateParameterSetId) > 0 {
		payload = append(payload, cmd.RateParameterSetId...)
	}

	return
}
