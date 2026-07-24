package bacnet

// NewWhoIsRouterToNetwork creates a Who-Is-Router-To-Network
// network-layer message.
//
// networkNumber is optional:
//
//	nil     asks routers to advertise all reachable networks
//	non-nil asks for a router to one specific network
//
// PDU contains only the message-specific data. The NPDU encoder is
// responsible for encoding the network message type.
func NewWhoIsRouterToNetwork(
	pdu []byte,
	networkNumber *uint16,
) *WhoIsRouterToNetwork {
	if pdu == nil {
		pdu = make([]byte, 2)
	}

	message := &WhoIsRouterToNetwork{
		PDU: pdu,
	}

	message.SetOptionalNetworkNumber(networkNumber)

	return message
}

type WhoIsRouterToNetwork struct {
	// NetworkNumber is nil when no particular network was requested.
	NetworkNumber *uint16

	// PDU contains only the message-specific data following the
	// NPDU network message type.
	PDU []byte

	// Length is the number of encoded or decoded message-data octets.
	//
	// A value of -1 indicates malformed message data.
	Length int
}

func (self *WhoIsRouterToNetwork) SetNetworkNumber(
	networkNumber uint16,
) *WhoIsRouterToNetwork {
	value := networkNumber
	self.NetworkNumber = &value

	return self
}

func (self *WhoIsRouterToNetwork) SetOptionalNetworkNumber(
	networkNumber *uint16,
) *WhoIsRouterToNetwork {
	if networkNumber == nil {
		self.NetworkNumber = nil
		return self
	}

	return self.SetNetworkNumber(*networkNumber)
}

func (self *WhoIsRouterToNetwork) ClearNetworkNumber() *WhoIsRouterToNetwork {
	self.NetworkNumber = nil
	return self
}

func (self *WhoIsRouterToNetwork) Encode() *WhoIsRouterToNetwork {
	if self.NetworkNumber == nil {
		self.PDU = self.PDU[:0]
		self.Length = 0
		return self
	}

	self.ensurePDUSize(2)

	encode_unsigned16(
		self.PDU,
		*self.NetworkNumber,
	)

	self.PDU = self.PDU[:2]
	self.Length = 2

	return self
}

func (self *WhoIsRouterToNetwork) Decode() *WhoIsRouterToNetwork {
	self.NetworkNumber = nil

	switch len(self.PDU) {
	case 0:
		// An empty payload asks for all reachable networks.
		self.Length = 0
		return self

	case 2:
		decodedLength, networkNumber := decode_unsigned16(self.PDU)
		if decodedLength != 2 {
			self.Length = -1
			return self
		}

		self.SetNetworkNumber(networkNumber)
		self.Length = decodedLength

		return self

	default:
		// The payload may contain either zero octets or exactly one
		// two-octet network number.
		self.Length = -1
		return self
	}
}

func (self *WhoIsRouterToNetwork) ensurePDUSize(size int) {
	if cap(self.PDU) < size {
		self.PDU = make([]byte, size)
		return
	}

	self.PDU = self.PDU[:size]
}
