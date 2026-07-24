package bacnet

func NewIHaveAPDU(
	pdu []byte,
	deviceId uint32,
	objectId BACNET_OBJECT_ID,
	objectName string,
) *IHave {
	if pdu == nil {
		pdu = make([]byte, MAX_APDU)
	}

	name := BACNET_CHARACTER_STRING{
		Value: make([]byte, MAX_CHARACTER_STRING_BYTES),
	}

	if !characterstring_init_ansi(
		&name,
		[]byte(objectName),
	) {
		panic("I-Have object name exceeds BACnet character-string capacity")
	}

	return &IHave{
		PDU:        pdu,
		DeviceId:   deviceId,
		ObjectId:   objectId,
		ObjectName: name,
	}
}

type IHave struct {
	PDU []byte

	// Device containing the matching object.
	DeviceId uint32

	// Object matching the Who-Has request.
	ObjectId BACNET_OBJECT_ID

	// Object_Name of the matching object.
	ObjectName BACNET_CHARACTER_STRING

	Length int
}

func (self *IHave) Encode() *IHave {
	encodeIdx := 2

	self.PDU[0] = PDU_TYPE_UNCONFIRMED_SERVICE_REQUEST
	self.PDU[1] = SERVICE_UNCONFIRMED_I_HAVE

	// Device Identifier
	encodeIdx += encode_application_object_id(
		self.PDU[encodeIdx:],
		int(OBJECT_DEVICE),
		self.DeviceId,
	)

	// Object Identifier
	encodeIdx += encode_application_object_id(
		self.PDU[encodeIdx:],
		int(self.ObjectId.Type),
		self.ObjectId.Instance,
	)

	// Object Name
	encodeIdx += encode_application_character_string(
		self.PDU[encodeIdx:],
		&self.ObjectName,
	)

	self.PDU = self.PDU[:encodeIdx]
	self.Length = encodeIdx

	return self
}

func (self *IHave) Decode() *IHave {
	var (
		decodeIdx      int
		offset         int
		lenTmp         int
		tagNumber      byte
		lenValue       uint32
		objectType     uint16
		objectInstance uint32
		slice          []byte
	)

	if len(self.PDU) >= 2 &&
		self.PDU[0] == PDU_TYPE_UNCONFIRMED_SERVICE_REQUEST &&
		self.PDU[1] == SERVICE_UNCONFIRMED_I_HAVE {

		slice = self.PDU[2:]
		offset = 2
	} else {
		slice = self.PDU
	}

	/*
	 * Device Identifier
	 */

	lenTmp, tagNumber, lenValue =
		decode_tag_number_and_value(slice[decodeIdx:])
	decodeIdx += lenTmp

	if tagNumber != BACNET_APPLICATION_TAG_OBJECT_ID {
		panic("invalid I-Have Device Identifier tag")
	}

	lenTmp, objectType, objectInstance =
		decode_object_id(slice[decodeIdx:])
	decodeIdx += lenTmp

	if objectType != uint16(OBJECT_DEVICE) {
		panic("I-Have Device Identifier is not a Device object")
	}

	self.DeviceId = objectInstance

	/*
	 * Object Identifier
	 */

	lenTmp, tagNumber, lenValue =
		decode_tag_number_and_value(slice[decodeIdx:])
	decodeIdx += lenTmp

	if tagNumber != BACNET_APPLICATION_TAG_OBJECT_ID {
		panic("invalid I-Have Object Identifier tag")
	}

	lenTmp, objectType, objectInstance =
		decode_object_id(slice[decodeIdx:])
	decodeIdx += lenTmp

	self.ObjectId = BACNET_OBJECT_ID{
		Type:     objectType,
		Instance: objectInstance,
	}

	/*
	 * Object Name
	 */

	lenTmp, tagNumber, lenValue =
		decode_tag_number_and_value(slice[decodeIdx:])
	decodeIdx += lenTmp

	if tagNumber != BACNET_APPLICATION_TAG_CHARACTER_STRING {
		panic("invalid I-Have Object Name tag")
	}

	lenTmp, self.ObjectName =
		decode_character_string(slice[decodeIdx:], lenValue)

	if lenTmp <= 0 {
		panic("invalid I-Have Object Name value")
	}

	decodeIdx += lenTmp

	self.Length = decodeIdx + offset

	return self
}
