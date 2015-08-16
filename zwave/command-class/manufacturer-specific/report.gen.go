// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package manufacturerspecific

import (
	"encoding/binary"
	"errors"
)

// <no value>

type ManufacturerSpecificReport struct {
	ManufacturerId uint16

	ProductTypeId uint16

	ProductId uint16
}

func (cmd *ManufacturerSpecificReport) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ManufacturerId = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ProductTypeId = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ProductId = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	return nil
}

func (cmd *ManufacturerSpecificReport) MarshalBinary() (payload []byte, err error) {

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.ManufacturerId)
		payload = append(payload, buf...)
	}

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.ProductTypeId)
		payload = append(payload, buf...)
	}

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.ProductId)
		payload = append(payload, buf...)
	}

	return
}
