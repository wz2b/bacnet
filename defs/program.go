package defs

type ProgramRequest byte

const (
	ProgramRequestReady   ProgramRequest = 0x00
	ProgramRequestLoad    ProgramRequest = 0x01
	ProgramRequestRun     ProgramRequest = 0x02
	ProgramRequestHalt    ProgramRequest = 0x03
	ProgramRequestRestart ProgramRequest = 0x04
	ProgramRequestUnload  ProgramRequest = 0x05
)

type ProgramState byte

const (
	ProgramStateIdle      ProgramState = 0x00
	ProgramStateLoading   ProgramState = 0x01
	ProgramStateRunning   ProgramState = 0x02
	ProgramStateWaiting   ProgramState = 0x03
	ProgramStateHalted    ProgramState = 0x04
	ProgramStateUnloading ProgramState = 0x05
)

type ProgramError uint16

const (
	ProgramErrorNormal     ProgramError = 0x00
	ProgramErrorLoadFailed ProgramError = 0x01
	ProgramErrorInternal   ProgramError = 0x02
	ProgramErrorProgram    ProgramError = 0x03
	ProgramErrorOther      ProgramError = 0x04

	ProgramErrorProprietaryMin ProgramError = 0x40
	ProgramErrorProprietaryMax ProgramError = 0xFFFF
)
