// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package security

import "errors"

// <no value>

type NetworkKeySet struct {
	NetworkKeyByte []byte
}

func (cmd *NetworkKeySet) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.NetworkKeyByte = payload[i:]

	return nil
}

func (cmd *NetworkKeySet) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.NetworkKeyByte...)

	return
}
