package codec

import (
	"encoding/binary"

	"github.com/wz2b/bacnet/defs"
)

/*
Clause 20.2.4: Encoding of an Unsigned Integer Value

	Clause 20.2.1: General Rules for Encoding BACnet Tags
*/
func EncodeContextTaggedUnsigned(apdu []byte, tagNumber defs.ApplicationTagType, value uint32) int {
	valueLength := unsignedEncodedLength(value)

	length := EncodeTag(apdu, tagNumber, true, uint32(valueLength))
	length += EncodeBACnetUnsigned(apdu[length:], value)

	return length
}

func EncodeApplicationTaggedUnsigned(value uint32) []byte {
	length := unsignedEncodedLength(value)
	result := make([]byte, 1+length)

	result[0] =
		(byte(defs.ApplicationTagUnsignedInt) << 4) |
			byte(length)

	switch length {
	case 1:
		result[1] = byte(value)

	case 2:
		binary.BigEndian.PutUint16(
			result[1:3],
			uint16(value),
		)

	case 3:
		result[1] = byte(value >> 16)
		result[2] = byte(value >> 8)
		result[3] = byte(value)

	case 4:
		binary.BigEndian.PutUint32(
			result[1:5],
			value,
		)
	}

	return result
}

// unsignedEncodedLength returns the minimum number of octets required
// to encode value as a BACnet unsigned integer.
func unsignedEncodedLength(value uint32) int {
	switch {
	case value <= 0xFF:
		return 1

	case value <= 0xFFFF:
		return 2

	case value <= 0xFFFFFF:
		return 3

	default:
		return 4
	}
}
