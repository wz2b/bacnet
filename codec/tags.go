package codec

/* from clause 20.2.1 General Rules for Encoding BACnet Tags */
/* returns the number of apdu bytes consumed */
func EncodeTag(
	apdu []byte,
	tagNumber byte,
	contextSpecific bool,
	lenValueType uint32,
) int {
	length := 1

	apdu[0] = 0

	if contextSpecific {
		apdu[0] = BIT3
	}

	/*
	 * Tag number.
	 */
	if tagNumber <= 14 {
		apdu[0] |= tagNumber << 4
	} else {
		apdu[0] |= 0xF0
		apdu[length] = tagNumber
		length++
	}

	/*
	 * Length/value/type field.
	 */
	if lenValueType <= 4 {
		apdu[0] |= byte(lenValueType)
		return length
	}

	apdu[0] |= 5

	/*
	 * At this point, length is the index of the next unused byte.
	 *
	 * For an ordinary tag, this is byte 1.
	 * For an extended-number tag, this is byte 2.
	 */
	switch {
	case lenValueType <= 253:
		apdu[length] = byte(lenValueType)
		length++

	case lenValueType <= 65535:
		apdu[length] = 254
		length++

		length += EncodeUnsigned16(
			apdu[length:],
			uint16(lenValueType),
		)

	default:
		apdu[length] = 255
		length++

		length += EncodeUnsigned32(
			apdu[length:],
			lenValueType,
		)
	}

	return length
}

/* from clause 20.2.1.3.2 Constructed Data */
/* returns the number of apdu bytes consumed */
func EncodeOpeningTag(apdu []byte, tagNumber byte) int {
	length := 1

	/* set class field to context specific */
	apdu[0] = BIT3

	/* additional tag byte after this byte for extended tag byte */
	if tagNumber <= 14 {
		apdu[0] |= tagNumber << 4
	} else {
		apdu[0] |= 0xF0
		apdu[1] = tagNumber
		length++
	}

	/* set type field to opening tag */
	apdu[0] |= 6

	return length
}

/* from clause 20.2.1.3.2 Constructed Data */
/* returns the number of apdu bytes consumed */
func EncodeClosingTag(apdu []byte, tagNumber byte) int {
	length := 1

	/* set class field to context specific */
	apdu[0] = BIT3

	/* additional tag byte after this byte for extended tag byte */
	if tagNumber <= 14 {
		apdu[0] |= tagNumber << 4
	} else {
		apdu[0] |= 0xF0
		apdu[1] = tagNumber
		length++
	}

	/* set type field to closing tag */
	apdu[0] |= 7

	return length
}

func DecodeTagNumber(apdu []byte) (int, byte) {
	if len(apdu) == 0 {
		return 0, 0
	}

	if IsExtendedTagNumber(apdu[0]) {
		if len(apdu) < 2 {
			return 0, 0
		}

		return 2, apdu[1]
	}

	return 1, apdu[0] >> 4
}

func DecodeTagNumberAndValue(apdu []byte) (int, byte, uint32) {
	if len(apdu) == 0 {
		return 0, 0, 0
	}

	length, tagNumber := DecodeTagNumber(apdu)
	if length <= 0 {
		return 0, 0, 0
	}

	var value uint32

	if IsExtendedValue(apdu[0]) {
		if length >= len(apdu) {
			return 0, 0, 0
		}

		switch apdu[length] {
		case 255:
			length++

			if len(apdu)-length < 4 {
				return 0, 0, 0
			}

			n, value32 := DecodeUnsigned32(apdu[length:])
			length += n
			value = value32

		case 254:
			length++

			if len(apdu)-length < 2 {
				return 0, 0, 0
			}

			n, value16 := DecodeUnsigned16(apdu[length:])
			length += n
			value = uint32(value16)

		default:
			value = uint32(apdu[length])
			length++
		}

	} else if IsOpeningTagByte(apdu[0]) {
		value = 0

	} else if IsClosingTagByte(apdu[0]) {
		value = 0

	} else {
		value = uint32(apdu[0] & 0x07)
	}

	return length, tagNumber, value
}

/* from clause 20.2.1.3.2 Constructed Data */
/* returns true if the tag is context specific and matches */
func IsContextTag(apdu []byte, tagNumber byte) bool {
	if len(apdu) == 0 {
		return false
	}

	_, actualTagNumber := DecodeTagNumber(apdu)

	return IsContextSpecific(apdu[0]) &&
		tagNumber == actualTagNumber
}

/*
 * IsOpeningTagByte and IsClosingTagByte inspect a single encoded tag octet.
 *
 * The low three bits contain the length/value/type field.
 * Values 6 and 7 identify opening and closing tags respectively.
 */
func IsOpeningTagByte(value byte) bool {
	return value&0x07 == 6
}

func IsClosingTagByte(value byte) bool {
	return value&0x07 == 7
}

/*
 * IsOpeningTag and IsClosingTag inspect the tag at the beginning
 * of an encoded byte slice.
 */
func IsOpeningTag(apdu []byte) bool {
	if len(apdu) == 0 {
		return false
	}

	return IsContextSpecific(apdu[0]) &&
		IsOpeningTagByte(apdu[0])
}

func IsClosingTag(apdu []byte) bool {
	if len(apdu) == 0 {
		return false
	}

	return IsContextSpecific(apdu[0]) &&
		IsClosingTagByte(apdu[0])
}

func IsOpeningTagNumber(apdu []byte, tagNumber byte) bool {
	if !IsOpeningTag(apdu) {
		return false
	}

	_, actualTagNumber := DecodeTagNumber(apdu)

	return actualTagNumber == tagNumber
}

func IsClosingTagNumber(apdu []byte, tagNumber byte) bool {
	if !IsClosingTag(apdu) {
		return false
	}

	_, actualTagNumber := DecodeTagNumber(apdu)

	return actualTagNumber == tagNumber
}

/* from clause 20.2.1.2 Tag Number */
/* true if extended tag numbering is used */
func IsExtendedTagNumber(x byte) bool {
	return (x & 0xF0) == 0xF0
}

/* from clause 20.2.1.3.1 Primitive Data */
/* true if the extended value is used */
func IsExtendedValue(x byte) bool {
	return (x & 0x07) == 5
}

/* from clause 20.2.1.1 Class */
/* true if the tag is context specific */
func IsContextSpecific(x byte) bool {
	//fmt.Printf("%d %d %d   is context specific: %t\n", x, x&BIT3, BIT3, (x&BIT3) == BIT3)
	return (x & BIT3) == BIT3
}
