package defs

// Restart constants.
const (
	RestartReasonUnknown byte = 0x00
	RestartReasonColdStart byte = 0x01
	RestartReasonWarmStart byte = 0x02
	RestartReasonDetectedPowerLost byte = 0x03
	RestartReasonDetectedPowerOff byte = 0x04
	RestartReasonHardwareWatchdog byte = 0x05
	RestartReasonSoftwareWatchdog byte = 0x06
	RestartReasonSuspended byte = 0x07
	RestartReasonProprietaryMin byte = 0x40
	RestartReasonProprietaryMax byte = 0xFF
)
