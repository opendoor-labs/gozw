// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package sensormultilevelv5

import "errors"

// <no value>

type SensorMultilevelSupportedGetScale struct {
	SensorType byte
}

func (cmd *SensorMultilevelSupportedGetScale) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.SensorType = payload[i]
	i++

	return nil
}

func (cmd *SensorMultilevelSupportedGetScale) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.SensorType)

	return
}
