package defs

// Binary constants.
const (
	MinBinaryPV    byte = 0x00 // PV
	BinaryInactive byte = 0x00
	BinaryActive   byte = 0x01
	MaxBinaryPV    byte = 0x01 // for validating incoming values
	BinaryNull     byte = 0xFF // local sentinel for unavailable PV
)
