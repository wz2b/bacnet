package defs

// Network constants.
const (
	BroadcastNetwork uint16 = 0xFFFF
	MaxInstance int32 = 0x3FFFFF
)

// Network constants.
const (
	InstanceBits int = 22
	MaxObject uint32 = 0x3FF
)

// Network constants.
const (
	HopCountDefault byte = 0xFF
)

// Network constants.
const (
	NetworkMessageWhoIsRouterToNetwork uint16 = 0x00
	NetworkMessageIAmRouterToNetwork uint16 = 0x01
	NetworkMessageICouldBeRouterToNetwork uint16 = 0x02
	NetworkMessageRejectMessageToNetwork uint16 = 0x03
	NetworkMessageRouterBusyToNetwork uint16 = 0x04
	NetworkMessageRouterAvailableToNetwork uint16 = 0x05
	NetworkMessageInitializeRoutingTable uint16 = 0x06
	NetworkMessageInitializeRoutingTableACK uint16 = 0x07
	NetworkMessageEstablishConnectionToNetwork uint16 = 0x08
	NetworkMessageDisconnectConnectionToNetwork uint16 = 0x09
	NetworkMessageInvalid uint16 = 0x100
)

// Network constants.
const (
	NetworkRejectUnknownError byte = 0x00
	NetworkRejectNoRoute byte = 0x01
	NetworkRejectRouterBusy byte = 0x02
	NetworkRejectUnknownMessageType byte = 0x03
	NetworkRejectMessageTooLong byte = 0x04
)
