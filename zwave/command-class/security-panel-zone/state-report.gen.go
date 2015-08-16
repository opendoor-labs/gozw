// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package securitypanelzone

import "errors"

// <no value>

type SecurityPanelZoneStateReport struct {
	ZoneNumber byte

	ZoneState byte
}

func (cmd *SecurityPanelZoneStateReport) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ZoneNumber = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ZoneState = payload[i]
	i++

	return nil
}

func (cmd *SecurityPanelZoneStateReport) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.ZoneNumber)

	payload = append(payload, cmd.ZoneState)

	return
}
