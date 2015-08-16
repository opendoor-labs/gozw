// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package climatecontrolschedule

import "errors"

// <no value>

type ScheduleOverrideReport struct {
	Properties1 struct {
		OverrideType byte
	}

	OverrideState byte
}

func (cmd *ScheduleOverrideReport) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.OverrideType = (payload[i] & 0x03)

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.OverrideState = payload[i]
	i++

	return nil
}

func (cmd *ScheduleOverrideReport) MarshalBinary() (payload []byte, err error) {

	{
		var val byte

		val |= (cmd.Properties1.OverrideType) & byte(0x03)

		payload = append(payload, val)
	}

	payload = append(payload, cmd.OverrideState)

	return
}
