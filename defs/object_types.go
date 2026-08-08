package defs

type ObjectType uint16

// Object types constants.
const (
	ObjectAnalogInput           ObjectType = 0x00
	ObjectAnalogOutput          ObjectType = 0x01
	ObjectAnalogValue           ObjectType = 0x02
	ObjectBinaryInput           ObjectType = 0x03
	ObjectBinaryOutput          ObjectType = 0x04
	ObjectBinaryValue           ObjectType = 0x05
	ObjectCalendar              ObjectType = 0x06
	ObjectCommand               ObjectType = 0x07
	ObjectDevice                ObjectType = 0x08
	ObjectEventEnrollment       ObjectType = 0x09
	ObjectFile                  ObjectType = 0x0A
	ObjectGroup                 ObjectType = 0x0B
	ObjectLoop                  ObjectType = 0x0C
	ObjectMultiStateInput       ObjectType = 0x0D
	ObjectMultiStateOutput      ObjectType = 0x0E
	ObjectNotificationClass     ObjectType = 0x0F
	ObjectProgram               ObjectType = 0x10
	ObjectSchedule              ObjectType = 0x11
	ObjectAveraging             ObjectType = 0x12
	ObjectMultiStateValue       ObjectType = 0x13
	ObjectTrendLog              ObjectType = 0x14
	ObjectLifeSafetyPoint       ObjectType = 0x15
	ObjectLifeSafetyZone        ObjectType = 0x16
	ObjectAccumulator           ObjectType = 0x17
	ObjectPulseConverter        ObjectType = 0x18
	ObjectEventLog              ObjectType = 0x19
	ObjectGlobalGroup           ObjectType = 0x1A
	ObjectTrendLogMultiple      ObjectType = 0x1B
	ObjectLoadControl           ObjectType = 0x1C
	ObjectStructuredView        ObjectType = 0x1D
	ObjectAccessDoor            ObjectType = 0x1E
	ObjectTimer                 ObjectType = 0x1F
	ObjectAccessCredential      ObjectType = 0x20
	ObjectAccessPoint           ObjectType = 0x21
	ObjectAccessRights          ObjectType = 0x22
	ObjectAccessUser            ObjectType = 0x23
	ObjectAccessZone            ObjectType = 0x24
	ObjectCredentialDataInput   ObjectType = 0x25
	ObjectNetworkSecurity       ObjectType = 0x26
	ObjectBitStringValue        ObjectType = 0x27
	ObjectCharacterStringValue  ObjectType = 0x28
	ObjectDatePatternValue      ObjectType = 0x29
	ObjectDateValue             ObjectType = 0x2A
	ObjectDatetimePatternValue  ObjectType = 0x2B
	ObjectDatetimeValue         ObjectType = 0x2C
	ObjectIntegerValue          ObjectType = 0x2D
	ObjectLargeAnalogValue      ObjectType = 0x2E
	ObjectOctetStringValue      ObjectType = 0x2F
	ObjectPositiveIntegerValue  ObjectType = 0x30
	ObjectTimePatternValue      ObjectType = 0x31
	ObjectTimeValue             ObjectType = 0x32
	ObjectNotificationForwarder ObjectType = 0x33
	ObjectAlertEnrollment       ObjectType = 0x34
	ObjectChannel               ObjectType = 0x35
	ObjectLightingOutput        ObjectType = 0x36
	ObjectBinaryLightingOutput  ObjectType = 0x37
	ObjectNetworkPort           ObjectType = 0x38

	ObjectProprietaryMin ObjectType = 0x80
	ObjectProprietaryMax ObjectType = 0x3FF
)
