package defs

// Lighting constants.
const (
	LightsNone uint16 = 0x00
	LightsFadeTo uint16 = 0x01
	LightsRampTo uint16 = 0x02
	LightsStepUp uint16 = 0x03
	LightsStepDown uint16 = 0x04
	LightsStepOn uint16 = 0x05
	LightsStepOff uint16 = 0x06
	LightsWarn uint16 = 0x07
	LightsWarnOff uint16 = 0x08
	LightsWarnRelinquish uint16 = 0x09
	LightsStop uint16 = 0x0A
	MaxLightingOperation uint16 = 0x0B
	LightsProprietaryFirst uint16 = 0x100
	LightsProprietaryLast uint16 = 0xFFFF
)

// Lighting constants.
const (
	LightingIdle byte = 0x00
	LightingFadeActive byte = 0x01
	LightingRampActive byte = 0x02
	LightingNotControlled byte = 0x03
	LightingOther byte = 0x04
	MaxLightingInProgress byte = 0x05
)

// Lighting constants.
const (
	LightingTransitionIdle byte = 0x00
	LightingTransitionFade byte = 0x01
	LightingTransitionRamp byte = 0x02
	MaxLightingTransition byte = 0x03
	LightingTransitionProprietaryFirst byte = 0x40
	LightingTransitionProprietaryLast byte = 0xFF
)
