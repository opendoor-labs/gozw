// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package screenmd

import "errors"

// <no value>

type ScreenMdGet struct {
	NumberOfReports byte

	NodeId byte
}

func (cmd *ScreenMdGet) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.NumberOfReports = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.NodeId = payload[i]
	i++

	return nil
}

func (cmd *ScreenMdGet) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.NumberOfReports)

	payload = append(payload, cmd.NodeId)

	return
}
