package bacnet

func NewWhoHasAPDU(pdu []byte, high, low int32, isObjectName bool, identifier BACNET_OBJECT_ID, name BACNET_CHARACTER_STRING) *WhoHas {
	if pdu == nil {
		pdu = make([]byte, 50)
	}
	apdu := &WhoHas{
		PDU:          pdu,
		LowLimit:     low,
		HighLimit:    high,
		IsObjectName: isObjectName,
		Identifier:   identifier,
		Name:         name,
	}
	return apdu
}

type WhoHas struct {
	LowLimit     int32
	HighLimit    int32
	IsObjectName bool
	Identifier   BACNET_OBJECT_ID
	Name         BACNET_CHARACTER_STRING
	PDU          []byte
	Length       int
}

func (self *WhoHas) SetLimits(low_limit, high_limit int32) *WhoHas {
	self.LowLimit = low_limit
	self.HighLimit = high_limit
	return self
}

func (self *WhoHas) SetIdentifier(identifier BACNET_OBJECT_ID) *WhoHas {
	self.Identifier = identifier
	return self
}

func (self *WhoHas) SetName(name BACNET_CHARACTER_STRING) *WhoHas {
	self.Name = name
	return self
}

func (self *WhoHas) Encode() Service {
	var encodeIdx int = 2
	self.PDU[0] = PDU_TYPE_UNCONFIRMED_SERVICE_REQUEST
	self.PDU[1] = SERVICE_UNCONFIRMED_WHO_HAS

	if self.LowLimit >= 0 && self.LowLimit <= BACNET_MAX_INSTANCE && self.HighLimit >= 0 && self.HighLimit <= BACNET_MAX_INSTANCE {
		encodeIdx += encode_context_unsigned(self.PDU[encodeIdx:], 0, uint32(self.LowLimit))
		encodeIdx += encode_context_unsigned(self.PDU[encodeIdx:], 1, uint32(self.HighLimit))
	}
	if self.IsObjectName {
		encodeIdx += encode_context_character_string(self.PDU[encodeIdx:], 3, &self.Name)
	} else {
		encodeIdx += encode_context_object_id(self.PDU[encodeIdx:], 2, int(self.Identifier.Type), self.Identifier.Instance)
	}

	self.PDU = self.PDU[:encodeIdx]
	self.Length = encodeIdx
	return self
}

// TODO: WhoHas Service Request Decode
func (self *WhoHas) Decode() Service {
	self.Length = -1

	if len(self.PDU) < 2 {
		return self
	}

	if self.PDU[0] != PDU_TYPE_UNCONFIRMED_SERVICE_REQUEST ||
		self.PDU[1] != SERVICE_UNCONFIRMED_WHO_HAS {
		return self
	}

	decodeIdx := 2

	self.LowLimit = -1
	self.HighLimit = -1

	/*
		Optional device-instance limits.

		If either limit is present, both must be present.
	*/
	if decodeIdx < len(self.PDU) &&
		decode_is_context_tag(self.PDU[decodeIdx:], 0) {

		lenTmp, _, lenValue :=
			decode_tag_number_and_value(self.PDU[decodeIdx:])
		decodeIdx += lenTmp

		if lenTmp <= 0 ||
			lenValue == 0 ||
			lenValue > 4 ||
			int(lenValue) > len(self.PDU)-decodeIdx {
			return self
		}

		lenTmp, decodedValue :=
			decode_unsigned(self.PDU[decodeIdx:], lenValue)
		decodeIdx += lenTmp

		if lenTmp <= 0 ||
			decodedValue > uint32(BACNET_MAX_INSTANCE) {
			return self
		}

		self.LowLimit = int32(decodedValue)

		if decodeIdx >= len(self.PDU) ||
			!decode_is_context_tag(self.PDU[decodeIdx:], 1) {
			return self
		}

		lenTmp, _, lenValue =
			decode_tag_number_and_value(self.PDU[decodeIdx:])
		decodeIdx += lenTmp

		if lenTmp <= 0 ||
			lenValue == 0 ||
			lenValue > 4 ||
			int(lenValue) > len(self.PDU)-decodeIdx {
			return self
		}

		lenTmp, decodedValue =
			decode_unsigned(self.PDU[decodeIdx:], lenValue)
		decodeIdx += lenTmp

		if lenTmp <= 0 ||
			decodedValue > uint32(BACNET_MAX_INSTANCE) {
			return self
		}

		self.HighLimit = int32(decodedValue)

		if self.LowLimit > self.HighLimit {
			return self
		}
	}

	if decodeIdx >= len(self.PDU) {
		return self
	}

	/*
		The requested object must be specified by either:

		  1. Context tag 2: object identifier
		  2. Context tag 3: object name
	*/
	switch {
	case decode_is_context_tag(self.PDU[decodeIdx:], 2):
		self.IsObjectName = false

		lenTmp, _, lenValue :=
			decode_tag_number_and_value(self.PDU[decodeIdx:])
		decodeIdx += lenTmp

		if lenTmp <= 0 ||
			lenValue != 4 ||
			int(lenValue) > len(self.PDU)-decodeIdx {
			return self
		}

		lenTmp,
			self.Identifier.Type,
			self.Identifier.Instance =
			decode_object_id(self.PDU[decodeIdx:])
		decodeIdx += lenTmp

		if lenTmp <= 0 {
			return self
		}

	case decode_is_context_tag(self.PDU[decodeIdx:], 3):
		self.IsObjectName = true

		lenTmp, _, lenValue :=
			decode_tag_number_and_value(self.PDU[decodeIdx:])
		decodeIdx += lenTmp

		// A BACnet character string contains at least the
		// character-set encoding byte.
		if lenTmp <= 0 ||
			lenValue < 1 ||
			int(lenValue) > len(self.PDU)-decodeIdx {
			return self
		}

		lenTmp, self.Name =
			decode_character_string(
				self.PDU[decodeIdx:],
				lenValue,
			)
		decodeIdx += lenTmp

		if lenTmp <= 0 {
			return self
		}

	default:
		return self
	}

	self.Length = decodeIdx
	return self
}
