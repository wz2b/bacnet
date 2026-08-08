package defs

type FileAccessType byte

// File access constants.
const (
	FileRecordAccess          FileAccessType = 0x00
	FileStreamAccess          FileAccessType = 0x01
	FileRecordAndStreamAccess FileAccessType = 0x02
)
