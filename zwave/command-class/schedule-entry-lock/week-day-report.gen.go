// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package scheduleentrylock

import "errors"

// <no value>

type ScheduleEntryLockWeekDayReport struct {
	UserIdentifier byte

	ScheduleSlotId byte

	DayOfWeek byte

	StartHour byte

	StartMinute byte

	StopHour byte

	StopMinute byte
}

func (cmd *ScheduleEntryLockWeekDayReport) UnmarshalBinary(payload []byte) error {
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

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.DayOfWeek = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StartHour = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StartMinute = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StopHour = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StopMinute = payload[i]
	i++

	return nil
}

func (cmd *ScheduleEntryLockWeekDayReport) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.UserIdentifier)

	payload = append(payload, cmd.ScheduleSlotId)

	payload = append(payload, cmd.DayOfWeek)

	payload = append(payload, cmd.StartHour)

	payload = append(payload, cmd.StartMinute)

	payload = append(payload, cmd.StopHour)

	payload = append(payload, cmd.StopMinute)

	return
}
