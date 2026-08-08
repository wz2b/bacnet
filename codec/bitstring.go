package codec

import (
	"fmt"

	"github.com/wz2b/bacnet/defs"
)

const BIT0 byte = 0x01
const BIT1 byte = 0x02
const BIT2 byte = 0x04
const BIT3 byte = 0x08
const BIT4 byte = 0x10
const BIT5 byte = 0x20
const BIT6 byte = 0x40
const BIT7 byte = 0x80

func EncodeApplicationTaggedBitString(
	bitCount int,
	setBits ...int,
) ([]byte, error) {
	if bitCount <= 0 {
		return nil, fmt.Errorf(
			"BACnet bit string must contain at least one bit",
		)
	}

	byteCount := (bitCount + 7) / 8
	unusedBits := byteCount*8 - bitCount

	bitData := make([]byte, byteCount)

	for _, bit := range setBits {
		if bit < 0 || bit >= bitCount {
			return nil, fmt.Errorf(
				"BACnet bit %d outside bit-string length %d",
				bit,
				bitCount,
			)
		}

		byteIndex := bit / 8
		bitIndex := bit % 8

		// BACnet bit zero is the most-significant bit of the first octet.
		bitData[byteIndex] |= byte(
			1 << (7 - bitIndex),
		)
	}

	/*
		BACnet BIT STRING contents are:

		    number of unused bits
		    packed bit data
	*/
	value := make(
		[]byte,
		1+len(bitData),
	)

	value[0] = byte(unusedBits)
	copy(value[1:], bitData)

	/*
		Application tag 8 is BIT STRING.

		Lengths greater than four use BACnet's extended-length form:
		    0x85
		    one-byte content length
		    content
	*/
	if len(value) <= 4 {
		tag := byte(
			defs.ApplicationTagBitString<<4,
		) | byte(len(value))

		return append([]byte{tag}, value...), nil
	}

	if len(value) <= 253 {
		tag := byte(
			defs.ApplicationTagBitString<<4,
		) | 5

		result := make([]byte, 0, 2+len(value))
		result = append(
			result,
			tag,
			byte(len(value)),
		)
		result = append(result, value...)

		return result, nil
	}

	return nil, fmt.Errorf(
		"BACnet bit string is too large: %d content bytes",
		len(value),
	)
}
