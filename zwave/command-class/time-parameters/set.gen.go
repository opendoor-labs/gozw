// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package timeparameters

import (
	"encoding/binary"
	"errors"
)

// <no value>

type TimeParametersSet struct {
	Year uint16

	Month byte

	Day byte

	HourUtc byte

	MinuteUtc byte

	SecondUtc byte
}

func (cmd *TimeParametersSet) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Year = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Month = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Day = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.HourUtc = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.MinuteUtc = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.SecondUtc = payload[i]
	i++

	return nil
}

func (cmd *TimeParametersSet) MarshalBinary() (payload []byte, err error) {

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.Year)
		payload = append(payload, buf...)
	}

	payload = append(payload, cmd.Month)

	payload = append(payload, cmd.Day)

	payload = append(payload, cmd.HourUtc)

	payload = append(payload, cmd.MinuteUtc)

	payload = append(payload, cmd.SecondUtc)

	return
}
