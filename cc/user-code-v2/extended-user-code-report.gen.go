// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package usercodev2

import (
	"encoding/binary"
	"encoding/gob"
	"errors"

	"github.com/opendoor-labs/gozw/cc"
)

const CommandExtendedUserCodeReport cc.CommandID = 0x0D

func init() {
	gob.Register(ExtendedUserCodeReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x63),
		Command:      cc.CommandID(0x0D),
		Version:      2,
	}, NewExtendedUserCodeReport)
}

func NewExtendedUserCodeReport() cc.Command {
	return &ExtendedUserCodeReport{}
}

// <no value>
type ExtendedUserCodeReport struct {
	NumberOfUserCodes byte

	NextUserIdentifier uint16

	Vg1 []ExtendedUserCodeReportVg1
}

type ExtendedUserCodeReportVg1 struct {
	UserIdentifier uint16

	UserIdStatus byte

	Properties1 struct {
		UserCodeLength byte
	}

	UserCode []byte
}

func (cmd ExtendedUserCodeReport) CommandClassID() cc.CommandClassID {
	return 0x63
}

func (cmd ExtendedUserCodeReport) CommandID() cc.CommandID {
	return CommandExtendedUserCodeReport
}

func (cmd ExtendedUserCodeReport) CommandIDString() string {
	return "EXTENDED_USER_CODE_REPORT"
}

func (cmd *ExtendedUserCodeReport) UnmarshalBinary(data []byte) error {
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

	cmd.NumberOfUserCodes = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.NextUserIdentifier = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	for i < len(payload) {

		vg1 := ExtendedUserCodeReportVg1{}

		if len(payload) <= i {
			return errors.New("slice index out of bounds")
		}

		userIdentifier := binary.BigEndian.Uint16(payload[i : i+2])
		i += 2

		if len(payload) <= i {
			return errors.New("slice index out of bounds")
		}

		userIdStatus := payload[i]
		i++

		if len(payload) <= i {
			return errors.New("slice index out of bounds")
		}

		vg1.Properties1.UserCodeLength = (payload[i] & 0x0F)

		i += 1

		if len(payload) <= i {
			return errors.New("slice index out of bounds")
		}

		length := (payload[2+2]) & 0x0F
		userCode := payload[i : i+int(length)]
		i += int(length)

		vg1.UserIdentifier = userIdentifier

		vg1.UserIdStatus = userIdStatus

		// struct byte fields are assigned to the variant group when computed

		vg1.UserCode = userCode

		cmd.Vg1 = append(cmd.Vg1, vg1)
	}

	return nil
}

func (cmd *ExtendedUserCodeReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.NumberOfUserCodes)

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.NextUserIdentifier)
		payload = append(payload, buf...)
	}

	for _, vg := range cmd.Vg1 {

		{
			buf := make([]byte, 2)
			binary.BigEndian.PutUint16(buf, vg.UserIdentifier)
			payload = append(payload, buf...)
		}

		payload = append(payload, vg.UserIdStatus)

		{
			var val byte

			val |= (vg.Properties1.UserCodeLength) & byte(0x0F)

			payload = append(payload, val)
		}

		if vg.UserCode != nil && len(vg.UserCode) > 0 {
			payload = append(payload, vg.UserCode...)
		}

	}

	return
}
