// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package powerlevel

import "errors"

// <no value>

type PowerlevelSet struct {
	PowerLevel byte

	Timeout byte
}

func (cmd *PowerlevelSet) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.PowerLevel = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Timeout = payload[i]
	i++

	return nil
}

func (cmd *PowerlevelSet) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.PowerLevel)

	payload = append(payload, cmd.Timeout)

	return
}