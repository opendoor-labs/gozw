// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package configuration

import (
	"encoding/gob"
	"errors"

	"github.com/opendoor-labs/gozw/cc"
)

const CommandSet cc.CommandID = 0x04

func init() {
	gob.Register(Set{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x70),
		Command:      cc.CommandID(0x04),
		Version:      1,
	}, NewSet)
}

func NewSet() cc.Command {
	return &Set{}
}

// <no value>
type Set struct {
	ParameterNumber byte

	Level struct {
		Size byte

		Default bool
	}

	ConfigurationValue []byte
}

func (cmd Set) CommandClassID() cc.CommandClassID {
	return 0x70
}

func (cmd Set) CommandID() cc.CommandID {
	return CommandSet
}

func (cmd Set) CommandIDString() string {
	return "CONFIGURATION_SET"
}

func (cmd *Set) UnmarshalBinary(data []byte) error {
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

	cmd.ParameterNumber = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Level.Size = (payload[i] & 0x07)

	cmd.Level.Default = payload[i]&0x80 == 0x80

	i += 1

	if len(payload) <= i {
		return nil // field is optional
	}

	{
		length := (payload[1+2]) & 0x07
		cmd.ConfigurationValue = payload[i : i+int(length)]
		i += int(length)
	}

	return nil
}

func (cmd *Set) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.ParameterNumber)

	{
		var val byte

		val |= (cmd.Level.Size) & byte(0x07)

		if cmd.Level.Default {
			val |= byte(0x80) // flip bits on
		} else {
			val &= ^byte(0x80) // flip bits off
		}

		payload = append(payload, val)
	}

	if cmd.ConfigurationValue != nil && len(cmd.ConfigurationValue) > 0 {
		payload = append(payload, cmd.ConfigurationValue...)
	}

	return
}
