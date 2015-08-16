// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package association

import "errors"

// <no value>

type AssociationReport struct {
	GroupingIdentifier byte

	MaxNodesSupported byte

	ReportsToFollow byte

	Nodeid []byte
}

func (cmd *AssociationReport) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.GroupingIdentifier = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.MaxNodesSupported = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ReportsToFollow = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Nodeid = payload[i:]

	return nil
}

func (cmd *AssociationReport) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.GroupingIdentifier)

	payload = append(payload, cmd.MaxNodesSupported)

	payload = append(payload, cmd.ReportsToFollow)

	payload = append(payload, cmd.Nodeid...)

	return
}
