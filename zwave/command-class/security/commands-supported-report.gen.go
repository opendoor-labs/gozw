// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package security

import "errors"

// <no value>

type SecurityCommandsSupportedReport struct {
	ReportsToFollow byte

	CommandClassSupport []byte

	CommandClassControl []byte
}

func (cmd *SecurityCommandsSupportedReport) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ReportsToFollow = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	{
		markerIndex := i
		for ; markerIndex < len(payload) && payload[markerIndex] != 0xEF; markerIndex++ {
		}
		cmd.CommandClassSupport = payload[i:markerIndex]
	}

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	i += 1 // skipping MARKER

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.CommandClassControl = payload[i:]

	return nil
}

func (cmd *SecurityCommandsSupportedReport) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.ReportsToFollow)

	{
		if cmd.CommandClassSupport != nil && len(cmd.CommandClassSupport) > 0 {
			payload = append(payload, cmd.CommandClassSupport...)
		}
		payload = append(payload, 0xEF)
	}

	payload = append(payload, 0xEF) // marker

	payload = append(payload, cmd.CommandClassControl...)

	return
}
