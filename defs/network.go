package defs

// BACnet network and object identifier limits.
const (
	BroadcastNetwork uint16 = 0xFFFF

	InstanceBits        = 22
	MaxInstance  uint32 = 0x3FFFFF
	MaxObject    uint32 = 0x3FF

	HopCountDefault byte = 0xFF
)

// NetworkMessageType identifies a BACnet network-layer message.
//
// Values 0x80 through 0xFF identify vendor-proprietary network
// messages and are followed by a vendor identifier.
type NetworkMessageType byte

const (
	NetworkMessageWhoIsRouterToNetwork          NetworkMessageType = 0x00
	NetworkMessageIAmRouterToNetwork            NetworkMessageType = 0x01
	NetworkMessageICouldBeRouterToNetwork       NetworkMessageType = 0x02
	NetworkMessageRejectMessageToNetwork        NetworkMessageType = 0x03
	NetworkMessageRouterBusyToNetwork           NetworkMessageType = 0x04
	NetworkMessageRouterAvailableToNetwork      NetworkMessageType = 0x05
	NetworkMessageInitializeRoutingTable        NetworkMessageType = 0x06
	NetworkMessageInitializeRoutingTableACK     NetworkMessageType = 0x07
	NetworkMessageEstablishConnectionToNetwork  NetworkMessageType = 0x08
	NetworkMessageDisconnectConnectionToNetwork NetworkMessageType = 0x09
)

// NetworkRejectReason identifies the reason carried by a
// Reject-Message-To-Network network-layer message.
type NetworkRejectReason byte

const (
	NetworkRejectUnknownError       NetworkRejectReason = 0x00
	NetworkRejectNoRoute            NetworkRejectReason = 0x01
	NetworkRejectRouterBusy         NetworkRejectReason = 0x02
	NetworkRejectUnknownMessageType NetworkRejectReason = 0x03
	NetworkRejectMessageTooLong     NetworkRejectReason = 0x04
)
