// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package cc

import "fmt"

type BasicDeviceType byte

func (b BasicDeviceType) String() string {
	if val, ok := BasicDeviceTypeNames[b]; ok {
		return val + fmt.Sprintf(" (0x%X)", byte(b))
	} else {
		return "Unknown" + fmt.Sprintf(" (0x%X)", byte(b))
	}
}

type GenericDeviceType byte

func (g GenericDeviceType) String() string {
	if val, ok := GenericDeviceTypeNames[g]; ok {
		return val + fmt.Sprintf(" (0x%X)", byte(g))
	} else {
		return "Unknown" + fmt.Sprintf(" (0x%X)", byte(g))
	}
}

type SpecificDeviceType byte

type DeviceType struct {
	BasicType    BasicDeviceType
	GenericType  GenericDeviceType
	SpecificType SpecificDeviceType
}

func getSpecificDeviceTypeName(genericType GenericDeviceType, specificType SpecificDeviceType) string {
	if val, ok := SpecificDeviceTypeNames[genericType][specificType]; ok {
		return val + fmt.Sprintf(" (0x%X)", byte(specificType))
	} else {
		return "Unknown" + fmt.Sprintf(" (0x%X)", byte(specificType))
	}
}

func (d DeviceType) String() string {
	return fmt.Sprintf("%s / %s / %s", d.BasicType, d.GenericType, d.SpecificDeviceTypeString())
}

func (d DeviceType) SpecificDeviceTypeString() string {
	return getSpecificDeviceTypeName(d.GenericType, d.SpecificType)
}

const (
	BasicTypeController       BasicDeviceType = 0x01
	BasicTypeRoutingSlave     BasicDeviceType = 0x04
	BasicTypeSlave            BasicDeviceType = 0x03
	BasicTypeStaticController BasicDeviceType = 0x02
)

var BasicDeviceTypeNames map[BasicDeviceType]string = map[BasicDeviceType]string{

	BasicTypeController:       "Controller",
	BasicTypeRoutingSlave:     "Routing Slave",
	BasicTypeSlave:            "Slave",
	BasicTypeStaticController: "Static Controller",
}

const (
	GenericTypeAvControlPoint     GenericDeviceType = 0x03
	GenericTypeDisplay            GenericDeviceType = 0x04
	GenericTypeEntryControl       GenericDeviceType = 0x40
	GenericTypeGenericController  GenericDeviceType = 0x01
	GenericTypeMeter              GenericDeviceType = 0x31
	GenericTypeMeterPulse         GenericDeviceType = 0x30
	GenericTypeNonInteroperable   GenericDeviceType = 0xFF
	GenericTypeRepeaterSlave      GenericDeviceType = 0x0F
	GenericTypeSecurityPanel      GenericDeviceType = 0x17
	GenericTypeSemiInteroperable  GenericDeviceType = 0x50
	GenericTypeSensorAlarm        GenericDeviceType = 0xA1
	GenericTypeSensorBinary       GenericDeviceType = 0x20
	GenericTypeSensorMultilevel   GenericDeviceType = 0x21
	GenericTypeStaticController   GenericDeviceType = 0x02
	GenericTypeSwitchBinary       GenericDeviceType = 0x10
	GenericTypeSwitchMultilevel   GenericDeviceType = 0x11
	GenericTypeSwitchRemote       GenericDeviceType = 0x12
	GenericTypeSwitchToggle       GenericDeviceType = 0x13
	GenericTypeThermostat         GenericDeviceType = 0x08
	GenericTypeVentilation        GenericDeviceType = 0x16
	GenericTypeWindowCovering     GenericDeviceType = 0x09
	GenericTypeZipNode            GenericDeviceType = 0x15
	GenericTypeWallController     GenericDeviceType = 0x18
	GenericTypeNetworkExtender    GenericDeviceType = 0x05
	GenericTypeAppliance          GenericDeviceType = 0x06
	GenericTypeSensorNotification GenericDeviceType = 0x07
)

