// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package networkmanagementinclusion

import "errors"

// <no value>

type NodeNeighborUpdateRequest struct {
	SeqNo byte

	NodeId byte
}

func (cmd *NodeNeighborUpdateRequest) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.SeqNo = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.NodeId = payload[i]
	i++

	return nil
}

func (cmd *NodeNeighborUpdateRequest) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.SeqNo)

	payload = append(payload, cmd.NodeId)

	return
}