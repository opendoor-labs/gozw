// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package thermostatfanmodev4

import (
	"encoding/gob"
	"errors"

	"github.com/opendoor-labs/gozw/cc"
)

const CommandSupportedReport cc.CommandID = 0x05

func init() {
	gob.Register(SupportedReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x44),
		Command:      cc.CommandID(0x05),
		Version:      4,
	}, NewSupportedReport)
}

func NewSupportedReport() cc.Command {
	return &SupportedReport{}
}

// <no value>
type SupportedReport struct {
	BitMask []byte
}

func (cmd SupportedReport) CommandClassID() cc.CommandClassID {
	return 0x44
}

func (cmd SupportedReport) CommandID() cc.CommandID {
	return CommandSupportedReport
}

func (cmd SupportedReport) CommandIDString() string {
	return "THERMOSTAT_FAN_MODE_SUPPORTED_REPORT"
}

func (cmd *SupportedReport) UnmarshalBinary(data []byte) error {
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

	cmd.BitMask = payload[i:]

	return nil
}

func (cmd *SupportedReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.BitMask...)

	return
}
