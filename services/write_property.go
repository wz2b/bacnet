package services

import (
	"fmt"

	"github.com/wz2b/bacnet/apdu"
	"github.com/wz2b/bacnet/bactypes"
	"github.com/wz2b/bacnet/codec"
	"github.com/wz2b/bacnet/defs"
)

// WritePropertyRequest is the service-specific portion of a BACnet
// WriteProperty confirmed request.
//
// EncodedValue contains the BACnet application-encoded property value,
// without the surrounding context tag [3] opening/closing tags.
type WritePropertyRequest struct {
	Object             bactypes.ObjectID
	PropertyIdentifier uint32
	ArrayIndex         *uint32
	EncodedValue       []byte
	Priority           *uint32
}

func (r WritePropertyRequest) APDU() apdu.APDU {
	data := make([]byte, 64+len(r.EncodedValue))
	offset := 0

	// [0] objectIdentifier
	offset += codec.EncodeContextTaggedObjectID(
		data[offset:],
		0,
		r.Object,
	)

	// [1] propertyIdentifier
	offset += codec.EncodeContextTaggedEnum(
		data[offset:],
		1,
		r.PropertyIdentifier,
	)

	// [2] propertyArrayIndex OPTIONAL
	if r.ArrayIndex != nil {
		offset += codec.EncodeContextTaggedUnsigned(
			data[offset:],
			2,
			*r.ArrayIndex,
		)
	}

	// [3] propertyValue
	offset += codec.EncodeOpeningTag(
		data[offset:],
		3,
	)

	copy(data[offset:], r.EncodedValue)
	offset += len(r.EncodedValue)

	offset += codec.EncodeClosingTag(
		data[offset:],
		3,
	)

	// [4] priority OPTIONAL
	if r.Priority != nil {
		offset += codec.EncodeContextTaggedUnsigned(
			data[offset:],
			4,
			*r.Priority,
		)
	}

	return apdu.APDU{
		Type:                   defs.PDUTypeConfirmedServiceRequest,
		ConfirmedServiceChoice: defs.ServiceConfirmedWriteProperty,
		Data:                   data[:offset],
	}
}

func DecodeWritePropertyRequest(a *apdu.APDU) (*WritePropertyRequest, error) {
	if a.Type != defs.PDUTypeConfirmedServiceRequest {
		return nil, fmt.Errorf(
			"WriteProperty requires a confirmed service request APDU",
		)
	}

	if a.ConfirmedServiceChoice != defs.ServiceConfirmedWriteProperty {
		return nil, fmt.Errorf(
			"APDU is not a WriteProperty request",
		)
	}

	result := &WritePropertyRequest{}
	data := a.Data
	offset := 0

	/*
		[0] objectIdentifier
	*/
	if offset >= len(data) || !codec.IsContextTag(data[offset:], 0) {
		return nil, fmt.Errorf(
			"WriteProperty missing object identifier",
		)
	}

	tagLength, _, valueLength :=
		codec.DecodeTagNumberAndValue(data[offset:])
	offset += tagLength

	if tagLength <= 0 ||
		valueLength != 4 ||
		int(valueLength) > len(data)-offset {
		return nil, fmt.Errorf(
			"invalid WriteProperty object identifier",
		)
	}

	n, objectType, instance :=
		codec.DecodeObjectID(data[offset:])
	if n <= 0 {
		return nil, fmt.Errorf(
			"invalid WriteProperty object identifier value",
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
	if offset >= len(data) || !codec.IsContextTag(data[offset:], 1) {
		return nil, fmt.Errorf(
			"WriteProperty missing property identifier",
		)
	}

	tagLength, _, valueLength =
		codec.DecodeTagNumberAndValue(data[offset:])
	offset += tagLength

	if tagLength <= 0 ||
		valueLength == 0 ||
		valueLength > 4 ||
		int(valueLength) > len(data)-offset {
		return nil, fmt.Errorf(
			"invalid WriteProperty property identifier",
		)
	}

	n, value := codec.DecodeUnsigned(
		data[offset:],
		valueLength,
	)
	if n <= 0 {
		return nil, fmt.Errorf(
			"invalid WriteProperty property identifier value",
		)
	}
	offset += n

	result.PropertyIdentifier = value

	/*
		[2] propertyArrayIndex OPTIONAL
	*/
	if offset < len(data) && codec.IsContextTag(data[offset:], 2) {
		tagLength, _, valueLength =
			codec.DecodeTagNumberAndValue(data[offset:])
		offset += tagLength

		if tagLength <= 0 ||
			valueLength == 0 ||
			valueLength > 4 ||
			int(valueLength) > len(data)-offset {
			return nil, fmt.Errorf(
				"invalid WriteProperty array index",
			)
		}

		n, value = codec.DecodeUnsigned(
			data[offset:],
			valueLength,
		)
		if n <= 0 {
			return nil, fmt.Errorf(
				"invalid WriteProperty array index value",
			)
		}
		offset += n

		result.ArrayIndex = new(uint32)
		*result.ArrayIndex = value
	}

	/*
		[3] propertyValue

		Opening tag 3 = 0x3e
		Closing tag 3 = 0x3f

		The contents are retained in their original BACnet encoding.
	*/
	if offset >= len(data) || data[offset] != 0x3e {
		return nil, fmt.Errorf(
			"WriteProperty missing property value opening tag",
		)
	}
	offset++

	valueStart := offset

	/*
		For the moment, property values we're interested in are simple
		application-tagged values such as REAL. Find the closing [3].

		If/when we support constructed property values, this should use
		the same constructed-tag walker used by ReadPropertyACK.
	*/
	for offset < len(data) && data[offset] != 0x3f {
		tagLength, _, valueLength =
			codec.DecodeTagNumberAndValue(data[offset:])

		if tagLength <= 0 {
			return nil, fmt.Errorf(
				"invalid WriteProperty encoded value",
			)
		}

		offset += tagLength

		if int(valueLength) > len(data)-offset {
			return nil, fmt.Errorf(
				"truncated WriteProperty encoded value",
			)
		}

		offset += int(valueLength)
	}

	if offset >= len(data) || data[offset] != 0x3f {
		return nil, fmt.Errorf(
			"WriteProperty missing property value closing tag",
		)
	}

	result.EncodedValue = append(
		[]byte(nil),
		data[valueStart:offset]...,
	)

	offset++ // closing [3]

	/*
		[4] priority OPTIONAL
	*/
	if offset < len(data) {
		if !codec.IsContextTag(data[offset:], 4) {
			return nil, fmt.Errorf(
				"unexpected WriteProperty data at offset %d",
				offset,
			)
		}

		tagLength, _, valueLength =
			codec.DecodeTagNumberAndValue(data[offset:])
		offset += tagLength

		if tagLength <= 0 ||
			valueLength == 0 ||
			valueLength > 4 ||
			int(valueLength) > len(data)-offset {
			return nil, fmt.Errorf(
				"invalid WriteProperty priority",
			)
		}

		n, value = codec.DecodeUnsigned(
			data[offset:],
			valueLength,
		)
		if n <= 0 {
			return nil, fmt.Errorf(
				"invalid WriteProperty priority value",
			)
		}
		offset += n

		if value < 1 || value > 16 {
			return nil, fmt.Errorf(
				"WriteProperty priority must be 1..16, got %d",
				value,
			)
		}

		result.Priority = new(uint32)
		*result.Priority = value
	}

	if offset != len(data) {
		return nil, fmt.Errorf(
			"unexpected trailing data in WriteProperty: %d bytes",
			len(data)-offset,
		)
	}

	return result, nil
}
