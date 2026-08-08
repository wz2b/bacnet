package defs

// Program constants.
const (
	ProgramRequestReady byte = 0x00
	ProgramRequestLoad byte = 0x01
	ProgramRequestRun byte = 0x02
	ProgramRequestHalt byte = 0x03
	ProgramRequestRestart byte = 0x04
	ProgramRequestUnload byte = 0x05
)

// Program constants.
const (
	ProgramStateIdle byte = 0x00
	ProgramStateLoading byte = 0x01
	ProgramStateRunning byte = 0x02
	ProgramStateWaiting byte = 0x03
	ProgramStateHalted byte = 0x04
	ProgramStateUnloading byte = 0x05
)

// Program constants.
const (
	ProgramErrorNormal uint16 = 0x00
	ProgramErrorLoadFailed uint16 = 0x01
	ProgramErrorInternal uint16 = 0x02
	ProgramErrorProgram uint16 = 0x03
	ProgramErrorOther uint16 = 0x04
	ProgramErrorProprietaryMin uint16 = 0x40
	ProgramErrorProprietaryMax uint16 = 0xFFFF
)
