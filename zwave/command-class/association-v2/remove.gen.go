// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package associationv2

import "errors"

// <no value>

type AssociationRemove struct {
	GroupingIdentifier byte

	NodeId []byte
}

func (cmd *AssociationRemove) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.GroupingIdentifier = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.NodeId = payload[i:]

	return nil
}

func (cmd *AssociationRemove) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.GroupingIdentifier)

	payload = append(payload, cmd.NodeId...)

	return
}
