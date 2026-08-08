package defs

// Reject reason constants.
const (
	RejectReasonOther uint16 = 0x00
	RejectReasonBufferOverflow uint16 = 0x01
	RejectReasonInconsistentParameters uint16 = 0x02
	RejectReasonInvalidParameterDataType uint16 = 0x03
	RejectReasonInvalidTag uint16 = 0x04
	RejectReasonMissingRequiredParameter uint16 = 0x05
	RejectReasonParameterOutOfRange uint16 = 0x06
	RejectReasonTooManyArguments uint16 = 0x07
	RejectReasonUndefinedEnumeration uint16 = 0x08
	RejectReasonUnrecognizedService uint16 = 0x09
	MaxRejectReason uint16 = 0x0A
	RejectReasonProprietaryFirst uint16 = 0x40
	RejectReasonProprietaryLast uint16 = 0xFFFF
)
