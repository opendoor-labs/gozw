// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package security

import (
	"encoding/gob"
	"errors"
	"fmt"

	"github.com/gozwave/gozw/cc"
)

const CommandCommandsSupportedReport cc.CommandID = 0x03

func init() {
	gob.Register(CommandsSupportedReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x98),
		Command:      cc.CommandID(0x03),
		Version:      1,
	}, NewCommandsSupportedReport)
}

func NewCommandsSupportedReport() cc.Command {
	return &CommandsSupportedReport{}
}

// <no value>
type CommandsSupportedReport struct {
	ReportsToFollow byte

	CommandClassSupport []byte

	CommandClassControl []byte
}

func (cmd CommandsSupportedReport) CommandClassID() cc.CommandClassID {
	return 0x98
}

func (cmd CommandsSupportedReport) CommandID() cc.CommandID {
	return CommandCommandsSupportedReport
}

func (cmd CommandsSupportedReport) CommandIDString() string {
	return "SECURITY_COMMANDS_SUPPORTED_REPORT"
}

func (cmd *CommandsSupportedReport) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	payload := make([]byte, len(data))
	copy(payload, data)

	if len(payload) < 2 {
		return errors.New("Payload length underflow")
	}

	i := 2

	if len(payload) <= i {
		return fmt.Errorf("slice index out of bounds (.ReportsToFollow) %d<=%d", len(payload), i)
	}

	cmd.ReportsToFollow = payload[i]
	i++

	if len(payload) <= i {
		return fmt.Errorf("slice index out of bounds (.CommandClassSupport) %d<=%d", len(payload), i)
	}

	{
		fieldStart := i
		for ; i < len(payload) && payload[i] != 0xEF; i++ {
		}
		cmd.CommandClassSupport = payload[fieldStart:i]
	}

	if len(payload) <= i {
		return fmt.Errorf("slice index out of bounds (.CommandClassMark) %d<=%d", len(payload), i)
	}

	i += 1 // skipping MARKER
	if len(payload) <= i {
		return nil
	}

	if len(payload) <= i {
		return nil
	}

	cmd.CommandClassControl = payload[i:]

	return nil
}

func (cmd *CommandsSupportedReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

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
