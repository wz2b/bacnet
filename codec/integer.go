package codec

func EncodeUnsigned16(apdu []byte, value uint16) int {
	apdu[0] = byte((value & 0xff00) >> 8)
	apdu[1] = byte(value & 0x00ff)
	return 2
}

func DecodeUnsigned16(apdu []byte) (int, uint16) {
	var value uint16 = 0
	value = (uint16(apdu[0]) << 8) & 0xff00
	value |= uint16(apdu[1]) & 0x00ff
	return 2, value
}

func EncodeUnsigned24(apdu []byte, value uint32) int {
	apdu[0] = byte((value & 0xff0000) >> 16)
	apdu[1] = byte((value & 0x00ff00) >> 8)
	apdu[2] = byte(value & 0x0000ff)
	return 3
}

func DecodeUnsigned24(apdu []byte) (int, uint32) {
	var value uint32 = 0
	value = (uint32(apdu[0]) << 16) & 0x00ff0000
	value |= (uint32(apdu[1]) << 8) & 0x0000ff00
	value |= uint32(apdu[2]) & 0x000000ff
	return 3, value
}

func EncodeUnsigned32(apdu []byte, value uint32) int {
	apdu[0] = byte((value & 0xff000000) >> 24)
	apdu[1] = byte((value & 0x00ff0000) >> 16)
	apdu[2] = byte((value & 0x0000ff00) >> 8)
	apdu[3] = byte(value & 0x000000ff)
	return 4
}

func DecodeUnsigned32(apdu []byte) (int, uint32) {
	var value uint32 = 0
	value = (uint32(apdu[0]) << 24) & 0xff000000
	value |= (uint32(apdu[1]) << 16) & 0x00ff0000
	value |= (uint32(apdu[2]) << 8) & 0x0000ff00
	value |= uint32(apdu[3]) & 0x000000ff
	return 4, value
}

func EncodeSigned8(apdu []byte, value int8) int {
	apdu[0] = byte(value)
	return 1
}

// TODO: I had to change value to int64 to compile
func DecodeSigned8(apdu []byte, value *int64) int {
	if *value != 0 {
		if apdu[0]&0x80 != 0 {
			*value = 0xFFFFFF00
		} else {
			*value = 0
		}
		*value |= int64(apdu[0]) & 0x000000ff
	}
	return 1
}

func EncodeSigned16(apdu []byte, value int32) int {
	apdu[0] = byte((value & 0xff00) >> 8)
	apdu[1] = byte(value & 0x00ff)
	return 2
}

// TODO: I had to change value to int64 to compile
func DecodeSigned16(apdu []byte, value *int64) int {
	if *value != 0 {
		if apdu[0]&0x80 != 0 {
			*value = 0xFFFF0000
		} else {
			*value = 0
		}
		*value |= (int64(apdu[0]) << 8) & 0x0000ff00
		*value |= int64(apdu[1]) & 0x000000ff
	}
	return 2
}

func EncodeSigned24(apdu []byte, value int32) int {
	apdu[0] = byte((value & 0xff0000) >> 16)
	apdu[1] = byte((value & 0x00ff00) >> 8)
	apdu[2] = byte(value & 0x0000ff)
	return 3
}

// TODO: I had to change value to int64 to compile
func DecodeSigned24(apdu []byte, value *int64) int {
	if *value != 0 {
		if apdu[0]&0x80 != 0 {
			*value = 0xFF000000
		} else {
			*value = 0
		}
		*value |= (int64(apdu[0]) << 16) & 0x00ff0000
		*value |= (int64(apdu[1]) << 8) & 0x0000ff00
		*value |= int64(apdu[2]) & 0x000000ff
	}
	return 3
}

func EncodeSigned32(apdu []byte, value int64) int {
	apdu[0] = byte((value & 0xff000000) >> 24)
	apdu[1] = byte((value & 0x00ff0000) >> 16)
	apdu[2] = byte((value & 0x0000ff00) >> 8)
	apdu[3] = byte(value & 0x000000ff)
	return 4
}

// TODO: I had to change value to int64 to compile
func DecodeSigned32(apdu []byte, value *int64) int {
	if *value != 0 {
		*value = (int64(apdu[0]) << 24) & 0xff000000
		*value |= (int64(apdu[1]) << 16) & 0x00ff0000
		*value |= (int64(apdu[2]) << 8) & 0x0000ff00
		*value |= int64(apdu[3]) & 0x000000ff
	}
	return 4
}

/* from clause 20.2.4 Encoding of an Unsigned Integer Value */
/* and 20.2.1 General Rules for Encoding BACnet Tags */
/* returns the number of apdu bytes consumed */
func DecodeUnsigned(apdu []byte, len_value uint32) (int, uint32) {
	var unsigned16_value uint16 = 0
	var value uint32 = 0
	switch len_value {
	case 1:
		value = uint32(apdu[0])
	case 2:
		_, unsigned16_value = DecodeUnsigned16(apdu)
		value = uint32(unsigned16_value)
	case 3:
		_, value = DecodeUnsigned24(apdu)
	case 4:
		_, value = DecodeUnsigned32(apdu)
	default:
		value = 0
	}
	return int(len_value), value
}

/* from clause 20.2.4 Encoding of an Unsigned Integer Value */
/* and 20.2.1 General Rules for Encoding BACnet Tags */
/* returns the number of apdu bytes consumed */
func EncodeBACnetUnsigned(apdu []byte, value uint32) int {
	var len int = 0 /* return value */
	if value < 0x100 {
		apdu[0] = byte(value)
		len = 1
	} else if value < 0x10000 {
		len = EncodeUnsigned16(apdu, uint16(value))
	} else if value < 0x1000000 {
		len = EncodeUnsigned24(apdu, value)
	} else {
		len = EncodeUnsigned32(apdu, value)
	}
	return len
}
