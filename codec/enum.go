package codec

import "github.com/wz2b/bacnet/defs"

/* from clause 20.2.11 Encoding of an Enumerated Value */
/* and 20.2.1 General Rules for Encoding BACnet Tags */
/* returns the number of apdu bytes consumed */
func EncodeBacnetEnum(apdu []byte, value uint32) int {
	return EncodeBACnetUnsigned(apdu, value)
}

/* from clause 20.2.11 Encoding of an Enumerated Value */
/* and 20.2.1 General Rules for Encoding BACnet Tags */
/* returns the number of apdu bytes consumed */
func DecodeEnumerated(apdu []byte, len_value uint32) (int, uint32) {
	var unsigned_value uint32 = 0
	var length int

	length, unsigned_value = DecodeUnsigned(apdu, len_value)

	return length, unsigned_value
}

/* from clause 20.2.11 Encoding of an Enumerated Value */
/* and 20.2.1 General Rules for Encoding BACnet Tags */
/* returns the number of apdu bytes consumed */
func EncodeContextTaggedEnum(apdu []byte, tag_number byte, value uint32) int {
	var length int = 0

	if value < 0x100 {
		length = 1
	} else if value < 0x10000 {
		length = 2
	} else if value < 0x1000000 {
		length = 3
	} else {
		length = 4
	}

	length = EncodeTag(apdu, tag_number, true, uint32(length))
	length += EncodeBacnetEnum(apdu[length:], value)

	return length
}

/* from clause 20.2.11 Encoding of an Enumerated Value */
/* and 20.2.1 General Rules for Encoding BACnet Tags */
/* returns the number of apdu bytes consumed */
func EncodeApplicationTaggedEnum(apdu []byte, value uint32) int {
	var length int = 0 /* return value */

	/* assumes that the tag only consumes 1 octet */
	length = EncodeBacnetEnum(apdu[1:], value)
	length += EncodeTag(apdu, defs.ApplicationTagEnumerated, false, uint32(length))

	return length
}
