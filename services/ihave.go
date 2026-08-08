package services

import (
	"fmt"

	"github.com/wz2b/bacnet/apdu"
	"github.com/wz2b/bacnet/bactypes"
	"github.com/wz2b/bacnet/codec"
	"github.com/wz2b/bacnet/defs"
)

type IHave struct {
	// Device containing the matching object.
	DeviceID uint32

	// Object matching the Who-Has request.
	ObjectID bactypes.ObjectID

	// Object name of the matching object.
	ObjectName codec.CharacterString
}

func (i IHave) APDU() apdu.APDU {
	data := make([]byte, 128)
	encodeIdx := 0

	// Device Identifier
	encodeIdx += codec.EncodeApplicationTaggedObjectID(
		data[encodeIdx:],
		bactypes.ObjectID{
			Type:     bactypes.ObjectType(defs.ObjectDevice),
			Instance: i.DeviceID,
		},
	)

	// Object Identifier
	encodeIdx += codec.EncodeApplicationTaggedObjectID(
		data[encodeIdx:],
		i.ObjectID,
	)

	// Object Name
	encodeIdx += codec.EncodeApplicationTaggedCharacterString(
		data[encodeIdx:],
		&i.ObjectName,
	)

	return apdu.APDU{
		Type:          defs.PDUTypeUnconfirmedServiceRequest,
		ServiceChoice: defs.ServiceUnconfirmedIHave,
		Data:          data[:encodeIdx],
	}
}

func DecodeIHave(a *apdu.APDU) (*IHave, error) {
	if a.Type != defs.PDUTypeUnconfirmedServiceRequest {
		return nil, fmt.Errorf("I-Have requires an unconfirmed service request APDU")
	}

	if a.ServiceChoice != defs.ServiceUnconfirmedIHave {
		return nil, fmt.Errorf("APDU is not an I-Have request")
	}

	result := &IHave{}
	data := a.Data
	decodeIdx := 0

	/*
		Device Identifier
	*/
	if decodeIdx >= len(data) {
		return nil, fmt.Errorf("I-Have missing device identifier")
	}

	tagLength, tagNumber, valueLength :=
		codec.DecodeTagNumberAndValue(data[decodeIdx:])
	if tagLength <= 0 {
		return nil, fmt.Errorf("invalid I-Have device identifier tag")
	}
	decodeIdx += tagLength

	if tagNumber != defs.ApplicationTagObjectID {
		return nil, fmt.Errorf("I-Have expected device object identifier")
	}

	if valueLength != 4 || len(data)-decodeIdx < 4 {
		return nil, fmt.Errorf("invalid I-Have device identifier length")
	}

	n, objectType, instance :=
		codec.DecodeObjectID(data[decodeIdx:])
	if n <= 0 {
		return nil, fmt.Errorf("invalid I-Have device identifier")
	}
	decodeIdx += n

	if objectType != bactypes.ObjectType(defs.ObjectDevice) {
		return nil, fmt.Errorf(
			"I-Have device identifier is not a Device object: %d",
			objectType,
		)
	}

	result.DeviceID = instance

	/*
		Object Identifier
	*/
	if decodeIdx >= len(data) {
		return nil, fmt.Errorf("I-Have missing object identifier")
	}

	tagLength, tagNumber, valueLength =
		codec.DecodeTagNumberAndValue(data[decodeIdx:])
	if tagLength <= 0 {
		return nil, fmt.Errorf("invalid I-Have object identifier tag")
	}
	decodeIdx += tagLength

	if tagNumber != defs.ApplicationTagObjectID {
		return nil, fmt.Errorf("I-Have expected object identifier")
	}

	if valueLength != 4 || len(data)-decodeIdx < 4 {
		return nil, fmt.Errorf("invalid I-Have object identifier length")
	}

	n, objectType, instance =
		codec.DecodeObjectID(data[decodeIdx:])
	if n <= 0 {
		return nil, fmt.Errorf("invalid I-Have object identifier")
	}
	decodeIdx += n

	result.ObjectID = bactypes.ObjectID{
		Type:     objectType,
		Instance: instance,
	}

	/*
		Object Name
	*/
	if decodeIdx >= len(data) {
		return nil, fmt.Errorf("I-Have missing object name")
	}

	tagLength, tagNumber, valueLength =
		codec.DecodeTagNumberAndValue(data[decodeIdx:])
	if tagLength <= 0 {
		return nil, fmt.Errorf("invalid I-Have object name tag")
	}
	decodeIdx += tagLength

	if tagNumber != defs.ApplicationTagCharacterString {
		return nil, fmt.Errorf("I-Have expected character string object name")
	}

	if valueLength < 1 ||
		int(valueLength) > len(data)-decodeIdx {
		return nil, fmt.Errorf("invalid I-Have object name length")
	}

	n, objectName := codec.DecodeCharacterString(
		data[decodeIdx:],
		valueLength,
	)
	if n <= 0 {
		return nil, fmt.Errorf("invalid I-Have object name")
	}
	decodeIdx += n

	result.ObjectName = objectName

	if decodeIdx != len(data) {
		return nil, fmt.Errorf(
			"unexpected trailing data in I-Have: %d bytes",
			len(data)-decodeIdx,
		)
	}

	return result, nil
}
