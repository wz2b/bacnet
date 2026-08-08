package defs

type PolarityType byte

// Polarity constants.
const (
	PolarityNormal  PolarityType = 0x00
	PolarityReverse PolarityType = 0x01
	MaxPolarity     PolarityType = 0x02
)
