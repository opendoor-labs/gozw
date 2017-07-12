// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package security

import (
	"encoding/gob"
	"errors"
	"fmt"

	"github.com/gozwave/gozw/cc"
)

const CommandNonceReport cc.CommandID = 0x80

func init() {
	gob.Register(NonceReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x98),
		Command:      cc.CommandID(0x80),
		Version:      1,
	}, NewNonceReport)
}

func NewNonceReport() cc.Command {
	return &NonceReport{}
}

// <no value>
type NonceReport struct {
	NonceByte []byte
}

func (cmd NonceReport) CommandClassID() cc.CommandClassID {
	return 0x98
}

func (cmd NonceReport) CommandID() cc.CommandID {
	return CommandNonceReport
}

func (cmd NonceReport) CommandIDString() string {
	return "SECURITY_NONCE_REPORT"
}

func (cmd *NonceReport) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	payload := make([]byte, len(data))
	copy(payload, data)

	if len(payload) < 2 {
		return errors.New("Payload length underflow")
	}

	i := 2

	if len(payload) <= i {
		return fmt.Errorf("slice index out of bounds (.NonceByte) %d<=%d", len(payload), i)
	}

	cmd.NonceByte = payload[i : i+8]

	i += 8

	return nil
}

func (cmd *NonceReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	if paramLen := len(cmd.NonceByte); paramLen > 8 {
		return nil, errors.New("Length overflow in array parameter NonceByte")
	}

	payload = append(payload, cmd.NonceByte...)

	return
}
