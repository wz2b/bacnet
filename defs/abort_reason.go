package defs

// Abort reason constants.
const (
	AbortReasonOther uint16 = 0x00
	AbortReasonBufferOverflow uint16 = 0x01
	AbortReasonInvalidAPDUInThisState uint16 = 0x02
	AbortReasonPreemptedByHigherPriorityTask uint16 = 0x03
	AbortReasonSegmentationNotSupported uint16 = 0x04
	MaxAbortReason uint16 = 0x05
	AbortReasonProprietaryFirst uint16 = 0x40
	AbortReasonProprietaryLast uint16 = 0xFFFF
)
