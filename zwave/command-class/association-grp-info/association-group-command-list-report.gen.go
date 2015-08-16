// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package associationgrpinfo

import "errors"

// <no value>

type AssociationGroupCommandListReport struct {
	GroupingIdentifier byte

	ListLength byte

	Command []byte
}

func (cmd *AssociationGroupCommandListReport) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.GroupingIdentifier = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ListLength = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Command = payload[i : i+1]
	i += 1

	return nil
}

func (cmd *AssociationGroupCommandListReport) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.GroupingIdentifier)

	payload = append(payload, cmd.ListLength)

	if cmd.Command != nil && len(cmd.Command) > 0 {
		payload = append(payload, cmd.Command...)
	}

	return
}
