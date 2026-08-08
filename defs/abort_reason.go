package defs

type AbortReason uint16

// Abort reason constants.
const (
	AbortReasonOther                         AbortReason = 0x00
	AbortReasonBufferOverflow                AbortReason = 0x01
	AbortReasonInvalidAPDUInThisState        AbortReason = 0x02
	AbortReasonPreemptedByHigherPriorityTask AbortReason = 0x03
	AbortReasonSegmentationNotSupported      AbortReason = 0x04
	MaxAbortReason                           AbortReason = 0x05
	AbortReasonProprietaryFirst              AbortReason = 0x40
	AbortReasonProprietaryLast               AbortReason = 0xFFFF
)
