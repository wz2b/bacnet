package defs

type RestartReason byte

// Restart constants.
const (
	RestartReasonUnknown           RestartReason = 0x00
	RestartReasonColdStart         RestartReason = 0x01
	RestartReasonWarmStart         RestartReason = 0x02
	RestartReasonDetectedPowerLost RestartReason = 0x03
	RestartReasonDetectedPowerOff  RestartReason = 0x04
	RestartReasonHardwareWatchdog  RestartReason = 0x05
	RestartReasonSoftwareWatchdog  RestartReason = 0x06
	RestartReasonSuspended         RestartReason = 0x07
	RestartReasonProprietaryMin    RestartReason = 0x40
	RestartReasonProprietaryMax    RestartReason = 0xFF
)
