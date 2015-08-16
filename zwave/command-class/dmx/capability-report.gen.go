// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package dmx

import (
	"encoding/binary"
	"errors"
)

// <no value>

type DmxCapabilityReport struct {
	ChannelId byte

	PropertyId uint16

	DeviceChannels byte

	MaxChannels byte
}

func (cmd *DmxCapabilityReport) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ChannelId = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.PropertyId = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.DeviceChannels = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.MaxChannels = payload[i]
	i++

	return nil
}

func (cmd *DmxCapabilityReport) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.ChannelId)

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.PropertyId)
		payload = append(payload, buf...)
	}

	payload = append(payload, cmd.DeviceChannels)

	payload = append(payload, cmd.MaxChannels)

	return
}