var GenericDeviceTypeNames map[GenericDeviceType]string = map[GenericDeviceType]string{

	GenericTypeAvControlPoint:     "Av Control Point",
	GenericTypeDisplay:            "Display",
	GenericTypeEntryControl:       "Entry Control",
	GenericTypeGenericController:  "Generic Controller",
	GenericTypeMeter:              "Meter",
	GenericTypeMeterPulse:         "Meter Pulse",
	GenericTypeNonInteroperable:   "Non Interoperable",
	GenericTypeRepeaterSlave:      "Repeater Slave",
	GenericTypeSecurityPanel:      "Security Panel",
	GenericTypeSemiInteroperable:  "Semi Interoperable",
	GenericTypeSensorAlarm:        "Sensor Alarm",
	GenericTypeSensorBinary:       "Sensor Binary",
	GenericTypeSensorMultilevel:   "Sensor Multilevel",
	GenericTypeStaticController:   "Static Controller",
	GenericTypeSwitchBinary:       "Switch Binary",
	GenericTypeSwitchMultilevel:   "Switch Multilevel",
	GenericTypeSwitchRemote:       "Switch Remote",
	GenericTypeSwitchToggle:       "Switch Toggle",
	GenericTypeThermostat:         "Thermostat",
	GenericTypeVentilation:        "Ventilation",
	GenericTypeWindowCovering:     "Window Covering",
	GenericTypeZipNode:            "Zip Node",
	GenericTypeWallController:     "Wall Controller",
	GenericTypeNetworkExtender:    "Network Extender",
	GenericTypeAppliance:          "Appliance",
	GenericTypeSensorNotification: "Sensor Notification",
}

