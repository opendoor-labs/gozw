// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package switchmultilevel

import "errors"

// <no value>

type SwitchMultilevelSet struct {
	Value byte
}

func (cmd *SwitchMultilevelSet) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Value = payload[i]
	i++

	return nil
}

func (cmd *SwitchMultilevelSet) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.Value)

	return
}