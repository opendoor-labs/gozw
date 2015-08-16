// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package scenecontrollerconf

import "errors"

// <no value>

type SceneControllerConfReport struct {
	GroupId byte

	SceneId byte

	DimmingDuration byte
}

func (cmd *SceneControllerConfReport) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.GroupId = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.SceneId = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.DimmingDuration = payload[i]
	i++

	return nil
}

func (cmd *SceneControllerConfReport) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.GroupId)

	payload = append(payload, cmd.SceneId)

	payload = append(payload, cmd.DimmingDuration)

	return
}
