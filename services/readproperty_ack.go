package services

import (
	"fmt"

	"github.com/wz2b/bacnet/apdu"
	"github.com/wz2b/bacnet/bactypes"
	"github.com/wz2b/bacnet/codec"
	"github.com/wz2b/bacnet/defs"
)

type ReadPropertyACK struct {
	Object             bactypes.ObjectID
	PropertyIdentifier uint32
	ArrayIndex         *uint32
	EncodedValue       []byte
}

func DecodeReadPropertyACK(a *apdu.APDU) (*ReadPropertyACK, error) {
	if a.Type != defs.PDUTypeComplexACK {
		return nil, fmt.Errorf(
			"ReadProperty ACK requires a ComplexACK APDU",
		)
	}

	if a.ConfirmedServiceChoice != defs.ServiceConfirmedReadProperty {
		return nil, fmt.Errorf(
			"APDU is not a ReadProperty ACK",
		)
	}

	result := &ReadPropertyACK{}
	data := a.Data
	offset := 0

	/*
		[0] objectIdentifier
	*/
	if offset >= len(data) ||
		!codec.IsContextTag(data[offset:], 0) {
		return nil, fmt.Errorf(
			"ReadProperty ACK missing object identifier",
		)
	}

	tagLength, _, valueLength :=
		codec.DecodeTagNumberAndValue(data[offset:])

	if tagLength <= 0 {
		return nil, fmt.Errorf(
			"invalid ReadProperty ACK object identifier tag",
		)
	}

	offset += tagLength

	if valueLength != 4 ||
		int(valueLength) > len(data)-offset {
		return nil, fmt.Errorf(
			"invalid ReadProperty ACK object identifier length",
		)
	}

	n, objectType, instance :=
		codec.DecodeObjectID(data[offset:])

	if n <= 0 {
		return nil, fmt.Errorf(
			"invalid ReadProperty ACK object identifier",
		)
	}

	offset += n

	result.Object = bactypes.ObjectID{
		Type:     objectType,
		Instance: instance,
	}

	/*
		[1] propertyIdentifier
	*/
	if offset >= len(data) ||
		!codec.IsContextTag(data[offset:], 1) {
		return nil, fmt.Errorf(
			"ReadProperty ACK missing property identifier",
		)
	}

	tagLength, _, valueLength =
		codec.DecodeTagNumberAndValue(data[offset:])

	if tagLength <= 0 {
		return nil, fmt.Errorf(
			"invalid ReadProperty ACK property identifier tag",
		)
	}

	offset += tagLength

	if valueLength == 0 ||
		valueLength > 4 ||
		int(valueLength) > len(data)-offset {
		return nil, fmt.Errorf(
			"invalid ReadProperty ACK property identifier length",
		)
	}

	n, value := codec.DecodeUnsigned(
		data[offset:],
		valueLength,
	)

	if n <= 0 {
		return nil, fmt.Errorf(
			"invalid ReadProperty ACK property identifier",
		)
	}

	offset += n
	result.PropertyIdentifier = value

	/*
		[2] propertyArrayIndex OPTIONAL
	*/
	if offset < len(data) &&
		codec.IsContextTag(data[offset:], 2) {

		tagLength, _, valueLength =
			codec.DecodeTagNumberAndValue(data[offset:])

		if tagLength <= 0 {
			return nil, fmt.Errorf(
				"invalid ReadProperty ACK array index tag",
			)
		}

		offset += tagLength

		if valueLength == 0 ||
			valueLength > 4 ||
			int(valueLength) > len(data)-offset {
			return nil, fmt.Errorf(
				"invalid ReadProperty ACK array index length",
			)
		}

		n, value = codec.DecodeUnsigned(
			data[offset:],
			valueLength,
		)

		if n <= 0 {
			return nil, fmt.Errorf(
				"invalid ReadProperty ACK array index",
			)
		}

		offset += n

		arrayIndex := value
		result.ArrayIndex = &arrayIndex
	}

	/*
		[3] propertyValue
	*/
	if offset >= len(data) ||
		!codec.IsOpeningTagNumber(data[offset:], 3) {
		return nil, fmt.Errorf(
			"ReadProperty ACK missing property value opening tag",
		)
	}

	n, _ = codec.DecodeTagNumber(data[offset:])
	if n <= 0 {
		return nil, fmt.Errorf(
			"invalid ReadProperty ACK property value opening tag",
		)
	}

	offset += n
	valueStart := offset

	/*
		Find the matching context tag 3 closing tag.

		The property value may itself contain constructed values, so
		we track nested opening and closing tags rather than looking
		for a particular byte value.
	*/
	depth := 0

	for offset < len(data) {
		if codec.IsOpeningTag(data[offset:]) {
			depth++

			n, _ = codec.DecodeTagNumber(data[offset:])
			if n <= 0 {
				return nil, fmt.Errorf(
					"invalid opening tag in ReadProperty ACK value",
				)
			}

			offset += n
			continue
		}

		if codec.IsClosingTag(data[offset:]) {
			_, tagNumber := codec.DecodeTagNumber(
				data[offset:],
			)

			if depth == 0 {
				if tagNumber != 3 {
					return nil, fmt.Errorf(
						"unexpected closing tag %d in ReadProperty ACK value",
						tagNumber,
					)
				}

				break
			}

			depth--

			n, _ = codec.DecodeTagNumber(data[offset:])
			if n <= 0 {
				return nil, fmt.Errorf(
					"invalid closing tag in ReadProperty ACK value",
				)
			}

			offset += n
			continue
		}

		tagLength, _, valueLength =
			codec.DecodeTagNumberAndValue(data[offset:])

		if tagLength <= 0 {
			return nil, fmt.Errorf(
				"invalid value in ReadProperty ACK property value",
			)
		}

		consumed := tagLength + int(valueLength)

		if consumed > len(data)-offset {
			return nil, fmt.Errorf(
				"truncated value in ReadProperty ACK property value",
			)
		}

		offset += consumed
	}

	if offset >= len(data) {
		return nil, fmt.Errorf(
			"ReadProperty ACK missing property value closing tag",
		)
	}

	result.EncodedValue = append(
		[]byte(nil),
		data[valueStart:offset]...,
	)

	if !codec.IsClosingTagNumber(data[offset:], 3) {
		return nil, fmt.Errorf(
			"invalid ReadProperty ACK property value closing tag",
		)
	}

	n, _ = codec.DecodeTagNumber(data[offset:])
	if n <= 0 {
		return nil, fmt.Errorf(
			"invalid ReadProperty ACK property value closing tag",
		)
	}

	offset += n

	if offset != len(data) {
		return nil, fmt.Errorf(
			"unexpected trailing data in ReadProperty ACK: %d bytes",
			len(data)-offset,
		)
	}

	return result, nil
}
