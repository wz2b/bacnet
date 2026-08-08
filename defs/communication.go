package defs

type CommunicationState byte

// Communication constants.
const (
	CommunicationEnable            CommunicationState = 0x00
	CommunicationDisable           CommunicationState = 0x01
	CommunicationDisableInitiation CommunicationState = 0x02
	MaxCommunicationEnableDisable  CommunicationState = 0x03
)
