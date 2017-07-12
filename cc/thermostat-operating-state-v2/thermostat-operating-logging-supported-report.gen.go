// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package thermostatoperatingstatev2

import (
	"encoding/gob"
	"errors"
	"fmt"

	"github.com/gozwave/gozw/cc"
)

const CommandThermostatOperatingLoggingSupportedReport cc.CommandID = 0x04

func init() {
	gob.Register(ThermostatOperatingLoggingSupportedReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x42),
		Command:      cc.CommandID(0x04),
		Version:      2,
	}, NewThermostatOperatingLoggingSupportedReport)
}

func NewThermostatOperatingLoggingSupportedReport() cc.Command {
	return &ThermostatOperatingLoggingSupportedReport{}
}

// <no value>
type ThermostatOperatingLoggingSupportedReport struct {
	BitMask []byte
}

func (cmd ThermostatOperatingLoggingSupportedReport) CommandClassID() cc.CommandClassID {
	return 0x42
}

func (cmd ThermostatOperatingLoggingSupportedReport) CommandID() cc.CommandID {
	return CommandThermostatOperatingLoggingSupportedReport
}

func (cmd ThermostatOperatingLoggingSupportedReport) CommandIDString() string {
	return "THERMOSTAT_OPERATING_LOGGING_SUPPORTED_REPORT"
}

func (cmd *ThermostatOperatingLoggingSupportedReport) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	payload := make([]byte, len(data))
	copy(payload, data)

	if len(payload) < 2 {
		return errors.New("Payload length underflow")
	}

	i := 2

	if len(payload) <= i {
		return fmt.Errorf("slice index out of bounds (.BitMask) %d<=%d", len(payload), i)
	}

	cmd.BitMask = payload[i:]

	return nil
}

func (cmd *ThermostatOperatingLoggingSupportedReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.BitMask...)

	return
}
