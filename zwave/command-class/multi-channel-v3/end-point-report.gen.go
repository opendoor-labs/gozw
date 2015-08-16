// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package multichannelv3

import "errors"

// <no value>

type MultiChannelEndPointReport struct {
	Properties1 struct {
		Identical bool

		Dynamic bool
	}

	Properties2 struct {
		EndPoints byte
	}
}

func (cmd *MultiChannelEndPointReport) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	if payload[i]&0x40 == 0x40 {
		cmd.Properties1.Identical = true
	} else {
		cmd.Properties1.Identical = false
	}

	if payload[i]&0x80 == 0x80 {
		cmd.Properties1.Dynamic = true
	} else {
		cmd.Properties1.Dynamic = false
	}

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties2.EndPoints = (payload[i] & 0x7F)

	i += 1

	return nil
}

func (cmd *MultiChannelEndPointReport) MarshalBinary() (payload []byte, err error) {

	{
		var val byte

		if cmd.Properties1.Identical {
			val |= byte(0x40) // flip bits on
		} else {
			val &= ^byte(0x40) // flip bits off
		}

		if cmd.Properties1.Dynamic {
			val |= byte(0x80) // flip bits on
		} else {
			val &= ^byte(0x80) // flip bits off
		}

		payload = append(payload, val)
	}

	{
		var val byte

		val |= (cmd.Properties2.EndPoints) & byte(0x7F)

		payload = append(payload, val)
	}

	return
}
