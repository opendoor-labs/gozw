// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package notificationv4

import "errors"

// <no value>

type NotificationSet struct {
	NotificationType byte

	NotificationStatus byte
}

func (cmd *NotificationSet) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.NotificationType = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.NotificationStatus = payload[i]
	i++

	return nil
}

func (cmd *NotificationSet) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.NotificationType)

	payload = append(payload, cmd.NotificationStatus)

	return
}