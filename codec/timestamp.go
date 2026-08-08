package codec

import (
	"github.com/wz2b/bacnet/bactypes"
)

func EncodeContextTaggedTimestamp(apdu []byte, tag_number byte, value *bactypes.Timestamp) int {
	var len int = 0
	var apdu_len int = 0

	if value != nil {
		len = EncodeOpeningTag(apdu[apdu_len:], tag_number)
		apdu_len += len
		len = EncodeTimestamp(apdu[apdu_len:], value)
		apdu_len += len
		len = EncodeClosingTag(apdu[apdu_len:], tag_number)
		apdu_len += len
	}
	return apdu_len
}

func EncodeTimestamp(apdu []byte, value *bactypes.Timestamp) int {
	var length int = 0

	if value != nil {
		if value.Tag == bactypes.TimestampTime {
			length = EncodeContextTaggedTime(apdu, 0, &value.Time)
		} else if value.Tag == bactypes.TimestampSequence {
			length = EncodeContextTaggedUnsigned(apdu, 1, uint32(value.SequenceNum))
		} else if value.Tag == bactypes.TimestampDateTime {
			length = EncodeContextTaggedDatetime(apdu, 2, &value.DateTime)
		}
	}

	return length
}

func DecodeTimestamp(apdu []byte) (int, bactypes.Timestamp) {
	var timestamp bactypes.Timestamp

	if len(apdu) == 0 {
		return 0, timestamp
	}

	switch {
	case IsContextTag(apdu, 0):
		length, value := DecodeContextTaggedTime(apdu, 0)
		if length <= 0 {
			return 0, timestamp
		}

		timestamp.Tag = bactypes.TimestampTime
		timestamp.Time = value

		return length, timestamp

	case IsContextTag(apdu, 1):
		tagLength, _, valueLength := DecodeTagNumberAndValue(apdu)
		if tagLength <= 0 {
			return 0, timestamp
		}

		if valueLength == 0 ||
			valueLength > 4 ||
			int(valueLength) > len(apdu)-tagLength {
			return 0, timestamp
		}

		valueLengthActual, value := DecodeUnsigned(
			apdu[tagLength:],
			valueLength,
		)
		if valueLengthActual <= 0 {
			return 0, timestamp
		}

		timestamp.Tag = bactypes.TimestampSequence
		timestamp.SequenceNum = uint16(value)

		return tagLength + valueLengthActual, timestamp

	case IsContextTag(apdu, 2):
		length, value := DecodeContextTaggedDatetime(apdu, 2)
		if length <= 0 {
			return 0, timestamp
		}

		timestamp.Tag = bactypes.TimestampDateTime
		timestamp.DateTime = value

		return length, timestamp
	}

	return 0, timestamp
}

func DecodeContextTaggedTimestamp(
	apdu []byte,
	tagNumber byte,
) (int, bactypes.Timestamp) {
	var result bactypes.Timestamp
	decodeIdx := 0

	if !IsOpeningTagNumber(apdu, tagNumber) {
		return 0, result
	}

	n, _ := DecodeTagNumber(apdu)
	if n <= 0 {
		return 0, result
	}
	decodeIdx += n

	n, timestamp := DecodeTimestamp(apdu[decodeIdx:])
	if n <= 0 {
		return 0, result
	}
	decodeIdx += n

	if !IsClosingTagNumber(apdu[decodeIdx:], tagNumber) {
		return 0, result
	}

	n, _ = DecodeTagNumber(apdu[decodeIdx:])
	if n <= 0 {
		return 0, result
	}
	decodeIdx += n

	return decodeIdx, timestamp
}
