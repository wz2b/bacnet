package services

import (
	"fmt"

	"github.com/wz2b/bacnet/apdu"
	"github.com/wz2b/bacnet/bactypes"
	"github.com/wz2b/bacnet/codec"
	"github.com/wz2b/bacnet/defs"
)

type ACKAlarm struct {
	AckProcessIdentifier  uint32
	EventObjectIdentifier bactypes.ObjectID
	EventStateAcked       byte
	EventTimeStamp        bactypes.Timestamp
	AckSource             codec.CharacterString
	AckTimeStamp          bactypes.Timestamp
}

func (a ACKAlarm) APDU() apdu.APDU {
	data := make([]byte, 128)
	encodeIdx := 0

	encodeIdx += codec.EncodeContextTaggedUnsigned(
		data[encodeIdx:],
		0,
		a.AckProcessIdentifier,
	)

	encodeIdx += codec.EncodeContextTaggedObjectID(
		data[encodeIdx:],
		1,
		a.EventObjectIdentifier,
	)

	encodeIdx += codec.EncodeContextTaggedEnum(
		data[encodeIdx:],
		2,
		uint32(a.EventStateAcked),
	)

	encodeIdx += codec.EncodeContextTaggedTimestamp(
		data[encodeIdx:],
		3,
		&a.EventTimeStamp,
	)

	encodeIdx += codec.EncodeContextTaggedCharacterString(
		data[encodeIdx:],
		4,
		&a.AckSource,
	)

	encodeIdx += codec.EncodeContextTaggedTimestamp(
		data[encodeIdx:],
		5,
		&a.AckTimeStamp,
	)

	return apdu.APDU{
		Type:          defs.PDUTypeConfirmedServiceRequest,
		ServiceChoice: defs.ServiceConfirmedAcknowledgeAlarm,
		Data:          data[:encodeIdx],
	}
}

func DecodeACKAlarm(a *apdu.APDU) (*ACKAlarm, error) {
	if a.Type != defs.PDUTypeConfirmedServiceRequest {
		return nil, fmt.Errorf("AcknowledgeAlarm requires a confirmed service request APDU")
	}

	if a.ServiceChoice != defs.ServiceConfirmedAcknowledgeAlarm {
		return nil, fmt.Errorf("APDU is not an AcknowledgeAlarm request")
	}

	result := &ACKAlarm{}
	data := a.Data
	decodeIdx := 0

	/*
		[0] acknowledgingProcessIdentifier
	*/
	if decodeIdx >= len(data) ||
		!codec.IsContextTag(data[decodeIdx:], 0) {
		return nil, fmt.Errorf("AcknowledgeAlarm missing acknowledging process identifier")
	}

	n, _, length := codec.DecodeTagNumberAndValue(data[decodeIdx:])
	if n <= 0 {
		return nil, fmt.Errorf("invalid AcknowledgeAlarm process identifier tag")
	}
	decodeIdx += n

	if length == 0 ||
		length > 4 ||
		int(length) > len(data)-decodeIdx {
		return nil, fmt.Errorf("invalid AcknowledgeAlarm process identifier length")
	}

	n, value := codec.DecodeUnsigned(data[decodeIdx:], length)
	if n <= 0 {
		return nil, fmt.Errorf("invalid AcknowledgeAlarm process identifier")
	}
	decodeIdx += n

	result.AckProcessIdentifier = value

	/*
		[1] eventObjectIdentifier
	*/
	if decodeIdx >= len(data) ||
		!codec.IsContextTag(data[decodeIdx:], 1) {
		return nil, fmt.Errorf("AcknowledgeAlarm missing event object identifier")
	}

	n, _, length = codec.DecodeTagNumberAndValue(data[decodeIdx:])
	if n <= 0 {
		return nil, fmt.Errorf("invalid AcknowledgeAlarm object identifier tag")
	}
	decodeIdx += n

	if length != 4 || int(length) > len(data)-decodeIdx {
		return nil, fmt.Errorf("invalid AcknowledgeAlarm object identifier length")
	}

	n, objectType, instance := codec.DecodeObjectID(data[decodeIdx:])
	if n <= 0 {
		return nil, fmt.Errorf("invalid AcknowledgeAlarm object identifier")
	}
	decodeIdx += n

	result.EventObjectIdentifier = bactypes.ObjectID{
		Type:     objectType,
		Instance: instance,
	}

	/*
		[2] eventStateAcknowledged
	*/
	if decodeIdx >= len(data) ||
		!codec.IsContextTag(data[decodeIdx:], 2) {
		return nil, fmt.Errorf("AcknowledgeAlarm missing event state")
	}

	n, _, length = codec.DecodeTagNumberAndValue(data[decodeIdx:])
	if n <= 0 {
		return nil, fmt.Errorf("invalid AcknowledgeAlarm event state tag")
	}
	decodeIdx += n

	if length == 0 ||
		length > 4 ||
		int(length) > len(data)-decodeIdx {
		return nil, fmt.Errorf("invalid AcknowledgeAlarm event state length")
	}

	n, value = codec.DecodeUnsigned(data[decodeIdx:], length)
	if n <= 0 {
		return nil, fmt.Errorf("invalid AcknowledgeAlarm event state")
	}
	decodeIdx += n

	result.EventStateAcked = byte(value)

	/*
		[3] eventTimeStamp
	*/
	if decodeIdx >= len(data) ||
		!codec.IsContextTag(data[decodeIdx:], 3) {
		return nil, fmt.Errorf("AcknowledgeAlarm missing event timestamp")
	}

	n, timestamp := codec.DecodeContextTaggedTimestamp(
		data[decodeIdx:],
		3,
	)
	if n <= 0 {
		return nil, fmt.Errorf("invalid AcknowledgeAlarm event timestamp")
	}
	decodeIdx += n

	result.EventTimeStamp = timestamp

	/*
		[4] acknowledgementSource
	*/
	if decodeIdx >= len(data) ||
		!codec.IsContextTag(data[decodeIdx:], 4) {
		return nil, fmt.Errorf("AcknowledgeAlarm missing acknowledgement source")
	}

	n, _, length = codec.DecodeTagNumberAndValue(data[decodeIdx:])
	if n <= 0 {
		return nil, fmt.Errorf("invalid AcknowledgeAlarm acknowledgement source tag")
	}
	decodeIdx += n

	if length < 1 || int(length) > len(data)-decodeIdx {
		return nil, fmt.Errorf("invalid AcknowledgeAlarm acknowledgement source length")
	}

	n, source := codec.DecodeCharacterString(
		data[decodeIdx:],
		length,
	)
	if n <= 0 {
		return nil, fmt.Errorf("invalid AcknowledgeAlarm acknowledgement source")
	}
	decodeIdx += n

	result.AckSource = source

	/*
		[5] acknowledgementTimeStamp
	*/
	if decodeIdx >= len(data) ||
		!codec.IsContextTag(data[decodeIdx:], 5) {
		return nil, fmt.Errorf("AcknowledgeAlarm missing acknowledgement timestamp")
	}

	n, timestamp = codec.DecodeContextTaggedTimestamp(
		data[decodeIdx:],
		5,
	)
	if n <= 0 {
		return nil, fmt.Errorf("invalid AcknowledgeAlarm acknowledgement timestamp")
	}
	decodeIdx += n

	result.AckTimeStamp = timestamp

	if decodeIdx != len(data) {
		return nil, fmt.Errorf(
			"unexpected trailing data in AcknowledgeAlarm: %d bytes",
			len(data)-decodeIdx,
		)
	}

	return result, nil
}
