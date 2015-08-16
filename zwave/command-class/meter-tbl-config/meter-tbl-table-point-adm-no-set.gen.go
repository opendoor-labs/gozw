// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package metertblconfig

import "errors"

// <no value>

type MeterTblTablePointAdmNoSet struct {
	Properties1 struct {
		NumberOfCharacters byte
	}

	MeterPointAdmNumberCharacter []byte
}

func (cmd *MeterTblTablePointAdmNoSet) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.NumberOfCharacters = (payload[i] & 0x1F)

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.MeterPointAdmNumberCharacter = payload[i : i+0]
	i += 0

	return nil
}

func (cmd *MeterTblTablePointAdmNoSet) MarshalBinary() (payload []byte, err error) {

	{
		var val byte

		val |= (cmd.Properties1.NumberOfCharacters) & byte(0x1F)

		payload = append(payload, val)
	}

	if cmd.MeterPointAdmNumberCharacter != nil && len(cmd.MeterPointAdmNumberCharacter) > 0 {
		payload = append(payload, cmd.MeterPointAdmNumberCharacter...)
	}

	return
}