const (
	SpecificTypeNotUsed SpecificDeviceType = 0x00

	SpecificTypeDoorbell                     SpecificDeviceType = 0x12
	SpecificTypeSatelliteReceiver            SpecificDeviceType = 0x04
	SpecificTypeSatelliteReceiverV2          SpecificDeviceType = 0x11
	SpecificTypeSoundSwitch                  SpecificDeviceType = 0x01
	SpecificTypeSimpleDisplay                SpecificDeviceType = 0x01
	SpecificTypeDoorLock                     SpecificDeviceType = 0x01
	SpecificTypeAdvancedDoorLock             SpecificDeviceType = 0x02
	SpecificTypeSecureKeypadDoorLock         SpecificDeviceType = 0x03
	SpecificTypeSecureKeypadDoorLockDeadbolt SpecificDeviceType = 0x04
	SpecificTypeSecureDoor                   SpecificDeviceType = 0x05
	SpecificTypeSecureGate                   SpecificDeviceType = 0x06
	SpecificTypeSecureBarrierAddon           SpecificDeviceType = 0x07
	SpecificTypeSecureBarrierOpenOnly        SpecificDeviceType = 0x08
	SpecificTypeSecureBarrierCloseOnly       SpecificDeviceType = 0x09
	SpecificTypeSecureLockbox                SpecificDeviceType = 0x0A
	SpecificTypeSecureKeypad                 SpecificDeviceType = 0x0B
	SpecificTypePortableRemoteController     SpecificDeviceType = 0x01
	SpecificTypePortableSceneController      SpecificDeviceType = 0x02
	SpecificTypePortableInstallerTool        SpecificDeviceType = 0x03
	SpecificTypeRemoteControlAv              SpecificDeviceType = 0x04
	SpecificTypeRemoteControlSimple          SpecificDeviceType = 0x06
	SpecificTypeSimpleMeter                  SpecificDeviceType = 0x01
	SpecificTypeAdvEnergyControl             SpecificDeviceType = 0x02
	SpecificTypeWholeHomeMeterSimple         SpecificDeviceType = 0x03
	SpecificTypeRepeaterSlave                SpecificDeviceType = 0x01
	SpecificTypeVirtualNode                  SpecificDeviceType = 0x02
	SpecificTypeZonedSecurityPanel           SpecificDeviceType = 0x01
	SpecificTypeEnergyProduction             SpecificDeviceType = 0x01
	SpecificTypeAdvZensorNetAlarmSensor      SpecificDeviceType = 0x05
	SpecificTypeAdvZensorNetSmokeSensor      SpecificDeviceType = 0x0A
	SpecificTypeBasicRoutingAlarmSensor      SpecificDeviceType = 0x01
	SpecificTypeBasicRoutingSmokeSensor      SpecificDeviceType = 0x06
	SpecificTypeBasicZensorNetAlarmSensor    SpecificDeviceType = 0x03
	SpecificTypeBasicZensorNetSmokeSensor    SpecificDeviceType = 0x08
	SpecificTypeRoutingAlarmSensor           SpecificDeviceType = 0x02
	SpecificTypeRoutingSmokeSensor           SpecificDeviceType = 0x07
	SpecificTypeZensorNetAlarmSensor         SpecificDeviceType = 0x04
	SpecificTypeZensorNetSmokeSensor         SpecificDeviceType = 0x09
	SpecificTypeAlarmSensor                  SpecificDeviceType = 0x0B
	SpecificTypeRoutingSensorBinary          SpecificDeviceType = 0x01
	SpecificTypeRoutingSensorMultilevel      SpecificDeviceType = 0x01
	SpecificTypeChimneyFan                   SpecificDeviceType = 0x02
	SpecificTypePcController                 SpecificDeviceType = 0x01
	SpecificTypeSceneController              SpecificDeviceType = 0x02
	SpecificTypeStaticInstallerTool          SpecificDeviceType = 0x03
	SpecificTypeSetTopBox                    SpecificDeviceType = 0x04
	SpecificTypeSubSystemController          SpecificDeviceType = 0x05
	SpecificTypeTv                           SpecificDeviceType = 0x06
	SpecificTypeGateway                      SpecificDeviceType = 0x07
	SpecificTypePowerSwitchBinary            SpecificDeviceType = 0x01
	SpecificTypeSceneSwitchBinary            SpecificDeviceType = 0x03
	SpecificTypePowerStrip                   SpecificDeviceType = 0x04
	SpecificTypeSiren                        SpecificDeviceType = 0x05
	SpecificTypeValveOpenClose               SpecificDeviceType = 0x06
	SpecificTypeColorTunableBinary           SpecificDeviceType = 0x02
	SpecificTypeIrrigationController         SpecificDeviceType = 0x07
	SpecificTypeClassAMotorControl           SpecificDeviceType = 0x05
	SpecificTypeClassBMotorControl           SpecificDeviceType = 0x06
	SpecificTypeClassCMotorControl           SpecificDeviceType = 0x07
	SpecificTypeMotorMultiposition           SpecificDeviceType = 0x03
	SpecificTypePowerSwitchMultilevel        SpecificDeviceType = 0x01
	SpecificTypeSceneSwitchMultilevel        SpecificDeviceType = 0x04
	SpecificTypeFanSwitch                    SpecificDeviceType = 0x08
	SpecificTypeColorTunableMultilevel       SpecificDeviceType = 0x02
	SpecificTypeSwitchRemoteBinary           SpecificDeviceType = 0x01
	SpecificTypeSwitchRemoteMultilevel       SpecificDeviceType = 0x02
	SpecificTypeSwitchRemoteToggleBinary     SpecificDeviceType = 0x03
	SpecificTypeSwitchRemoteToggleMultilevel SpecificDeviceType = 0x04
	SpecificTypeSwitchToggleBinary           SpecificDeviceType = 0x01
	SpecificTypeSwitchToggleMultilevel       SpecificDeviceType = 0x02
	SpecificTypeSetbackScheduleThermostat    SpecificDeviceType = 0x03
	SpecificTypeSetbackThermostat            SpecificDeviceType = 0x05
	SpecificTypeSetpointThermostat           SpecificDeviceType = 0x04
	SpecificTypeThermostatGeneral            SpecificDeviceType = 0x02
	SpecificTypeThermostatGeneralV2          SpecificDeviceType = 0x06
	SpecificTypeThermostatHeating            SpecificDeviceType = 0x01
	SpecificTypeResidentialHrv               SpecificDeviceType = 0x01
	SpecificTypeSimpleWindowCovering         SpecificDeviceType = 0x01
	SpecificTypeZipAdvNode                   SpecificDeviceType = 0x02
	SpecificTypeZipTunNode                   SpecificDeviceType = 0x01
	SpecificTypeBasicWallController          SpecificDeviceType = 0x01
	SpecificTypeSecureExtender               SpecificDeviceType = 0x01
	SpecificTypeGeneralAppliance             SpecificDeviceType = 0x01
	SpecificTypeKitchenAppliance             SpecificDeviceType = 0x02
	SpecificTypeLaundryAppliance             SpecificDeviceType = 0x03
	SpecificTypeNotificationSensor           SpecificDeviceType = 0x01
)

