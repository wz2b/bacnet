package codec

import (
	"github.com/wz2b/bacnet/bactypes"
	"github.com/wz2b/bacnet/defs"
)

/* from clause 20.2.13 Encoding of a Time Value */
/* and 20.2.1 General Rules for Encoding BACnet Tags */
/* returns the number of apdu bytes consumed */
func EncodeTime(apdu []byte, btime *bactypes.Time) int {
	apdu[0] = btime.Hour
	apdu[1] = btime.Minute
	apdu[2] = btime.Second
	apdu[3] = btime.Hundredth
	return 4
}

func DecodeTime(apdu []byte) (int, bactypes.Time) {
	var result bactypes.Time

	if len(apdu) < 4 {
		return 0, result
	}

	result.Hour = apdu[0]
	result.Minute = apdu[1]
	result.Second = apdu[2]
	result.Hundredth = apdu[3]

	return 4, result
}

func EncodeContextTaggedTime(apdu []byte, tag_number defs.ApplicationTagType, btime *bactypes.Time) int {
	var length int = 0 /* return value */

	/* length of time is 4 octets, as per 20.2.13 */
	length = EncodeTag(apdu, tag_number, true, 4)
	length += EncodeTime(apdu[length:], btime)

	return length
}

func DecodeContextTaggedTime(
	apdu []byte,
	tagNumber byte,
) (int, bactypes.Time) {
	var result bactypes.Time

	if len(apdu) == 0 || !IsContextTag(apdu, tagNumber) {
		return 0, result
	}

	tagLength, _, valueLength := DecodeTagNumberAndValue(apdu)
	if tagLength <= 0 {
		return 0, result
	}

	if valueLength != 4 {
		return 0, result
	}

	if len(apdu)-tagLength < 4 {
		return 0, result
	}

	n, value := DecodeTime(apdu[tagLength:])
	if n != 4 {
		return 0, result
	}

	return tagLength + n, value
}

func EncodeApplicationTaggedTime(
	apdu []byte,
	btime *bactypes.Time,
) int {
	if btime == nil {
		return 0
	}

	length := EncodeTag(
		apdu,
		defs.ApplicationTagTime,
		false,
		4,
	)

	length += EncodeTime(
		apdu[length:],
		btime,
	)

	return length
}

func DecodeApplicationTaggedTime(
	apdu []byte,
) (int, bactypes.Time) {
	var result bactypes.Time

	if len(apdu) == 0 {
		return 0, result
	}

	tagLength, tagNumber, valueLength :=
		DecodeTagNumberAndValue(apdu)

	if tagLength <= 0 {
		return 0, result
	}

	if IsContextSpecific(apdu[0]) {
		return 0, result
	}

	if tagNumber != defs.ApplicationTagTime {
		return 0, result
	}

	if valueLength != 4 {
		return 0, result
	}

	if len(apdu)-tagLength < 4 {
		return 0, result
	}

	n, value := DecodeTime(apdu[tagLength:])
	if n != 4 {
		return 0, result
	}

	return tagLength + n, value
}
