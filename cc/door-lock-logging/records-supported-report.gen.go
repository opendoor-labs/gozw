// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package doorlocklogging

import (
	"encoding/gob"
	"errors"
	"fmt"

	"github.com/gozwave/gozw/cc"
)

const CommandRecordsSupportedReport cc.CommandID = 0x02

func init() {
	gob.Register(RecordsSupportedReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x4C),
		Command:      cc.CommandID(0x02),
		Version:      1,
	}, NewRecordsSupportedReport)
}

func NewRecordsSupportedReport() cc.Command {
	return &RecordsSupportedReport{}
}

// <no value>
type RecordsSupportedReport struct {
	MaxRecordsStored byte
}

func (cmd RecordsSupportedReport) CommandClassID() cc.CommandClassID {
	return 0x4C
}

func (cmd RecordsSupportedReport) CommandID() cc.CommandID {
	return CommandRecordsSupportedReport
}

func (cmd RecordsSupportedReport) CommandIDString() string {
	return "DOOR_LOCK_LOGGING_RECORDS_SUPPORTED_REPORT"
}

func (cmd *RecordsSupportedReport) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	payload := make([]byte, len(data))
	copy(payload, data)

	if len(payload) < 2 {
		return errors.New("Payload length underflow")
	}

	i := 2

	if len(payload) <= i {
		return fmt.Errorf("slice index out of bounds (.MaxRecordsStored) %d<=%d", len(payload), i)
	}

	cmd.MaxRecordsStored = payload[i]
	i++

	return nil
}

func (cmd *RecordsSupportedReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.MaxRecordsStored)

	return
}
