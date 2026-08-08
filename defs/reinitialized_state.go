package defs

// Reinitialized state constants.
const (
	ReinitializedStateColdStart byte = 0x00
	ReinitializedStateWarmStart byte = 0x01
	ReinitializedStateStartBackup byte = 0x02
	ReinitializedStateEndBackup byte = 0x03
	ReinitializedStateStartRestore byte = 0x04
	ReinitializedStateEndRestore byte = 0x05
	ReinitializedStateAbortRestore byte = 0x06
	ReinitializedStateIdle byte = 0xFF
)
