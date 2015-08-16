// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package networkmanagementinclusion

import "errors"

// <no value>

type FailedNodeRemoveStatus struct {
	SeqNo byte

	Status byte

	NodeId byte
}

func (cmd *FailedNodeRemoveStatus) UnmarshalBinary(payload []byte) error {
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

	cmd.NodeId = payload[i]
	i++

	return nil
}

func (cmd *FailedNodeRemoveStatus) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.SeqNo)

	payload = append(payload, cmd.Status)

	payload = append(payload, cmd.NodeId)

	return
}