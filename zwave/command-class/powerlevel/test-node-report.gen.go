// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package powerlevel

import (
	"encoding/binary"
	"errors"
)

// <no value>

type PowerlevelTestNodeReport struct {
	TestNodeid byte

	StatusOfOperation byte

	TestFrameCount uint16
}

func (cmd *PowerlevelTestNodeReport) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.TestNodeid = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StatusOfOperation = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.TestFrameCount = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	return nil
}

func (cmd *PowerlevelTestNodeReport) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.TestNodeid)

	payload = append(payload, cmd.StatusOfOperation)

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.TestFrameCount)
		payload = append(payload, buf...)
	}

	return
}
