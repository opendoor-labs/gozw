// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package appliance

import "errors"

// <no value>

type ApplianceReport struct {
	Properties1 struct {
		NoOfManufacturerDataFields byte

		ApplianceMode byte
	}

	ApplianceProgram byte

	ManufacturerData []byte
}

func (cmd *ApplianceReport) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.NoOfManufacturerDataFields = (payload[i] & 0xF0) >> 4

	cmd.Properties1.ApplianceMode = (payload[i] & 0x0F)

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ApplianceProgram = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ManufacturerData = payload[i : i+0]
	i += 0

	return nil
}

func (cmd *ApplianceReport) MarshalBinary() (payload []byte, err error) {

	{
		var val byte

		val |= (cmd.Properties1.NoOfManufacturerDataFields << byte(4)) & byte(0xF0)

		val |= (cmd.Properties1.ApplianceMode) & byte(0x0F)

		payload = append(payload, val)
	}

	payload = append(payload, cmd.ApplianceProgram)

	if cmd.ManufacturerData != nil && len(cmd.ManufacturerData) > 0 {
		payload = append(payload, cmd.ManufacturerData...)
	}

	return
}
