// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package associationcommandconfiguration

import "errors"

// <no value>

type CommandConfigurationSet struct {
	GroupingIdentifier byte

	NodeId byte

	CommandLength byte

	CommandClassIdentifier byte

	CommandIdentifier byte

	CommandByte []byte
}

func (cmd *CommandConfigurationSet) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.GroupingIdentifier = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.NodeId = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.CommandLength = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.CommandClassIdentifier = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.CommandIdentifier = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.CommandByte = payload[i:]

	return nil
}

func (cmd *CommandConfigurationSet) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.GroupingIdentifier)

	payload = append(payload, cmd.NodeId)

	payload = append(payload, cmd.CommandLength)

	payload = append(payload, cmd.CommandClassIdentifier)

	payload = append(payload, cmd.CommandIdentifier)

	payload = append(payload, cmd.CommandByte...)

	return
}