// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package networkmanagementbasic

import "errors"

// <no value>

type NetworkUpdateRequestStatus struct {
	SeqNo byte

	Status byte
}

func (cmd *NetworkUpdateRequestStatus) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.SeqNo = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Status = payload[i]
	i++

	return nil
}

func (cmd *NetworkUpdateRequestStatus) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.SeqNo)

	payload = append(payload, cmd.Status)

	return
}
