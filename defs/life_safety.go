package defs

type LifeSafetyMode uint16

const (
	LifeSafetyModeOff                      LifeSafetyMode = 0x00
	LifeSafetyModeOn                       LifeSafetyMode = 0x01
	LifeSafetyModeTest                     LifeSafetyMode = 0x02
	LifeSafetyModeManned                   LifeSafetyMode = 0x03
	LifeSafetyModeUnmanned                 LifeSafetyMode = 0x04
	LifeSafetyModeArmed                    LifeSafetyMode = 0x05
	LifeSafetyModeDisarmed                 LifeSafetyMode = 0x06
	LifeSafetyModePrearmed                 LifeSafetyMode = 0x07
	LifeSafetyModeSlow                     LifeSafetyMode = 0x08
	LifeSafetyModeFast                     LifeSafetyMode = 0x09
	LifeSafetyModeDisconnected             LifeSafetyMode = 0x0A
	LifeSafetyModeEnabled                  LifeSafetyMode = 0x0B
	LifeSafetyModeDisabled                 LifeSafetyMode = 0x0C
	LifeSafetyModeAutomaticReleaseDisabled LifeSafetyMode = 0x0D
	LifeSafetyModeDefault                  LifeSafetyMode = 0x0E

	LifeSafetyModeProprietaryMin LifeSafetyMode = 0x100
	LifeSafetyModeProprietaryMax LifeSafetyMode = 0xFFFF
)

type LifeSafetyOperation uint16

const (
	LifeSafetyOpNone             LifeSafetyOperation = 0x00
	LifeSafetyOpSilence          LifeSafetyOperation = 0x01
	LifeSafetyOpSilenceAudible   LifeSafetyOperation = 0x02
	LifeSafetyOpSilenceVisual    LifeSafetyOperation = 0x03
	LifeSafetyOpReset            LifeSafetyOperation = 0x04
	LifeSafetyOpResetAlarm       LifeSafetyOperation = 0x05
	LifeSafetyOpResetFault       LifeSafetyOperation = 0x06
	LifeSafetyOpUnsilence        LifeSafetyOperation = 0x07
	LifeSafetyOpUnsilenceAudible LifeSafetyOperation = 0x08
	LifeSafetyOpUnsilenceVisual  LifeSafetyOperation = 0x09

	LifeSafetyOpProprietaryMin LifeSafetyOperation = 0x40
	LifeSafetyOpProprietaryMax LifeSafetyOperation = 0xFFFF
)

type LifeSafetyState uint16

const (
	LifeSafetyStateQuiet           LifeSafetyState = 0x00
	LifeSafetyStatePreAlarm        LifeSafetyState = 0x01
	LifeSafetyStateAlarm           LifeSafetyState = 0x02
	LifeSafetyStateFault           LifeSafetyState = 0x03
	LifeSafetyStateFaultPreAlarm   LifeSafetyState = 0x04
	LifeSafetyStateFaultAlarm      LifeSafetyState = 0x05
	LifeSafetyStateNotReady        LifeSafetyState = 0x06
	LifeSafetyStateActive          LifeSafetyState = 0x07
	LifeSafetyStateTamper          LifeSafetyState = 0x08
	LifeSafetyStateTestAlarm       LifeSafetyState = 0x09
	LifeSafetyStateTestActive      LifeSafetyState = 0x0A
	LifeSafetyStateTestFault       LifeSafetyState = 0x0B
	LifeSafetyStateTestFaultAlarm  LifeSafetyState = 0x0C
	LifeSafetyStateHoldup          LifeSafetyState = 0x0D
	LifeSafetyStateDuress          LifeSafetyState = 0x0E
	LifeSafetyStateTamperAlarm     LifeSafetyState = 0x0F
	LifeSafetyStateAbnormal        LifeSafetyState = 0x10
	LifeSafetyStateEmergencyPower  LifeSafetyState = 0x11
	LifeSafetyStateDelayed         LifeSafetyState = 0x12
	LifeSafetyStateBlocked         LifeSafetyState = 0x13
	LifeSafetyStateLocalAlarm      LifeSafetyState = 0x14
	LifeSafetyStateGeneralAlarm    LifeSafetyState = 0x15
	LifeSafetyStateSupervisory     LifeSafetyState = 0x16
	LifeSafetyStateTestSupervisory LifeSafetyState = 0x17

	LifeSafetyStateProprietaryMin LifeSafetyState = 0x100
	LifeSafetyStateProprietaryMax LifeSafetyState = 0xFFFF
)
