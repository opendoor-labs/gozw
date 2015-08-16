// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package wakeupv2

import (
	"encoding/binary"
	"errors"
)

// <no value>

type WakeUpIntervalSet struct {
	Seconds uint32

	Nodeid byte
}

func (cmd *WakeUpIntervalSet) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Seconds = binary.BigEndian.Uint32(payload[i : i+3])
	i += 3

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Nodeid = payload[i]
	i++

	return nil
}

func (cmd *WakeUpIntervalSet) MarshalBinary() (payload []byte, err error) {

	{
		buf := make([]byte, 4)
		binary.BigEndian.PutUint32(buf, cmd.Seconds)
		if buf[0] != 0 {
			return nil, errors.New("BIT_24 value overflow")
		}
		payload = append(payload, buf[1:4]...)
	}

	payload = append(payload, cmd.Nodeid)

	return
}
