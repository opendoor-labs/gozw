// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package usercode

import "errors"

// <no value>

type UserCodeSet struct {
	UserIdentifier byte

	UserIdStatus byte

	UserCode string
}

func (cmd *UserCodeSet) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.UserIdentifier = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.UserIdStatus = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.UserCode = string(payload[i : i+10])

	i += 10

	return nil
}

func (cmd *UserCodeSet) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.UserIdentifier)

	payload = append(payload, cmd.UserIdStatus)

	if paramLen := len(cmd.UserCode); paramLen > 10 {
		return nil, errors.New("Length overflow in array parameter UserCode")
	}

	payload = append(payload, []byte(cmd.UserCode)...)

	return
}
