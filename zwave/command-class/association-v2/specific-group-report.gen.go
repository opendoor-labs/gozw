// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package associationv2

import "errors"

// <no value>

type AssociationSpecificGroupReport struct {
	Group byte
}

func (cmd *AssociationSpecificGroupReport) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Group = payload[i]
	i++

	return nil
}

func (cmd *AssociationSpecificGroupReport) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.Group)

	return
}