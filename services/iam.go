package services

import (
	"fmt"

	"github.com/wz2b/bacnet/apdu"
	"github.com/wz2b/bacnet/bactypes"
	"github.com/wz2b/bacnet/codec"
	"github.com/wz2b/bacnet/defs"
)

type IAm struct {
	DeviceID     uint32
	MaxAPDU      uint32
	Segmentation defs.Segmentation
	VendorID     uint16
}

func (i IAm) APDU() apdu.APDU {
	data := make([]byte, 32)
	encodeIdx := 0

	encodeIdx += codec.EncodeApplicationTaggedObjectID(
		data[encodeIdx:],
		bactypes.ObjectID{
			Type:     defs.ObjectDevice,
			Instance: i.DeviceID,
		},
	)

	encoded := codec.EncodeApplicationTaggedUnsigned(i.MaxAPDU)
	copy(data[encodeIdx:], encoded)
	encodeIdx += len(encoded)

	encodeIdx += codec.EncodeApplicationTaggedEnum(
		data[encodeIdx:],
		uint32(i.Segmentation),
	)

	encoded = codec.EncodeApplicationTaggedUnsigned(
		uint32(i.VendorID),
	)
	copy(data[encodeIdx:], encoded)
	encodeIdx += len(encoded)

	return apdu.APDU{
		Type:                     defs.PDUTypeUnconfirmedServiceRequest,
		UnconfirmedServiceChoice: defs.ServiceUnconfirmedIAm,
		Data:                     data[:encodeIdx],
	}
}

func DecodeIAm(a *apdu.APDU) (*IAm, error) {
	if a.Type != defs.PDUTypeUnconfirmedServiceRequest {
		return nil, fmt.Errorf("I-Am requires an unconfirmed service request APDU")
	}

	if a.UnconfirmedServiceChoice != defs.ServiceUnconfirmedIAm {
		return nil, fmt.Errorf("APDU is not an I-Am request")
	}

	result := &IAm{}
	data := a.Data
	decodeIdx := 0

	/*
		Device Object Identifier
	*/
	if decodeIdx >= len(data) {
		return nil, fmt.Errorf("I-Am missing device object identifier")
	}

	tagLength, tagNumber, valueLength :=
		codec.DecodeTagNumberAndValue(data[decodeIdx:])
	if tagLength <= 0 {
		return nil, fmt.Errorf("invalid I-Am object identifier tag")
	}
	decodeIdx += tagLength

	if tagNumber != defs.ApplicationTagObjectID {
		return nil, fmt.Errorf("I-Am expected object identifier")
	}

	if valueLength != 4 || len(data)-decodeIdx < 4 {
		return nil, fmt.Errorf("invalid I-Am object identifier length")
	}

	n, objectType, instance :=
		codec.DecodeObjectID(data[decodeIdx:])
	if n <= 0 {
		return nil, fmt.Errorf("invalid I-Am object identifier")
	}
	decodeIdx += n

	if objectType != defs.ObjectType(defs.ObjectDevice) {
		return nil, fmt.Errorf(
			"I-Am object identifier is not a device object: %d",
			objectType,
		)
	}

	result.DeviceID = instance

	/*
		Max APDU Length Accepted
	*/
	if decodeIdx >= len(data) {
		return nil, fmt.Errorf("I-Am missing max APDU length")
	}

	tagLength, tagNumber, valueLength =
		codec.DecodeTagNumberAndValue(data[decodeIdx:])
	if tagLength <= 0 {
		return nil, fmt.Errorf("invalid I-Am max APDU tag")
	}
	decodeIdx += tagLength

	if tagNumber != defs.ApplicationTagUnsignedInt {
		return nil, fmt.Errorf("I-Am expected unsigned max APDU length")
	}

	if valueLength == 0 ||
		valueLength > 4 ||
		int(valueLength) > len(data)-decodeIdx {
		return nil, fmt.Errorf("invalid I-Am max APDU length")
	}

	n, value := codec.DecodeUnsigned(
		data[decodeIdx:],
		valueLength,
	)
	if n <= 0 {
		return nil, fmt.Errorf("invalid I-Am max APDU value")
	}
	decodeIdx += n

	result.MaxAPDU = value

	/*
		Segmentation Supported
	*/
	if decodeIdx >= len(data) {
		return nil, fmt.Errorf("I-Am missing segmentation value")
	}

	tagLength, tagNumber, valueLength =
		codec.DecodeTagNumberAndValue(data[decodeIdx:])
	if tagLength <= 0 {
		return nil, fmt.Errorf("invalid I-Am segmentation tag")
	}
	decodeIdx += tagLength

	if tagNumber != defs.ApplicationTagEnumerated {
		return nil, fmt.Errorf("I-Am expected enumerated segmentation value")
	}

	if valueLength == 0 ||
		valueLength > 4 ||
		int(valueLength) > len(data)-decodeIdx {
		return nil, fmt.Errorf("invalid I-Am segmentation length")
	}

	n, value = codec.DecodeEnumerated(
		data[decodeIdx:],
		valueLength,
	)
	if n <= 0 {
		return nil, fmt.Errorf("invalid I-Am segmentation value")
	}
	decodeIdx += n

	segmentation := defs.Segmentation(value)

	if !segmentation.Valid() {
		return nil, fmt.Errorf(
			"I-Am segmentation value out of range: %d",
			value,
		)
	}

	result.Segmentation = segmentation
	/*
		Vendor Identifier
	*/
	if decodeIdx >= len(data) {
		return nil, fmt.Errorf("I-Am missing vendor identifier")
	}

	tagLength, tagNumber, valueLength =
		codec.DecodeTagNumberAndValue(data[decodeIdx:])
	if tagLength <= 0 {
		return nil, fmt.Errorf("invalid I-Am vendor identifier tag")
	}
	decodeIdx += tagLength

	if tagNumber != defs.ApplicationTagUnsignedInt {
		return nil, fmt.Errorf("I-Am expected unsigned vendor identifier")
	}

	if valueLength == 0 ||
		valueLength > 4 ||
		int(valueLength) > len(data)-decodeIdx {
		return nil, fmt.Errorf("invalid I-Am vendor identifier length")
	}

	n, value = codec.DecodeUnsigned(
		data[decodeIdx:],
		valueLength,
	)
	if n <= 0 {
		return nil, fmt.Errorf("invalid I-Am vendor identifier")
	}
	decodeIdx += n

	if value > 0xffff {
		return nil, fmt.Errorf("I-Am vendor identifier out of range: %d", value)
	}

	result.VendorID = uint16(value)

	if decodeIdx != len(data) {
		return nil, fmt.Errorf(
			"unexpected trailing data in I-Am: %d bytes",
			len(data)-decodeIdx,
		)
	}

	return result, nil
}
