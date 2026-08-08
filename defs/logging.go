package defs

type LogStatus byte

const (
	LogStatusLogDisabled    LogStatus = 0x00
	LogStatusBufferPurged   LogStatus = 0x01
	LogStatusLogInterrupted LogStatus = 0x02
)

type LoggingType byte

const (
	LoggingTypePolled    LoggingType = 0x00
	LoggingTypeCOV       LoggingType = 0x01
	LoggingTypeTriggered LoggingType = 0x02
)
