// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package sceneactivation

import "errors"

// <no value>

type SceneActivationSet struct {
	SceneId byte

	DimmingDuration byte
}

func (cmd *SceneActivationSet) UnmarshalBinary(payload []byte) error {
	i := 0

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

func (cmd *SceneActivationSet) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.SceneId)

	payload = append(payload, cmd.DimmingDuration)

	return
}
