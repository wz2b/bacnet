package defs

type ShedState byte

// Shed state constants.
const (
	ShedInactive       ShedState = 0x00
	ShedRequestPending ShedState = 0x01
	ShedCompliant      ShedState = 0x02
	ShedNonCompliant   ShedState = 0x03
)
