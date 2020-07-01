// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package usercodev2

import (
	"encoding/gob"
	"errors"

	"github.com/opendoor-labs/gozw/cc"
)

const CommandMasterCodeSet cc.CommandID = 0x0E

func init() {
	gob.Register(MasterCodeSet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x63),
		Command:      cc.CommandID(0x0E),
		Version:      2,
	}, NewMasterCodeSet)
}

func NewMasterCodeSet() cc.Command {
	return &MasterCodeSet{}
}

// <no value>
type MasterCodeSet struct {
	Properties1 struct {
		MasterCodeLength byte
	}

	MasterCode []byte
}

func (cmd MasterCodeSet) CommandClassID() cc.CommandClassID {
	return 0x63
}

func (cmd MasterCodeSet) CommandID() cc.CommandID {
	return CommandMasterCodeSet
}

func (cmd MasterCodeSet) CommandIDString() string {
	return "MASTER_CODE_SET"
}

func (cmd *MasterCodeSet) UnmarshalBinary(data []byte) error {
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

	cmd.Properties1.MasterCodeLength = (payload[i] & 0x0F)

	i += 1

	if len(payload) <= i {
		return nil // field is optional
	}

	{
		length := (payload[0+2]) & 0x0F
		cmd.MasterCode = payload[i : i+int(length)]
		i += int(length)
	}

	return nil
}

func (cmd *MasterCodeSet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	{
		var val byte

		val |= (cmd.Properties1.MasterCodeLength) & byte(0x0F)

		payload = append(payload, val)
	}

	if cmd.MasterCode != nil && len(cmd.MasterCode) > 0 {
		payload = append(payload, cmd.MasterCode...)
	}

	return
}
