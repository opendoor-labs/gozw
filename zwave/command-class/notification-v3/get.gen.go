// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package notificationv3

import "errors"

// <no value>

type NotificationGet struct {
	V1AlarmType byte

	NotificationType byte

	Event byte
}

func (cmd *NotificationGet) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.V1AlarmType = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.NotificationType = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Event = payload[i]
	i++

	return nil
}

func (cmd *NotificationGet) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.V1AlarmType)

	payload = append(payload, cmd.NotificationType)

	payload = append(payload, cmd.Event)

	return
}