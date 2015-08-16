// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package ratetblmonitor

import (
	"encoding/binary"
	"errors"
)

// <no value>

type RateTblSupportedReport struct {
	RatesSupported byte

	ParameterSetSupportedBitMask uint16
}

func (cmd *RateTblSupportedReport) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.RatesSupported = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ParameterSetSupportedBitMask = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	return nil
}

func (cmd *RateTblSupportedReport) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.RatesSupported)

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.ParameterSetSupportedBitMask)
		payload = append(payload, buf...)
	}

	return
}
