// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package alarmv2

import (
	"encoding/gob"
	"errors"

	"github.com/opendoor-labs/gozw/cc"
)

const CommandReport cc.CommandID = 0x05

func init() {
	gob.Register(Report{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x71),
		Command:      cc.CommandID(0x05),
		Version:      2,
	}, NewReport)
}

func NewReport() cc.Command {
	return &Report{}
}

// <no value>
type Report struct {
	AlarmType byte

	AlarmLevel byte

	ZensorNetSourceNodeId byte

	ZwaveAlarmStatus byte

	ZwaveAlarmType byte

	ZwaveAlarmEvent byte

	NumberOfEventParameters byte

	EventParameter []byte
}

func (cmd Report) CommandClassID() cc.CommandClassID {
	return 0x71
}

func (cmd Report) CommandID() cc.CommandID {
	return CommandReport
}

func (cmd Report) CommandIDString() string {
	return "ALARM_REPORT"
}

func (cmd *Report) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	payload := make([]byte, len(data))
	copy(payload, data)

	if len(payload) < 2 {
		return errors.New("Payload length underflow")
	}

	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.AlarmType = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.AlarmLevel = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ZensorNetSourceNodeId = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ZwaveAlarmStatus = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ZwaveAlarmType = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ZwaveAlarmEvent = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.NumberOfEventParameters = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	{
		length := (payload[6+2]) & 0xFF
		cmd.EventParameter = payload[i : i+int(length)]
		i += int(length)
	}

	return nil
}

func (cmd *Report) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.AlarmType)

	payload = append(payload, cmd.AlarmLevel)

	payload = append(payload, cmd.ZensorNetSourceNodeId)

	payload = append(payload, cmd.ZwaveAlarmStatus)

	payload = append(payload, cmd.ZwaveAlarmType)

	payload = append(payload, cmd.ZwaveAlarmEvent)

	payload = append(payload, cmd.NumberOfEventParameters)

	if cmd.EventParameter != nil && len(cmd.EventParameter) > 0 {
		payload = append(payload, cmd.EventParameter...)
	}

	return
}
