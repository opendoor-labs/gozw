// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package barrieroperator

import "errors"

// <no value>

type BarrierOperatorSignalSupportedReport struct {
	BitMask byte
}

func (cmd *BarrierOperatorSignalSupportedReport) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.BitMask = payload[i]
	i++

	return nil
}

func (cmd *BarrierOperatorSignalSupportedReport) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.BitMask)

	return
}