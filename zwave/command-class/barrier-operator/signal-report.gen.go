// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package barrieroperator

import "errors"

// <no value>

type BarrierOperatorSignalReport struct {
	SubsystemType byte

	SubsystemState byte
}

func (cmd *BarrierOperatorSignalReport) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.SubsystemType = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.SubsystemState = payload[i]
	i++

	return nil
}

func (cmd *BarrierOperatorSignalReport) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.SubsystemType)

	payload = append(payload, cmd.SubsystemState)

	return
}