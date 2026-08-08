package defs

type Reliability uint16

// Reliability constants.
const (
	ReliabilityNoFaultDetected      Reliability = 0x00
	ReliabilityNoSensor             Reliability = 0x01
	ReliabilityOverRange            Reliability = 0x02
	ReliabilityUnderRange           Reliability = 0x03
	ReliabilityOpenLoop             Reliability = 0x04
	ReliabilityShortedLoop          Reliability = 0x05
	ReliabilityNoOutput             Reliability = 0x06
	ReliabilityUnreliableOther      Reliability = 0x07
	ReliabilityProcessError         Reliability = 0x08
	ReliabilityMultiStateFault      Reliability = 0x09
	ReliabilityConfigurationError   Reliability = 0x0A
	ReliabilityMemberFault          Reliability = 0x0B
	ReliabilityCommunicationFailure Reliability = 0x0C
	ReliabilityTripped              Reliability = 0x0D
	ReliabilityProprietaryMin       Reliability = 0x40
	ReliabilityProprietaryMax       Reliability = 0xFFFF
)
