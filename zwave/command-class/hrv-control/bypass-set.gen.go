// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package hrvcontrol

import "errors"

// <no value>

type HrvControlBypassSet struct {
	Bypass byte
}

func (cmd *HrvControlBypassSet) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Bypass = payload[i]
	i++

	return nil
}

func (cmd *HrvControlBypassSet) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.Bypass)

	return
}
