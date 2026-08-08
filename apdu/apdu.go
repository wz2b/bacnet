package apdu

import (
	"net"

	"github.com/wz2b/bacnet/defs"
	"github.com/wz2b/bacnet/npdu"
)

type APDU struct {
	Type          defs.PDUType
	ServiceChoice uint8
	Data          []byte

	Segmented bool  // Only valid for PDU types that support segmentation.
	InvokeID  uint8 // Only valid for PDU types that carry an invoke ID.

	MaxAPDU     uint32 // Only valid for confirmed service request PDUs.
	MaxSegments uint8  // Only valid for confirmed service request PDUs.
}

type APDUContext struct {
	RawPacket []byte

	UDPSender      *net.UDPAddr
	OriginalSender *net.UDPAddr

	BVLCFunction defs.BVLCFunction

	NPDUData []byte
	NPDU     *npdu.NPDU

	APDUData []byte
	APDU     *APDU
}
