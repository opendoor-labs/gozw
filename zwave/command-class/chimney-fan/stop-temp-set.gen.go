// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package chimneyfan

import "errors"

// <no value>

type ChimneyFanStopTempSet struct {
	Properties1 struct {
		Size byte

		Scale byte

		Precision byte
	}

	Value []byte
}

func (cmd *ChimneyFanStopTempSet) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.Size = (payload[i] & 0x07)

	cmd.Properties1.Scale = (payload[i] & 0x18) >> 3

	cmd.Properties1.Precision = (payload[i] & 0xE0) >> 5

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Value = payload[i : i+0]
	i += 0

	return nil
}

func (cmd *ChimneyFanStopTempSet) MarshalBinary() (payload []byte, err error) {

	{
		var val byte

		val |= (cmd.Properties1.Size) & byte(0x07)

		val |= (cmd.Properties1.Scale << byte(3)) & byte(0x18)

		val |= (cmd.Properties1.Precision << byte(5)) & byte(0xE0)

		payload = append(payload, val)
	}

	if cmd.Value != nil && len(cmd.Value) > 0 {
		payload = append(payload, cmd.Value...)
	}

	return
}
