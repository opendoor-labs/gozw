// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package metertblmonitorv2

import (
	"encoding/binary"
	"errors"
)

// <no value>

type MeterTblCurrentDataGet struct {
	DatasetRequested uint32
}

func (cmd *MeterTblCurrentDataGet) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.DatasetRequested = binary.BigEndian.Uint32(payload[i : i+3])
	i += 3

	return nil
}

func (cmd *MeterTblCurrentDataGet) MarshalBinary() (payload []byte, err error) {

	{
		buf := make([]byte, 4)
		binary.BigEndian.PutUint32(buf, cmd.DatasetRequested)
		if buf[0] != 0 {
			return nil, errors.New("BIT_24 value overflow")
		}
		payload = append(payload, buf[1:4]...)
	}

	return
}
