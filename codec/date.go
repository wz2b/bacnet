package codec

import (
	"github.com/wz2b/bacnet/bactypes"
	"github.com/wz2b/bacnet/defs"
)

/* from clause 20.2.12 Encoding of a Date Value */
/* and 20.2.1 General Rules for Encoding BACnet Tags */
/* returns the number of apdu bytes consumed */
func EncodeDate(apdu []byte, bdate *bactypes.Date) int {
	/* allow 2 digit years */
	if bdate.Year >= 1900 {
		apdu[0] = byte(bdate.Year - 1900)
	} else if bdate.Year < 0x100 {
		apdu[0] = byte(bdate.Year)
	} else {
		/*
		 ** Don't try and guess what the user meant here. Just fail
		 */
		return -1
	}

	apdu[1] = bdate.Month
	apdu[2] = bdate.Day
	apdu[3] = bdate.Weekday

	return 4
}

func DecodeDate(apdu []byte) (int, bactypes.Date) {
	var result bactypes.Date

	if len(apdu) < 4 {
		return 0, result
	}

	result.Year = uint16(apdu[0]) + 1900
	result.Month = apdu[1]
	result.Day = apdu[2]
	result.Weekday = apdu[3]

	return 4, result
}

func EncodeContextTaggedDate(apdu []byte, tag_number byte, bdate *bactypes.Date) int {
	var len int = 0 /* return value */

	/* length of date is 4 octets, as per 20.2.12 */
	len = EncodeTag(apdu, tag_number, true, 4)
	len += EncodeDate(apdu[len:], bdate)
	return len
}

/* from clause 20.2.13 Encoding of a Time Value */
/* and 20.2.1 General Rules for Encoding BACnet Tags */
/* returns the number of apdu bytes consumed */
func EncodeApplicationTaggedDate(
	apdu []byte,
	value *bactypes.Date,
) int {
	if value == nil {
		return 0
	}

	length := EncodeTag(
		apdu,
		defs.ApplicationTagDate,
		false,
		4,
	)

	length += EncodeDate(
		apdu[length:],
		value,
	)

	return length
}

func DecodeApplicationTaggedDate(
	apdu []byte,
) (int, bactypes.Date) {
	var result bactypes.Date

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

	if tagNumber != defs.ApplicationTagDate {
		return 0, result
	}

	if valueLength != 4 {
		return 0, result
	}

	if len(apdu)-tagLength < 4 {
		return 0, result
	}

	n, value := DecodeDate(apdu[tagLength:])
	if n != 4 {
		return 0, result
	}

	return tagLength + n, value
}
