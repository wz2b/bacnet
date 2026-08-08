package bactypes

type TimestampType byte

const (
	TimestampTime     TimestampType = 0x00
	TimestampSequence TimestampType = 0x01
	TimestampDateTime TimestampType = 0x02
)

type Timestamp struct {
	Tag         TimestampType
	Time        Time
	DateTime    DateTime
	SequenceNum uint16
}
