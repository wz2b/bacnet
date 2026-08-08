package defs

type ReinitializedState byte

// Reinitialized state constants.
const (
	ReinitializedStateColdStart    ReinitializedState = 0x00
	ReinitializedStateWarmStart    ReinitializedState = 0x01
	ReinitializedStateStartBackup  ReinitializedState = 0x02
	ReinitializedStateEndBackup    ReinitializedState = 0x03
	ReinitializedStateStartRestore ReinitializedState = 0x04
	ReinitializedStateEndRestore   ReinitializedState = 0x05
	ReinitializedStateAbortRestore ReinitializedState = 0x06
	ReinitializedStateIdle         ReinitializedState = 0xFF
)
