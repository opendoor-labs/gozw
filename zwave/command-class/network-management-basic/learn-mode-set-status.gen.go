// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package networkmanagementbasic

import "errors"

// <no value>

type LearnModeSetStatus struct {
	SeqNo byte

	Status byte

	NewNodeId byte
}

func (cmd *LearnModeSetStatus) UnmarshalBinary(payload []byte) error {
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

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.NewNodeId = payload[i]
	i++

	return nil
}

func (cmd *LearnModeSetStatus) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.SeqNo)

	payload = append(payload, cmd.Status)

	payload = append(payload, cmd.NewNodeId)

	return
}
