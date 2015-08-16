// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package versionv2

import "errors"

// <no value>

type VersionCommandClassReport struct {
	RequestedCommandClass byte

	CommandClassVersion byte
}

func (cmd *VersionCommandClassReport) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.RequestedCommandClass = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.CommandClassVersion = payload[i]
	i++

	return nil
}

func (cmd *VersionCommandClassReport) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.RequestedCommandClass)

	payload = append(payload, cmd.CommandClassVersion)

	return
}