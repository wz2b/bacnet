package defs

// Event States
const (
	EventStateNormal byte = 0x00
	EventStateFault byte = 0x01
	EventStateOffnormal byte = 0x02
	EventStateHighLimit byte = 0x03
	EventStateLowLimit byte = 0x04
)

// Event Enables
const (
	EventEnableToOffnormal byte = 0x01
	EventEnableToFault byte = 0x02
	EventEnableToNormal byte = 0x04
	EventLowLimitEnable byte = 0x01
	EventHighLimitEnable byte = 0x02
)

// Event Types
const (
	EventChangeOfBitstring uint16 = 0x00
	EventChangeOfState uint16 = 0x01
	EventChangeOfValue uint16 = 0x02
	EventCommandFailure uint16 = 0x03
	EventFloatingLimit uint16 = 0x04
	EventOutOfRange uint16 = 0x05
	EventChangeOfLifeSafety uint16 = 0x08
	EventExtended uint16 = 0x09
	EventBufferReady uint16 = 0x0A
	EventUnsignedRange uint16 = 0x0B
	EventProprietaryMin uint16 = 0x40
	EventProprietaryMax uint16 = 0xFFFF
)

// Event state Filters
const (
	EventStateFilterOffnormal byte = 0x00
	EventStateFilterFault byte = 0x01
	EventStateFilterNormal byte = 0x02
	EventStateFilterAll byte = 0x03
	EventStateFilterActive byte = 0x04
)

// Event transitions and bitmasks
const (
	TransitionToOffnormal byte = 0x00
	TransitionToFault byte = 0x01
	TransitionToNormal byte = 0x02
	MaxEventTransition byte = 0x03
	TransitionToOffnormalMasked byte = 0x01
	TransitionToFaultMasked byte = 0x02
	TransitionToNormalMasked byte = 0x04
)
