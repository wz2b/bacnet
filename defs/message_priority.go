package defs

type MessagePriority byte

// Message priority constants.
const (
	MessagePriorityNormal            MessagePriority = 0x00
	MessagePriorityUrgent            MessagePriority = 0x01
	MessagePriorityCriticalEquipment MessagePriority = 0x02
	MessagePriorityLifeSafety        MessagePriority = 0x03
)
