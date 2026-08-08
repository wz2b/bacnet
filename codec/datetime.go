package codec

import (
	"github.com/wz2b/bacnet/bactypes"
)

func EncodeDatetime(apdu []byte, value *bactypes.DateTime) int {
	var len int = 0
	var apdu_len int = 0

	if value != nil {
		len = EncodeApplicationTaggedDate(apdu, &value.Date)
		apdu_len += len
		len = EncodeApplicationTaggedTime(apdu[apdu_len:], &value.Time)
		apdu_len += len
	}
	return apdu_len
}

func DecodeDatetime(
	apdu []byte,
) (int, bactypes.DateTime) {
	var result bactypes.DateTime
	decodeIdx := 0

	n, date := DecodeApplicationTaggedDate(
		apdu[decodeIdx:],
	)
	if n <= 0 {
		return 0, result
	}
	decodeIdx += n

	n, time := DecodeApplicationTaggedTime(
		apdu[decodeIdx:],
	)
	if n <= 0 {
		return 0, result
	}
	decodeIdx += n

	result.Date = date
	result.Time = time

	return decodeIdx, result
}

func DecodeContextTaggedDatetime(
	apdu []byte,
	tagNumber byte,
) (int, bactypes.DateTime) {
	var result bactypes.DateTime
	decodeIdx := 0

	if len(apdu) == 0 ||
		!IsOpeningTagNumber(apdu, tagNumber) {
		return 0, result
	}

	n, _ := DecodeTagNumber(apdu)
	if n <= 0 {
		return 0, result
	}
	decodeIdx += n

	n, value := DecodeDatetime(apdu[decodeIdx:])
	if n <= 0 {
		return 0, result
	}
	decodeIdx += n

	if decodeIdx >= len(apdu) ||
		!IsClosingTagNumber(apdu[decodeIdx:], tagNumber) {
		return 0, result
	}

	n, _ = DecodeTagNumber(apdu[decodeIdx:])
	if n <= 0 {
		return 0, result
	}
	decodeIdx += n

	return decodeIdx, value
}

func EncodeContextTaggedDatetime(apdu []byte, tag_number byte, value *bactypes.DateTime) int {
	var len int = 0
	var apdu_len int = 0

	if value != nil {
		len = EncodeOpeningTag(apdu[apdu_len:], tag_number)
		apdu_len += len
		len = EncodeDatetime(apdu[apdu_len:], value)
		apdu_len += len
		len = EncodeClosingTag(apdu[apdu_len:], tag_number)
		apdu_len += len
	}
	return apdu_len
}
