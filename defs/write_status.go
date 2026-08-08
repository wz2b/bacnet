package defs

type WriteStatus byte

// Write status constants.
const (
	WriteStatusIdle       WriteStatus = 0x00
	WriteStatusInProgress WriteStatus = 0x01
	WriteStatusSuccessful WriteStatus = 0x02
	WriteStatusFailed     WriteStatus = 0x03
)
