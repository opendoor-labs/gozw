// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package ratetblmonitor

import "errors"

// <no value>

type RateTblGet struct {
	RateParameterSetId byte
}

func (cmd *RateTblGet) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.RateParameterSetId = payload[i]
	i++

	return nil
}

func (cmd *RateTblGet) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.RateParameterSetId)

	return
}
