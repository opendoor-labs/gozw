// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package sensorconfiguration

import "errors"

// <no value>

type SensorTriggerLevelSet struct {
	Properties1 struct {
		Current bool

		Default bool
	}

	SensorType byte

	Properties2 struct {
		Size byte

		Scale byte

		Precision byte
	}

	TriggerValue []byte
}

func (cmd *SensorTriggerLevelSet) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	if payload[i]&0x40 == 0x40 {
		cmd.Properties1.Current = true
	} else {
		cmd.Properties1.Current = false
	}

	if payload[i]&0x80 == 0x80 {
		cmd.Properties1.Default = true
	} else {
		cmd.Properties1.Default = false
	}

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.SensorType = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties2.Size = (payload[i] & 0x07)

	cmd.Properties2.Scale = (payload[i] & 0x18) >> 3

	cmd.Properties2.Precision = (payload[i] & 0xE0) >> 5

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.TriggerValue = payload[i : i+2]
	i += 2

	return nil
}

func (cmd *SensorTriggerLevelSet) MarshalBinary() (payload []byte, err error) {

	{
		var val byte

		if cmd.Properties1.Current {
			val |= byte(0x40) // flip bits on
		} else {
			val &= ^byte(0x40) // flip bits off
		}

		if cmd.Properties1.Default {
			val |= byte(0x80) // flip bits on
		} else {
			val &= ^byte(0x80) // flip bits off
		}

		payload = append(payload, val)
	}

	payload = append(payload, cmd.SensorType)

	{
		var val byte

		val |= (cmd.Properties2.Size) & byte(0x07)

		val |= (cmd.Properties2.Scale << byte(3)) & byte(0x18)

		val |= (cmd.Properties2.Precision << byte(5)) & byte(0xE0)

		payload = append(payload, val)
	}

	if cmd.TriggerValue != nil && len(cmd.TriggerValue) > 0 {
		payload = append(payload, cmd.TriggerValue...)
	}

	return
}
