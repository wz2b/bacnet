package defs

type PropertyState byte

// Property states constants.
const (
	PropertyStateBooleanValue            PropertyState = 0x00
	PropertyStateBinaryValue             PropertyState = 0x01
	PropertyStateEventType               PropertyState = 0x02
	PropertyStatePolarity                PropertyState = 0x03
	PropertyStateProgramChange           PropertyState = 0x04
	PropertyStateProgramState            PropertyState = 0x05
	PropertyStateReasonForHalt           PropertyState = 0x06
	PropertyStateReliability             PropertyState = 0x07
	PropertyStateEventState              PropertyState = 0x08
	PropertyStateSystemStatus            PropertyState = 0x09
	PropertyStateUnits                   PropertyState = 0x0A
	PropertyStateUnsignedValue           PropertyState = 0x0B
	PropertyStateLifeSafetyMode          PropertyState = 0x0C
	PropertyStateLifeSafetyState         PropertyState = 0x0D
	PropertyStateRestartReason           PropertyState = 0x0E
	PropertyStateDoorAlarmState          PropertyState = 0x0F
	PropertyStateAction                  PropertyState = 0x10
	PropertyStateDoorSecuredStatus       PropertyState = 0x11
	PropertyStateDoorStatus              PropertyState = 0x12
	PropertyStateDoorValue               PropertyState = 0x13
	PropertyStateFileAccessMethod        PropertyState = 0x14
	PropertyStateLockStatus              PropertyState = 0x15
	PropertyStateLifeSafetyOperation     PropertyState = 0x16
	PropertyStateMaintenance             PropertyState = 0x17
	PropertyStateNodeType                PropertyState = 0x18
	PropertyStateNotifyType              PropertyState = 0x19
	PropertyStateSecurityLevel           PropertyState = 0x1A
	PropertyStateShedState               PropertyState = 0x1B
	PropertyStateSilencedState           PropertyState = 0x1C
	PropertyStateAccessEvent             PropertyState = 0x1E
	PropertyStateZoneOccupancyState      PropertyState = 0x1F
	PropertyStateAccessCredDisableReason PropertyState = 0x20
	PropertyStateAccessCredDisable       PropertyState = 0x21
	PropertyStateAuthenticationStatus    PropertyState = 0x22
)
