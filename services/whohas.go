package services

import (
	"fmt"

	"github.com/wz2b/bacnet/apdu"
	"github.com/wz2b/bacnet/bactypes"
	"github.com/wz2b/bacnet/codec"
	"github.com/wz2b/bacnet/defs"
)

type WhoHas struct {
	LowLimit  *uint32
	HighLimit *uint32

	IsObjectName bool
	Identifier   bactypes.ObjectID
	Name         codec.CharacterString
}

func (w WhoHas) APDU() apdu.APDU {
	data := make([]byte, 64)
	encodeIdx := 0

	if w.LowLimit != nil && w.HighLimit != nil {
		encodeIdx += codec.EncodeContextTaggedUnsigned(
			data[encodeIdx:],
			0,
			*w.LowLimit,
		)

		encodeIdx += codec.EncodeContextTaggedUnsigned(
			data[encodeIdx:],
			1,
			*w.HighLimit,
		)
	}

	if w.IsObjectName {
		encodeIdx += codec.EncodeContextTaggedCharacterString(
			data[encodeIdx:],
			3,
			&w.Name,
		)
	} else {
		encodeIdx += codec.EncodeContextTaggedObjectID(
			data[encodeIdx:],
			2,
			w.Identifier,
		)
	}

	return apdu.APDU{
		Type:                     defs.PDUTypeUnconfirmedServiceRequest,
		UnconfirmedServiceChoice: defs.ServiceUnconfirmedWhoHas,
		Data:                     data[:encodeIdx],
	}
}

func DecodeWhoHas(a *apdu.APDU) (*WhoHas, error) {
	if a.Type != defs.PDUTypeUnconfirmedServiceRequest {
		return nil, fmt.Errorf(
			"Who-Has requires an unconfirmed service request APDU",
		)
	}

	if a.UnconfirmedServiceChoice != defs.ServiceUnconfirmedWhoHas {
		return nil, fmt.Errorf(
			"APDU is not a Who-Has request",
		)
	}

	result := &WhoHas{}

	data := a.Data
	decodeIdx := 0

	/*
		Optional [0] lowLimit and [1] highLimit.
		If one is present, both must be present.
	*/
	if decodeIdx < len(data) &&
		codec.IsContextTag(data[decodeIdx:], 0) {

		tagLength, _, valueLength :=
			codec.DecodeTagNumberAndValue(data[decodeIdx:])
		if tagLength <= 0 {
			return nil, fmt.Errorf(
				"invalid Who-Has low limit tag",
			)
		}

		decodeIdx += tagLength

		if valueLength == 0 ||
			valueLength > 4 ||
			int(valueLength) > len(data)-decodeIdx {
			return nil, fmt.Errorf(
				"invalid Who-Has low limit length",
			)
		}

		n, value := codec.DecodeUnsigned(
			data[decodeIdx:],
			valueLength,
		)
		if n <= 0 {
			return nil, fmt.Errorf(
				"invalid Who-Has low limit",
			)
		}

		decodeIdx += n
		lowLimit := uint32(value)
		result.LowLimit = &lowLimit

		if decodeIdx >= len(data) ||
			!codec.IsContextTag(data[decodeIdx:], 1) {
			return nil, fmt.Errorf(
				"Who-Has low limit present without high limit",
			)
		}

		tagLength, _, valueLength =
			codec.DecodeTagNumberAndValue(data[decodeIdx:])
		if tagLength <= 0 {
			return nil, fmt.Errorf(
				"invalid Who-Has high limit tag",
			)
		}

		decodeIdx += tagLength

		if valueLength == 0 ||
			valueLength > 4 ||
			int(valueLength) > len(data)-decodeIdx {
			return nil, fmt.Errorf(
				"invalid Who-Has high limit length",
			)
		}

		n, value = codec.DecodeUnsigned(
			data[decodeIdx:],
			valueLength,
		)
		if n <= 0 {
			return nil, fmt.Errorf(
				"invalid Who-Has high limit",
			)
		}

		decodeIdx += n

		highLimit := value
		result.HighLimit = &highLimit
	}

	if decodeIdx >= len(data) {
		return nil, fmt.Errorf(
			"Who-Has missing object identifier or object name",
		)
	}

	/*
		Either [2] objectIdentifier or [3] objectName.
	*/
	switch {
	case codec.IsContextTag(data[decodeIdx:], 2):
		tagLength, _, valueLength :=
			codec.DecodeTagNumberAndValue(data[decodeIdx:])
		if tagLength <= 0 {
			return nil, fmt.Errorf(
				"invalid Who-Has object identifier tag",
			)
		}

		decodeIdx += tagLength

		if valueLength != 4 ||
			int(valueLength) > len(data)-decodeIdx {
			return nil, fmt.Errorf(
				"invalid Who-Has object identifier length",
			)
		}

		n, objectType, instance :=
			codec.DecodeObjectID(data[decodeIdx:])
		if n <= 0 {
			return nil, fmt.Errorf(
				"invalid Who-Has object identifier",
			)
		}

		decodeIdx += n

		result.Identifier = bactypes.ObjectID{
			Type:     objectType,
			Instance: instance,
		}

		result.IsObjectName = false

	case codec.IsContextTag(data[decodeIdx:], 3):
		tagLength, _, valueLength :=
			codec.DecodeTagNumberAndValue(data[decodeIdx:])
		if tagLength <= 0 {
			return nil, fmt.Errorf(
				"invalid Who-Has object name tag",
			)
		}

		decodeIdx += tagLength

		if valueLength < 1 ||
			int(valueLength) > len(data)-decodeIdx {
			return nil, fmt.Errorf(
				"invalid Who-Has object name length",
			)
		}

		n, name := codec.DecodeCharacterString(
			data[decodeIdx:],
			valueLength,
		)
		if n <= 0 {
			return nil, fmt.Errorf(
				"invalid Who-Has object name",
			)
		}

		decodeIdx += n

		result.Name = name
		result.IsObjectName = true

	default:
		return nil, fmt.Errorf(
			"Who-Has expected context tag 2 or 3",
		)
	}

	if decodeIdx != len(data) {
		return nil, fmt.Errorf(
			"unexpected trailing data in Who-Has: %d bytes",
			len(data)-decodeIdx,
		)
	}

	return result, nil
}
