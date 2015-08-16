// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package scheduleentrylockv3

import "errors"

// <no value>

type ScheduleEntryLockWeekDayGet struct {
	UserIdentifier byte

	ScheduleSlotId byte
}

func (cmd *ScheduleEntryLockWeekDayGet) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.UserIdentifier = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ScheduleSlotId = payload[i]
	i++

	return nil
}

func (cmd *ScheduleEntryLockWeekDayGet) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.UserIdentifier)

	payload = append(payload, cmd.ScheduleSlotId)

	return
}
