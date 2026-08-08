package defs

type LightingOperation uint16

const (
	LightingOperationNone           LightingOperation = 0x00
	LightingOperationFadeTo         LightingOperation = 0x01
	LightingOperationRampTo         LightingOperation = 0x02
	LightingOperationStepUp         LightingOperation = 0x03
	LightingOperationStepDown       LightingOperation = 0x04
	LightingOperationStepOn         LightingOperation = 0x05
	LightingOperationStepOff        LightingOperation = 0x06
	LightingOperationWarn           LightingOperation = 0x07
	LightingOperationWarnOff        LightingOperation = 0x08
	LightingOperationWarnRelinquish LightingOperation = 0x09
	LightingOperationStop           LightingOperation = 0x0A

	LightingOperationProprietaryFirst LightingOperation = 0x100
	LightingOperationProprietaryLast  LightingOperation = 0xFFFF
)

type LightingInProgress byte

const (
	LightingInProgressIdle          LightingInProgress = 0x00
	LightingInProgressFadeActive    LightingInProgress = 0x01
	LightingInProgressRampActive    LightingInProgress = 0x02
	LightingInProgressNotControlled LightingInProgress = 0x03
	LightingInProgressOther         LightingInProgress = 0x04
)

type LightingTransition byte

const (
	LightingTransitionIdle LightingTransition = 0x00
	LightingTransitionFade LightingTransition = 0x01
	LightingTransitionRamp LightingTransition = 0x02

	LightingTransitionProprietaryFirst LightingTransition = 0x40
	LightingTransitionProprietaryLast  LightingTransition = 0xFF
)