var SpecificDeviceTypeNames map[GenericDeviceType]map[SpecificDeviceType]string = map[GenericDeviceType]map[SpecificDeviceType]string{
	GenericTypeAvControlPoint: map[SpecificDeviceType]string{
		SpecificTypeNotUsed:             "Not Used",
		SpecificTypeDoorbell:            "Doorbell",
		SpecificTypeSatelliteReceiver:   "Satellite Receiver",
		SpecificTypeSatelliteReceiverV2: "Satellite Receiver V2",
		SpecificTypeSoundSwitch:         "Sound Switch",
	},
	GenericTypeDisplay: map[SpecificDeviceType]string{
		SpecificTypeNotUsed:       "Not Used",
		SpecificTypeSimpleDisplay: "Simple Display",
	},
	GenericTypeEntryControl: map[SpecificDeviceType]string{
		SpecificTypeNotUsed:                      "Not Used",
		SpecificTypeDoorLock:                     "Door Lock",
		SpecificTypeAdvancedDoorLock:             "Advanced Door Lock",
		SpecificTypeSecureKeypadDoorLock:         "Secure Keypad Door Lock",
		SpecificTypeSecureKeypadDoorLockDeadbolt: "Door Lock Keypad Deadbolt",
		SpecificTypeSecureDoor:                   "Secure Door",
		SpecificTypeSecureGate:                   "Secure Gate",
		SpecificTypeSecureBarrierAddon:           "Secure Barrier Addon",
		SpecificTypeSecureBarrierOpenOnly:        "Secure Barrier Open Only",
		SpecificTypeSecureBarrierCloseOnly:       "Secure Barrier Close Only",
		SpecificTypeSecureLockbox:                "Secure Lockbox",
		SpecificTypeSecureKeypad:                 "Secure Keypad",
	},
	GenericTypeGenericController: map[SpecificDeviceType]string{
		SpecificTypeNotUsed:                  "Not Used",
		SpecificTypePortableRemoteController: "Portable Remote Controller",
		SpecificTypePortableSceneController:  "Portable Scene Controller",
		SpecificTypePortableInstallerTool:    "Portable Installer Tool",
		SpecificTypeRemoteControlAv:          "Remote Control AV",
		SpecificTypeRemoteControlSimple:      "Remote Control Simple",
	},
	GenericTypeMeter: map[SpecificDeviceType]string{
		SpecificTypeNotUsed:              "Not Used",
		SpecificTypeSimpleMeter:          "Simple Meter",
		SpecificTypeAdvEnergyControl:     "Adv Energy Control",
		SpecificTypeWholeHomeMeterSimple: "Whole Home Meter Simple",
	},
	GenericTypeMeterPulse: map[SpecificDeviceType]string{
		SpecificTypeNotUsed: "Not Used",
	},
	GenericTypeNonInteroperable: map[SpecificDeviceType]string{
		SpecificTypeNotUsed: "Not Used",
	},
	GenericTypeRepeaterSlave: map[SpecificDeviceType]string{
		SpecificTypeNotUsed:       "Not Used",
		SpecificTypeRepeaterSlave: "Repeater Slave",
		SpecificTypeVirtualNode:   "Virtual Node",
	},
	GenericTypeSecurityPanel: map[SpecificDeviceType]string{
		SpecificTypeNotUsed:            "Not Used",
		SpecificTypeZonedSecurityPanel: "Zoned Security Panel",
	},
	GenericTypeSemiInteroperable: map[SpecificDeviceType]string{
		SpecificTypeNotUsed:          "Not Used",
		SpecificTypeEnergyProduction: "Energy Production",
	},
	GenericTypeSensorAlarm: map[SpecificDeviceType]string{
		SpecificTypeNotUsed:                   "Not Used",
		SpecificTypeAdvZensorNetAlarmSensor:   "Adv Zensor Net Alarm Sensor",
		SpecificTypeAdvZensorNetSmokeSensor:   "Adv Zensor Net Smoke Sensor",
		SpecificTypeBasicRoutingAlarmSensor:   "Basic Routing Alarm Sensor",
		SpecificTypeBasicRoutingSmokeSensor:   "Basic Routing Smoke Sensor",
		SpecificTypeBasicZensorNetAlarmSensor: "Basic Zensor Net Alarm Sensor",
		SpecificTypeBasicZensorNetSmokeSensor: "Basic Zensor Net Smoke Sensor",
		SpecificTypeRoutingAlarmSensor:        "Routing Alarm Sensor",
		SpecificTypeRoutingSmokeSensor:        "Routing Smoke Sensor",
		SpecificTypeZensorNetAlarmSensor:      "Zensor Net Alarm Sensor",
		SpecificTypeZensorNetSmokeSensor:      "Zensor Net Smoke Sensor",
		SpecificTypeAlarmSensor:               "Alarm Sensor",
	},
	GenericTypeSensorBinary: map[SpecificDeviceType]string{
		SpecificTypeNotUsed:             "Not Used",
		SpecificTypeRoutingSensorBinary: "Routing Sensor Binary",
	},
	GenericTypeSensorMultilevel: map[SpecificDeviceType]string{
		SpecificTypeNotUsed:                 "Not Used",
		SpecificTypeRoutingSensorMultilevel: "Routing Sensor Multilevel",
		SpecificTypeChimneyFan:              "Chimney Fan",
	},
	GenericTypeStaticController: map[SpecificDeviceType]string{
		SpecificTypeNotUsed:             "Not Used",
		SpecificTypePcController:        "Pc Controller",
		SpecificTypeSceneController:     "Scene Controller",
		SpecificTypeStaticInstallerTool: "Static Installer Tool",
		SpecificTypeSetTopBox:           "Set Top Box",
		SpecificTypeSubSystemController: "Sub System Controller",
		SpecificTypeTv:                  "TV",
		SpecificTypeGateway:             "Gateway",
	},
	GenericTypeSwitchBinary: map[SpecificDeviceType]string{
		SpecificTypeNotUsed:              "Not Used",
		SpecificTypePowerSwitchBinary:    "Power Switch Binary",
		SpecificTypeSceneSwitchBinary:    "Scene Switch Binary",
		SpecificTypePowerStrip:           "Power Strip",
		SpecificTypeSiren:                "Siren",
		SpecificTypeValveOpenClose:       "Valve Open/Close",
		SpecificTypeColorTunableBinary:   "Binary Tunable Color Light",
		SpecificTypeIrrigationController: "Irrigation Controller",
	},
	GenericTypeSwitchMultilevel: map[SpecificDeviceType]string{
		SpecificTypeNotUsed:                "Not Used",
		SpecificTypeClassAMotorControl:     "Class A Motor Control",
		SpecificTypeClassBMotorControl:     "Class B Motor Control",
		SpecificTypeClassCMotorControl:     "Class C Motor Control",
		SpecificTypeMotorMultiposition:     "Motor Multiposition",
		SpecificTypePowerSwitchMultilevel:  "Power Switch Multilevel",
		SpecificTypeSceneSwitchMultilevel:  "Scene Switch Multilevel",
		SpecificTypeFanSwitch:              "Fan Switch",
		SpecificTypeColorTunableMultilevel: "Multilevel Tunable Color Light",
	},
	GenericTypeSwitchRemote: map[SpecificDeviceType]string{
		SpecificTypeNotUsed:                      "Not Used",
		SpecificTypeSwitchRemoteBinary:           "Switch Remote Binary",
		SpecificTypeSwitchRemoteMultilevel:       "Switch Remote Multilevel",
		SpecificTypeSwitchRemoteToggleBinary:     "Switch Remote Toggle Binary",
		SpecificTypeSwitchRemoteToggleMultilevel: "Switch Remote Toggle Multilevel",
	},
	GenericTypeSwitchToggle: map[SpecificDeviceType]string{
		SpecificTypeNotUsed:                "Not Used",
		SpecificTypeSwitchToggleBinary:     "Switch Toggle Binary",
		SpecificTypeSwitchToggleMultilevel: "Switch Toggle Multilevel",
	},
	GenericTypeThermostat: map[SpecificDeviceType]string{
		SpecificTypeNotUsed:                   "Not Used",
		SpecificTypeSetbackScheduleThermostat: "Setback Schedule Thermostat",
		SpecificTypeSetbackThermostat:         "Setback Thermostat",
		SpecificTypeSetpointThermostat:        "Setpoint Thermostat",
		SpecificTypeThermostatGeneral:         "Thermostat General",
		SpecificTypeThermostatGeneralV2:       "Thermostat General V2",
		SpecificTypeThermostatHeating:         "Thermostat Heating",
	},
	GenericTypeVentilation: map[SpecificDeviceType]string{
		SpecificTypeNotUsed:        "Not Used",
		SpecificTypeResidentialHrv: "Residential Hrv",
	},
	GenericTypeWindowCovering: map[SpecificDeviceType]string{
		SpecificTypeNotUsed:              "Not Used",
		SpecificTypeSimpleWindowCovering: "Simple Window Covering",
	},
	GenericTypeZipNode: map[SpecificDeviceType]string{
		SpecificTypeNotUsed:    "Not Used",
		SpecificTypeZipAdvNode: "Zip Adv Node",
		SpecificTypeZipTunNode: "Zip Tun Node",
	},
	GenericTypeWallController: map[SpecificDeviceType]string{
		SpecificTypeNotUsed:             "Not Used",
		SpecificTypeBasicWallController: "Basic Wall Controller",
	},
	GenericTypeNetworkExtender: map[SpecificDeviceType]string{
		SpecificTypeNotUsed:        "Not Used",
		SpecificTypeSecureExtender: "Secure Extender",
	},
	GenericTypeAppliance: map[SpecificDeviceType]string{
		SpecificTypeNotUsed:          "Not Used",
		SpecificTypeGeneralAppliance: "General Appliance",
		SpecificTypeKitchenAppliance: "Kitchen Appliance",
		SpecificTypeLaundryAppliance: "Laundry Appliance",
	},
	GenericTypeSensorNotification: map[SpecificDeviceType]string{
		SpecificTypeNotUsed:            "Not Used",
		SpecificTypeNotificationSensor: "Notification Sensor",
	},
}
