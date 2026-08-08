package defs

type BinaryPV byte

const (
	BinaryInactive BinaryPV = 0x00
	BinaryActive   BinaryPV = 0x01
)

// Binary constants.
const (
	MinBinaryPV byte = 0x00 // PV
	MaxBinaryPV byte = 0x01 // for validating incoming values
	BinaryNull  byte = 0xFF // local sentinel for unavailable PV
)
