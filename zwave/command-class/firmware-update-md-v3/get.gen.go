// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package firmwareupdatemdv3

import "errors"

// <no value>

type FirmwareUpdateMdGet struct {
	NumberOfReports byte

	Properties1 struct {
		ReportNumber1 byte

		Zero bool
	}

	ReportNumber2 byte
}

func (cmd *FirmwareUpdateMdGet) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.NumberOfReports = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.ReportNumber1 = (payload[i] & 0x7F)

	if payload[i]&0x80 == 0x80 {
		cmd.Properties1.Zero = true
	} else {
		cmd.Properties1.Zero = false
	}

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ReportNumber2 = payload[i]
	i++

	return nil
}

func (cmd *FirmwareUpdateMdGet) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.NumberOfReports)

	{
		var val byte

		val |= (cmd.Properties1.ReportNumber1) & byte(0x7F)

		if cmd.Properties1.Zero {
			val |= byte(0x80) // flip bits on
		} else {
			val &= ^byte(0x80) // flip bits off
		}

		payload = append(payload, val)
	}

	payload = append(payload, cmd.ReportNumber2)

	return
}