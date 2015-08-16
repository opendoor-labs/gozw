// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package multichannelv3

import "errors"

// <no value>

type MultiChannelEndPointFind struct {
	GenericDeviceClass byte

	SpecificDeviceClass byte
}

func (cmd *MultiChannelEndPointFind) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.GenericDeviceClass = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.SpecificDeviceClass = payload[i]
	i++

	return nil
}

func (cmd *MultiChannelEndPointFind) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.GenericDeviceClass)

	payload = append(payload, cmd.SpecificDeviceClass)

	return
}
