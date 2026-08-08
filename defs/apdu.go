package defs

type PDUType byte

// APDU constants.
const (
	PDUTypeConfirmedServiceRequest   PDUType = 0x00
	PDUTypeUnconfirmedServiceRequest PDUType = 0x10
	PDUTypeSimpleACK                 PDUType = 0x20
	PDUTypeComplexACK                PDUType = 0x30
	PDUuTypeSegmentACK               PDUType = 0x40
	PDUTypeError                     PDUType = 0x50
	PDUTypeReject                    PDUType = 0x60
	PDUTypeAbort                     PDUType = 0x70
)
