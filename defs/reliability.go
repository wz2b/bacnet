package defs

// Reliability constants.
const (
	ReliabilityNoFaultDetected uint16 = 0x00
	ReliabilityNoSensor uint16 = 0x01
	ReliabilityOverRange uint16 = 0x02
	ReliabilityUnderRange uint16 = 0x03
	ReliabilityOpenLoop uint16 = 0x04
	ReliabilityShortedLoop uint16 = 0x05
	ReliabilityNoOutput uint16 = 0x06
	ReliabilityUnreliableOther uint16 = 0x07
	ReliabilityProcessError uint16 = 0x08
	ReliabilityMultiStateFault uint16 = 0x09
	ReliabilityConfigurationError uint16 = 0x0A
	ReliabilityMemberFault uint16 = 0x0B
	ReliabilityCommunicationFailure uint16 = 0x0C
	ReliabilityTripped uint16 = 0x0D
	ReliabilityProprietaryMin uint16 = 0x40
	ReliabilityProprietaryMax uint16 = 0xFFFF
)
