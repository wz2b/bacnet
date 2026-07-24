package bacnet

import "encoding/binary"

func EncodeApplicationReal(value float32) []byte {
	result := make([]byte, 5)
	encodeApplicationReal(result, value)
	return result
}

func EncodeApplicationUnsigned(value uint32) []byte {
	length := unsignedEncodedLength(value)
	result := make([]byte, 1+length)

	result[0] =
		(BACNET_APPLICATION_TAG_UNSIGNED_INT << 4) |
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
