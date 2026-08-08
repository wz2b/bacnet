package codec

import "github.com/wz2b/bacnet/defs"

type CharacterString struct {
	Encoding byte
	Value    []byte
}

/* from clause 20.2.9 Encoding of a Character String Value */
/* returns the number of apdu bytes consumed */
func EncodeCharacterString(
	apdu []byte,
	value *CharacterString,
) int {
	if value == nil {
		return 0
	}

	length := 1 + len(value.Value)

	if len(apdu) < length {
		return 0
	}

	apdu[0] = value.Encoding
	copy(apdu[1:], value.Value)

	return length
}

/* from clause 20.2.9 Encoding of a Character String Value */
/* returns the number of apdu bytes consumed */
func DecodeCharacterString(
	apdu []byte,
	lenValue uint32,
) (int, CharacterString) {
	var result CharacterString

	// A BACnet character string contains at least the encoding byte.
	if lenValue < 1 {
		return 0, result
	}

	if len(apdu) < int(lenValue) {
		return 0, result
	}

	result.Encoding = apdu[0]

	stringLength := int(lenValue) - 1

	result.Value = append(
		[]byte(nil),
		apdu[1:1+stringLength]...,
	)

	return int(lenValue), result
}

/* from clause 20.2.9 Encoding of a Character String Value */
/* and 20.2.1 General Rules for Encoding BACnet Tags */
/* returns the number of apdu bytes consumed */
func EncodeContextTaggedCharacterString(
	apdu []byte,
	tagNumber byte,
	value *CharacterString,
) int {
	if value == nil {
		return 0
	}

	stringLength := 1 + len(value.Value)

	tagLength := EncodeTag(
		apdu,
		tagNumber,
		true,
		uint32(stringLength),
	)

	valueLength := EncodeCharacterString(
		apdu[tagLength:],
		value,
	)

	if valueLength == 0 {
		return 0
	}

	return tagLength + valueLength
}

/* from clause 20.2.9 Encoding of a Character String Value */
/* and 20.2.1 General Rules for Encoding BACnet Tags */
/* returns the number of apdu bytes consumed */
func EncodeApplicationTaggedCharacterString(
	apdu []byte,
	value *CharacterString,
) int {
	if value == nil {
		return 0
	}

	stringLength := 1 + len(value.Value)

	tagLength := EncodeTag(
		apdu,
		defs.ApplicationTagCharacterString,
		false,
		uint32(stringLength),
	)

	valueLength := EncodeCharacterString(
		apdu[tagLength:],
		value,
	)

	if valueLength == 0 {
		return 0
	}

	return tagLength + valueLength
}
