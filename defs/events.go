package defs

type EventState byte

const (
	EventStateNormal    EventState = 0x00
	EventStateFault     EventState = 0x01
	EventStateOffnormal EventState = 0x02
	EventStateHighLimit EventState = 0x03
	EventStateLowLimit  EventState = 0x04
)

type EventEnable byte

const (
	EventEnableToOffnormal EventEnable = 0x01
	EventEnableToFault     EventEnable = 0x02
	EventEnableToNormal    EventEnable = 0x04
)

type LimitEnable byte

const (
	EventLowLimitEnable  LimitEnable = 0x01
	EventHighLimitEnable LimitEnable = 0x02
)

type EventType uint16

const (
	EventChangeOfBitstring  EventType = 0x00
	EventChangeOfState      EventType = 0x01
	EventChangeOfValue      EventType = 0x02
	EventCommandFailure     EventType = 0x03
	EventFloatingLimit      EventType = 0x04
	EventOutOfRange         EventType = 0x05
	EventChangeOfLifeSafety EventType = 0x08
	EventExtended           EventType = 0x09
	EventBufferReady        EventType = 0x0A
	EventUnsignedRange      EventType = 0x0B

	EventProprietaryMin EventType = 0x40
	EventProprietaryMax EventType = 0xFFFF
)

type EventStateFilter byte

const (
	EventStateFilterOffnormal EventStateFilter = 0x00
	EventStateFilterFault     EventStateFilter = 0x01
	EventStateFilterNormal    EventStateFilter = 0x02
	EventStateFilterAll       EventStateFilter = 0x03
	EventStateFilterActive    EventStateFilter = 0x04
)

type EventTransition byte

const (
	TransitionToOffnormal EventTransition = 0x00
	TransitionToFault     EventTransition = 0x01
	TransitionToNormal    EventTransition = 0x02
)

type EventTransitionMask byte

const (
	TransitionToOffnormalMasked EventTransitionMask = 0x01
	TransitionToFaultMasked     EventTransitionMask = 0x02
	TransitionToNormalMasked    EventTransitionMask = 0x04
)
