package services

import (
	"fmt"

	"github.com/wz2b/bacnet/apdu"
	"github.com/wz2b/bacnet/bactypes"
	"github.com/wz2b/bacnet/codec"
	"github.com/wz2b/bacnet/defs"
)

// ReadPropertyRequest is the service-specific portion of a BACnet
// ReadProperty confirmed request.
//
// PDU contains only the service data, not the confirmed-request APDU header.
type ReadPropertyRequest struct {
	Object             bactypes.ObjectID
	PropertyIdentifier uint32
	ArrayIndex         *uint32
}

func (r ReadPropertyRequest) APDU() apdu.APDU {
	data := make([]byte, 32)
	offset := 0

	offset += codec.EncodeContextTaggedObjectID(
		data[offset:],
		0,
		r.Object,
	)

	offset += codec.EncodeContextTaggedEnum(
		data[offset:],
		1,
		r.PropertyIdentifier,
	)

	if r.ArrayIndex != nil {
		offset += codec.EncodeContextTaggedUnsigned(
			data[offset:],
			2,
			*r.ArrayIndex,
		)
	}

	return apdu.APDU{
		Type:                   defs.PDUTypeConfirmedServiceRequest,
		ConfirmedServiceChoice: defs.ServiceConfirmedReadProperty,
		Data:                   data[:offset],
	}
}

func DecodeReadPropertyRequest(a *apdu.APDU) (*ReadPropertyRequest, error) {
	if a.Type != defs.PDUTypeConfirmedServiceRequest {
		return nil, fmt.Errorf(
			"ReadProperty requires a confirmed service request APDU",
		)
	}

	if a.ConfirmedServiceChoice != defs.ServiceConfirmedReadProperty {
		return nil, fmt.Errorf(
			"APDU is not a ReadProperty request",
		)
	}

	result := &ReadPropertyRequest{}
	data := a.Data
	offset := 0

	/*
		[0] objectIdentifier
	*/
	if offset >= len(data) || !codec.IsContextTag(data[offset:], 0) {
		return nil, fmt.Errorf(
			"ReadProperty missing object identifier",
		)
	}

	tagLength, _, valueLength :=
		codec.DecodeTagNumberAndValue(data[offset:])
	offset += tagLength

	if tagLength <= 0 ||
		valueLength != 4 ||
		int(valueLength) > len(data)-offset {
		return nil, fmt.Errorf(
			"invalid ReadProperty object identifier",
		)
	}

	n, objectType, instance :=
		codec.DecodeObjectID(data[offset:])
	if n <= 0 {
		return nil, fmt.Errorf(
			"invalid ReadProperty object identifier value",
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
			"ReadProperty missing property identifier",
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
			"invalid ReadProperty property identifier",
		)
	}

	n, value := codec.DecodeUnsigned(
		data[offset:],
		valueLength,
	)
	if n <= 0 {
		return nil, fmt.Errorf(
			"invalid ReadProperty property identifier value",
		)
	}
	offset += n

	result.PropertyIdentifier = value

	/*
		[2] propertyArrayIndex OPTIONAL
	*/
	if offset < len(data) {
		if !codec.IsContextTag(data[offset:], 2) {
			return nil, fmt.Errorf(
				"unexpected ReadProperty data at offset %d",
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
				"invalid ReadProperty array index",
			)
		}

		n, value = codec.DecodeUnsigned(
			data[offset:],
			valueLength,
		)
		if n <= 0 {
			return nil, fmt.Errorf(
				"invalid ReadProperty array index value",
			)
		}
		offset += n

		result.ArrayIndex = new(uint32)
		*result.ArrayIndex = value
	}

	if offset != len(data) {
		return nil, fmt.Errorf(
			"unexpected trailing data in ReadProperty: %d bytes",
			len(data)-offset,
		)
	}

	return result, nil
}

func (r ReadPropertyACK) APDU() apdu.APDU {
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

	return apdu.APDU{
		Type:                   defs.PDUTypeComplexACK,
		ConfirmedServiceChoice: defs.ServiceConfirmedReadProperty,
		Data:                   data[:offset],
	}
}
