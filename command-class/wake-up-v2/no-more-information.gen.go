// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package wakeupv2

import "encoding/gob"

func init() {
	gob.Register(NoMoreInformation{})
}

// <no value>
type NoMoreInformation struct {
}

func (cmd NoMoreInformation) CommandClassID() byte {
	return 0x84
}

func (cmd NoMoreInformation) CommandID() byte {
	return byte(CommandNoMoreInformation)
}

func (cmd *NoMoreInformation) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	return nil
}

func (cmd *NoMoreInformation) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = cmd.CommandClassID()
	payload[1] = cmd.CommandID()

	return
}