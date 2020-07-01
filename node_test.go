package gozw

import (
	"testing"

	"github.com/opendoor-labs/gozw/cc"
	"github.com/stretchr/testify/assert"
)

func TestIsVersionCommandClass(t *testing.T) {
	assert.True(t, isVersionCommandClassID(cc.CommandClassID(0x86)))
	assert.True(t, isVersionCommandClassID(cc.Version))
	assert.True(t, isVersionCommandClassID(cc.VersionV2))
	assert.True(t, isVersionCommandClassID(cc.VersionV3))

	assert.False(t, isVersionCommandClassID(cc.CommandClassID(0x85)))
}

func TestIsSecurityCommandClass(t *testing.T) {
	assert.True(t, isSecurityCommandClassID(cc.CommandClassID(0x98)))
	assert.True(t, isSecurityCommandClassID(cc.CommandClassID(0x9F)))
	assert.True(t, isSecurityCommandClassID(cc.Security))
	assert.True(t, isSecurityCommandClassID(cc.Security2))

	assert.False(t, isSecurityCommandClassID(cc.CommandClassID(0x97)))
}

func TestCommandClassesInOrderToInterview(t *testing.T) {
	emptySet := cc.CommandClassSet{}
	assert.Len(t, commandClassesInOrderToInterview(emptySet), 0)

	noVersionOrSecurity := cc.CommandClassSet{
		// ccSupport defaults to insecure
		cc.Alarm:       &cc.CommandClassSupport{},
		cc.Antitheft:   &cc.CommandClassSupport{},
		cc.Association: &cc.CommandClassSupport{},
	}
	assert.Len(t, commandClassesInOrderToInterview(noVersionOrSecurity), 3,
		"mostly testing that we don't blow up")

	justVersion := cc.CommandClassSet{
		cc.Alarm:   &cc.CommandClassSupport{},
		cc.Version: &cc.CommandClassSupport{},
	}
	assert.Len(t, commandClassesInOrderToInterview(justVersion), 2,
		"we make no attempt to optimize the order in this case")

	justSecurity := cc.CommandClassSet{
		cc.Alarm:    &cc.CommandClassSupport{},
		cc.Security: &cc.CommandClassSupport{},
	}
	assert.Len(t, commandClassesInOrderToInterview(justSecurity), 2,
		"we make no attempt to optimize the order in this case")

	var inorder []cc.CommandClassID

	onlyVersionAndSecurity := cc.CommandClassSet{
		cc.Version:  &cc.CommandClassSupport{},
		cc.Security: &cc.CommandClassSupport{},
	}
	inorder = commandClassesInOrderToInterview(onlyVersionAndSecurity)
	assert.Len(t, inorder, 2,
		"50/50 shot at accidentally getting the order right")
	assert.Equal(t, cc.Security, inorder[0])
	assert.Equal(t, cc.Version, inorder[1])

	oneMoreThanSecurityAndVersion := cc.CommandClassSet{
		cc.Version:  &cc.CommandClassSupport{},
		cc.Security: &cc.CommandClassSupport{},
		cc.Alarm:    &cc.CommandClassSupport{},
	}
	inorder = commandClassesInOrderToInterview(oneMoreThanSecurityAndVersion)
	assert.Len(t, inorder, 3, "1/3 shot at accidentally getting the order right")
	assert.Equal(t, cc.Security, inorder[0])
	assert.Equal(t, cc.Version, inorder[1])
	assert.Equal(t, cc.Alarm, inorder[2])

	realWorldExample := cc.CommandClassSet{
		cc.AssociationGrpInfo:       &cc.CommandClassSupport{},
		cc.SensorMultilevel:         &cc.CommandClassSupport{},
		cc.Crc16Encap:               &cc.CommandClassSupport{},
		cc.ManufacturerSpecific:     &cc.CommandClassSupport{},
		cc.Battery:                  &cc.CommandClassSupport{},
		cc.ZwaveplusInfo:            &cc.CommandClassSupport{},
		cc.Powerlevel:               &cc.CommandClassSupport{},
		cc.VersionV2:                &cc.CommandClassSupport{},
		cc.Security:                 &cc.CommandClassSupport{},
		cc.ApplicationStatus:        &cc.CommandClassSupport{},
		cc.FirmwareUpdateMd:         &cc.CommandClassSupport{},
		cc.Basic:                    &cc.CommandClassSupport{Secure: true},
		cc.WakeUp:                   &cc.CommandClassSupport{Secure: true},
		cc.Alarm:                    &cc.CommandClassSupport{Secure: true},
		cc.Configuration:            &cc.CommandClassSupport{Secure: true},
		cc.SensorAlarm:              &cc.CommandClassSupport{Secure: true},
		cc.Association:              &cc.CommandClassSupport{Secure: true},
		cc.MultiInstanceAssociation: &cc.CommandClassSupport{Secure: true},
		cc.SensorBinary:             &cc.CommandClassSupport{Secure: true},
		cc.DeviceResetLocally:       &cc.CommandClassSupport{Secure: true},
	}
	inorder = commandClassesInOrderToInterview(realWorldExample)
	assert.Len(t, inorder, 20)
	assert.Equal(t, cc.Security, inorder[0])
	assert.Equal(t, cc.Version, inorder[1])

	find := func(sl []cc.CommandClassID, ccid cc.CommandClassID) int {
		for idx := range sl {
			if sl[idx] == ccid {
				return idx
			}
		}
		return -1
	}

	assert.True(t, find(inorder, cc.AssociationGrpInfo) < 11, "should be in insecure section")
	assert.True(t, find(inorder, cc.SensorMultilevel) < 11, "should be in insecure section")
	assert.True(t, find(inorder, cc.Crc16Encap) < 11, "should be in insecure section")
	assert.True(t, find(inorder, cc.ManufacturerSpecific) < 11, "should be in insecure section")
	assert.True(t, find(inorder, cc.Battery) < 11, "should be in insecure section")
	assert.True(t, find(inorder, cc.ZwaveplusInfo) < 11, "should be in insecure section")
	assert.True(t, find(inorder, cc.Powerlevel) < 11, "should be in insecure section")
	assert.True(t, find(inorder, cc.ApplicationStatus) < 11, "should be in insecure section")
	assert.True(t, find(inorder, cc.FirmwareUpdateMd) < 11, "should be in insecure section")

	assert.True(t, find(inorder, cc.Basic) > 10, "should be in secure section")
	assert.True(t, find(inorder, cc.WakeUp) > 10, "should be in secure section")
	assert.True(t, find(inorder, cc.Alarm) > 10, "should be in secure section")
	assert.True(t, find(inorder, cc.Configuration) > 10, "should be in secure section")
	assert.True(t, find(inorder, cc.SensorAlarm) > 10, "should be in secure section")
	assert.True(t, find(inorder, cc.Association) > 10, "should be in secure section")
	assert.True(t, find(inorder, cc.MultiInstanceAssociation) > 10, "should be in secure section")
	assert.True(t, find(inorder, cc.SensorBinary) > 10, "should be in secure section")
	assert.True(t, find(inorder, cc.DeviceResetLocally) > 10, "should be in secure section")

}
