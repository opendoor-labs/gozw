// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package multichannelv2

import "errors"

// <no value>

type MultiInstanceCmdEncap struct {
	Properties1 struct {
		Instance byte
	}

	CommandClass byte

	Command byte

	Parameter []byte
}

func (cmd *MultiInstanceCmdEncap) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.Instance = (payload[i] & 0x7F)

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.CommandClass = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Command = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Parameter = payload[i:]

	return nil
}

func (cmd *MultiInstanceCmdEncap) MarshalBinary() (payload []byte, err error) {

	{
		var val byte

		val |= (cmd.Properties1.Instance) & byte(0x7F)

		payload = append(payload, val)
	}

	payload = append(payload, cmd.CommandClass)

	payload = append(payload, cmd.Command)

	payload = append(payload, cmd.Parameter...)

	return
}