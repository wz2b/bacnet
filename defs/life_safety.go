package defs

// Life safety constants.
const (
	MinLifeSafetyMode uint16 = 0x00
	LifeSafetyModeOff uint16 = 0x00
	LifeSafetyModeOn uint16 = 0x01
	LifeSafetyModeTest uint16 = 0x02
	LifeSafetyModeManned uint16 = 0x03
	LifeSafetyModeUnmanned uint16 = 0x04
	LifeSafetyModeArmed uint16 = 0x05
	LifeSafetyModeDisarmed uint16 = 0x06
	LifeSafetyModePrearmed uint16 = 0x07
	LifeSafetyModeSlow uint16 = 0x08
	LifeSafetyModeFast uint16 = 0x09
	LifeSafetyModeDisconnected uint16 = 0x0A
	LifeSafetyModeEnabled uint16 = 0x0B
	LifeSafetyModeDisabled uint16 = 0x0C
	LifeSafetyModeAutomaticReleaseDisabled uint16 = 0x0D
	LifeSafetyModeDefault uint16 = 0x0E
	MaxLifeSafetyMode uint16 = 0x0F
	LifeSafetyModeProprietaryMin uint16 = 0x100
	LifeSafetyModeProprietaryMax uint16 = 0xFFFF
)

// Life safety constants.
const (
	LifeSafetyOpNone uint16 = 0x00
	LifeSafetyOpSilence uint16 = 0x01
	LifeSafetyOpSilenceAudible uint16 = 0x02
	LifeSafetyOpSilenceVisual uint16 = 0x03
	LifeSafetyOpReset uint16 = 0x04
	LifeSafetyOpResetAlarm uint16 = 0x05
	LifeSafetyOpResetFault uint16 = 0x06
	LifeSafetyOpUnsilence uint16 = 0x07
	LifeSafetyOpUnsilenceAudible uint16 = 0x08
	LifeSafetyOpUnsilenceVisual uint16 = 0x09
	LifeSafetyOpProprietaryMin uint16 = 0x40
	LifeSafetyOpProprietaryMax uint16 = 0xFFFF
)

// Life safety constants.
const (
	MinLifeSafetyState uint16 = 0x00
	LifeSafetyStateQuiet uint16 = 0x00
	LifeSafetyStatePreAlarm uint16 = 0x01
	LifeSafetyStateAlarm uint16 = 0x02
	LifeSafetyStateFault uint16 = 0x03
	LifeSafetyStateFaultPreAlarm uint16 = 0x04
	LifeSafetyStateFaultAlarm uint16 = 0x05
	LifeSafetyStateNotReady uint16 = 0x06
	LifeSafetyStateActive uint16 = 0x07
	LifeSafetyStateTamper uint16 = 0x08
	LifeSafetyStateTestAlarm uint16 = 0x09
	LifeSafetyStateTestActive uint16 = 0x0A
	LifeSafetyStateTestFault uint16 = 0x0B
	LifeSafetyStateTestFaultAlarm uint16 = 0x0C
	LifeSafetyStateHoldup uint16 = 0x0D
	LifeSafetyStateDuress uint16 = 0x0E
	LifeSafetyStateTamperAlarm uint16 = 0x0F
	LifeSafetyStateAbnormal uint16 = 0x10
	LifeSafetyStateEmergencyPower uint16 = 0x11
	LifeSafetyStateDelayed uint16 = 0x12
	LifeSafetyStateBlocked uint16 = 0x13
	LifeSafetyStateLocalAlarm uint16 = 0x14
	LifeSafetyStateGeneralAlarm uint16 = 0x15
	LifeSafetyStateSupervisory uint16 = 0x16
	LifeSafetyStateTestSupervisory uint16 = 0x17
	MaxLifeSafetyState uint16 = 0x18
	LifeSafetyStateProprietaryMin uint16 = 0x100
	LifeSafetyStateProprietaryMax uint16 = 0xFFFF
)
