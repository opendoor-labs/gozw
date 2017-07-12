// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package association

import (
	"encoding/gob"
	"errors"
	"fmt"

	"github.com/gozwave/gozw/cc"
)

const CommandGroupingsReport cc.CommandID = 0x06

func init() {
	gob.Register(GroupingsReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x85),
		Command:      cc.CommandID(0x06),
		Version:      1,
	}, NewGroupingsReport)
}

func NewGroupingsReport() cc.Command {
	return &GroupingsReport{}
}

// <no value>
type GroupingsReport struct {
	SupportedGroupings byte
}

func (cmd GroupingsReport) CommandClassID() cc.CommandClassID {
	return 0x85
}

func (cmd GroupingsReport) CommandID() cc.CommandID {
	return CommandGroupingsReport
}

func (cmd GroupingsReport) CommandIDString() string {
	return "ASSOCIATION_GROUPINGS_REPORT"
}

func (cmd *GroupingsReport) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	payload := make([]byte, len(data))
	copy(payload, data)

	if len(payload) < 2 {
		return errors.New("Payload length underflow")
	}

	i := 2

	if len(payload) <= i {
		return fmt.Errorf("slice index out of bounds (.SupportedGroupings) %d<=%d", len(payload), i)
	}

	cmd.SupportedGroupings = payload[i]
	i++

	return nil
}

func (cmd *GroupingsReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.SupportedGroupings)

	return
}
