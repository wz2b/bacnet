package defs

type RejectReason uint32

// Reject reason constants.
const (
	RejectReasonOther                    RejectReason = 0x00
	RejectReasonBufferOverflow           RejectReason = 0x01
	RejectReasonInconsistentParameters   RejectReason = 0x02
	RejectReasonInvalidParameterDataType RejectReason = 0x03
	RejectReasonInvalidTag               RejectReason = 0x04
	RejectReasonMissingRequiredParameter RejectReason = 0x05
	RejectReasonParameterOutOfRange      RejectReason = 0x06
	RejectReasonTooManyArguments         RejectReason = 0x07
	RejectReasonUndefinedEnumeration     RejectReason = 0x08
	RejectReasonUnrecognizedService      RejectReason = 0x09
	MaxRejectReason                      RejectReason = 0x0A
	RejectReasonProprietaryFirst         RejectReason = 0x40
	RejectReasonProprietaryLast          RejectReason = 0xFFFF
)
