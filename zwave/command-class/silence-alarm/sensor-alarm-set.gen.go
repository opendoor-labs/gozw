// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package silencealarm

import (
	"encoding/binary"
	"errors"
)

// <no value>

type SensorAlarmSet struct {
	Mode byte

	Seconds uint16

	NumberOfBitMasks byte

	BitMask []byte
}

func (cmd *SensorAlarmSet) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Mode = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Seconds = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.NumberOfBitMasks = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.BitMask = payload[i : i+2]
	i += 2

	return nil
}

func (cmd *SensorAlarmSet) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.Mode)

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.Seconds)
		payload = append(payload, buf...)
	}

	payload = append(payload, cmd.NumberOfBitMasks)

	if cmd.BitMask != nil && len(cmd.BitMask) > 0 {
		payload = append(payload, cmd.BitMask...)
	}

	return
}
