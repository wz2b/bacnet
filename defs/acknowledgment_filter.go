package defs

type AcknowledgementFilterType byte

// Acknowledgment filter constants.
const (
	AcknowledgmentFilterAll      AcknowledgementFilterType = 0x00
	AcknowledgmentFilterAcked    AcknowledgementFilterType = 0x01
	AcknowledgmentFilterNotAcked byte                      = 0x02
)
