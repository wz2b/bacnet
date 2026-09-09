package services

import (
	"encoding/binary"
	"fmt"
	"math"

	"github.com/wz2b/bacnet/apdu"
	"github.com/wz2b/bacnet/bactypes"
	"github.com/wz2b/bacnet/codec"
	"github.com/wz2b/bacnet/defs"
)

// SubscribeCOV is the service-specific portion of a BACnet
// SubscribeCOV confirmed request.
//
// If both IssueConfirmedNotifications and Lifetime are nil,
// the request cancels the subscription.
type SubscribeCOV struct {
	SubscriberProcessIdentifier uint32
	MonitoredObjectIdentifier   bactypes.ObjectID

	IssueConfirmedNotifications *bool
	Lifetime                    *uint32
}

func (s SubscribeCOV) CancellationRequest() bool {
	return s.IssueConfirmedNotifications == nil &&
		s.Lifetime == nil
}

func (s SubscribeCOV) APDU() apdu.APDU {
	data := make([]byte, 32)
	offset := 0

	// [0] subscriberProcessIdentifier
	offset += codec.EncodeContextTaggedUnsigned(
		data[offset:],
		0,
		s.SubscriberProcessIdentifier,
	)

	// [1] monitoredObjectIdentifier
	offset += codec.EncodeContextTaggedObjectID(
		data[offset:],
		1,
		s.MonitoredObjectIdentifier,
	)

	// [2] issueConfirmedNotifications OPTIONAL
	if s.IssueConfirmedNotifications != nil {
		offset += encodeContextBoolean(
			data[offset:],
			2,
			*s.IssueConfirmedNotifications,
		)
	}

	// [3] lifetime OPTIONAL
	if s.Lifetime != nil {
		offset += codec.EncodeContextTaggedUnsigned(
			data[offset:],
			3,
			*s.Lifetime,
		)
	}

	return apdu.APDU{
		Type:                   defs.PDUTypeConfirmedServiceRequest,
		ConfirmedServiceChoice: defs.ServiceConfirmedSubscribeCOV,
		Data:                   data[:offset],
	}
}

func DecodeSubscribeCOV(a *apdu.APDU) (*SubscribeCOV, error) {
	if a == nil {
		return nil, fmt.Errorf("nil APDU")
	}

	if a.Type != defs.PDUTypeConfirmedServiceRequest {
		return nil, fmt.Errorf(
			"SubscribeCOV requires a confirmed service request APDU",
		)
	}

	if a.ConfirmedServiceChoice != defs.ServiceConfirmedSubscribeCOV {
		return nil, fmt.Errorf(
			"APDU is not a SubscribeCOV request",
		)
	}

	result := &SubscribeCOV{}
	data := a.Data
	offset := 0

	// [0] subscriberProcessIdentifier
	n, value, err := decodeContextUnsigned(data[offset:], 0)
	if err != nil {
		return nil, fmt.Errorf(
			"decode SubscribeCOV subscriber process identifier: %w",
			err,
		)
	}
	offset += n
	result.SubscriberProcessIdentifier = value

	// [1] monitoredObjectIdentifier
	n, objectID, err := decodeContextObjectID(data[offset:], 1)
	if err != nil {
		return nil, fmt.Errorf(
			"decode SubscribeCOV monitored object identifier: %w",
			err,
		)
	}
	offset += n
	result.MonitoredObjectIdentifier = objectID

	// [2] issueConfirmedNotifications OPTIONAL
	if offset < len(data) &&
		codec.IsContextTag(data[offset:], 2) {

		n, value, err := decodeContextBoolean(data[offset:], 2)
		if err != nil {
			return nil, fmt.Errorf(
				"decode SubscribeCOV issue confirmed notifications: %w",
				err,
			)
		}

		offset += n
		result.IssueConfirmedNotifications = new(bool)
		*result.IssueConfirmedNotifications = value
	}

	// [3] lifetime OPTIONAL
	if offset < len(data) &&
		codec.IsContextTag(data[offset:], 3) {

		n, value, err := decodeContextUnsigned(data[offset:], 3)
		if err != nil {
			return nil, fmt.Errorf(
				"decode SubscribeCOV lifetime: %w",
				err,
			)
		}

		offset += n
		result.Lifetime = new(uint32)
		*result.Lifetime = value
	}

	if offset != len(data) {
		return nil, fmt.Errorf(
			"unexpected trailing data in SubscribeCOV: %d bytes",
			len(data)-offset,
		)
	}

	return result, nil
}

// SubscribeCOVProperty is the service-specific portion of a BACnet
// SubscribeCOVProperty confirmed request.
//
// If both IssueConfirmedNotifications and Lifetime are nil,
// the request cancels the subscription.
type SubscribeCOVProperty struct {
	SubscriberProcessIdentifier uint32
	MonitoredObjectIdentifier   bactypes.ObjectID

	IssueConfirmedNotifications *bool
	Lifetime                    *uint32

	MonitoredProperty bactypes.PropertyReference
	COVIncrement      *float32
}

func (s SubscribeCOVProperty) CancellationRequest() bool {
	return s.IssueConfirmedNotifications == nil &&
		s.Lifetime == nil
}

func (s SubscribeCOVProperty) APDU() apdu.APDU {
	data := make([]byte, 64)
	offset := 0

	// [0] subscriberProcessIdentifier
	offset += codec.EncodeContextTaggedUnsigned(
		data[offset:],
		0,
		s.SubscriberProcessIdentifier,
	)

	// [1] monitoredObjectIdentifier
	offset += codec.EncodeContextTaggedObjectID(
		data[offset:],
		1,
		s.MonitoredObjectIdentifier,
	)

	// [2] issueConfirmedNotifications OPTIONAL
	if s.IssueConfirmedNotifications != nil {
		offset += encodeContextBoolean(
			data[offset:],
			2,
			*s.IssueConfirmedNotifications,
		)
	}

	// [3] lifetime OPTIONAL
	if s.Lifetime != nil {
		offset += codec.EncodeContextTaggedUnsigned(
			data[offset:],
			3,
			*s.Lifetime,
		)
	}

	// [4] monitoredPropertyIdentifier
	offset += encodePropertyReference(
		data[offset:],
		4,
		s.MonitoredProperty,
	)

	// [5] covIncrement OPTIONAL
	if s.COVIncrement != nil {
		offset += encodeContextReal(
			data[offset:],
			5,
			*s.COVIncrement,
		)
	}

	return apdu.APDU{
		Type:                   defs.PDUTypeConfirmedServiceRequest,
		ConfirmedServiceChoice: defs.ServiceConfirmedSubscribeCOVProperty,
		Data:                   data[:offset],
	}
}

func DecodeSubscribeCOVProperty(
	a *apdu.APDU,
) (*SubscribeCOVProperty, error) {
	if a == nil {
		return nil, fmt.Errorf("nil APDU")
	}

	if a.Type != defs.PDUTypeConfirmedServiceRequest {
		return nil, fmt.Errorf(
			"SubscribeCOVProperty requires a confirmed service request APDU",
		)
	}

	if a.ConfirmedServiceChoice !=
		defs.ServiceConfirmedSubscribeCOVProperty {

		return nil, fmt.Errorf(
			"APDU is not a SubscribeCOVProperty request",
		)
	}

	result := &SubscribeCOVProperty{}
	data := a.Data
	offset := 0

	// [0] subscriberProcessIdentifier
	n, value, err := decodeContextUnsigned(data[offset:], 0)
	if err != nil {
		return nil, fmt.Errorf(
			"decode SubscribeCOVProperty subscriber process identifier: %w",
			err,
		)
	}
	offset += n
	result.SubscriberProcessIdentifier = value

	// [1] monitoredObjectIdentifier
	n, objectID, err := decodeContextObjectID(data[offset:], 1)
	if err != nil {
		return nil, fmt.Errorf(
			"decode SubscribeCOVProperty monitored object identifier: %w",
			err,
		)
	}
	offset += n
	result.MonitoredObjectIdentifier = objectID

	// [2] issueConfirmedNotifications OPTIONAL
	if offset < len(data) &&
		codec.IsContextTag(data[offset:], 2) {

		n, value, err := decodeContextBoolean(data[offset:], 2)
		if err != nil {
			return nil, fmt.Errorf(
				"decode SubscribeCOVProperty issue confirmed notifications: %w",
				err,
			)
		}

		offset += n
		result.IssueConfirmedNotifications = new(bool)
		*result.IssueConfirmedNotifications = value
	}

	// [3] lifetime OPTIONAL
	if offset < len(data) &&
		codec.IsContextTag(data[offset:], 3) {

		n, value, err := decodeContextUnsigned(data[offset:], 3)
		if err != nil {
			return nil, fmt.Errorf(
				"decode SubscribeCOVProperty lifetime: %w",
				err,
			)
		}

		offset += n
		result.Lifetime = new(uint32)
		*result.Lifetime = value
	}

	// [4] monitoredPropertyIdentifier
	n, property, err := decodePropertyReference(data[offset:], 4)
	if err != nil {
		return nil, fmt.Errorf(
			"decode SubscribeCOVProperty monitored property: %w",
			err,
		)
	}
	offset += n
	result.MonitoredProperty = property

	// [5] covIncrement OPTIONAL
	if offset < len(data) &&
		codec.IsContextTag(data[offset:], 5) {

		n, value, err := decodeContextReal(data[offset:], 5)
		if err != nil {
			return nil, fmt.Errorf(
				"decode SubscribeCOVProperty COV increment: %w",
				err,
			)
		}

		offset += n
		result.COVIncrement = new(float32)
		*result.COVIncrement = value
	}

	if offset != len(data) {
		return nil, fmt.Errorf(
			"unexpected trailing data in SubscribeCOVProperty: %d bytes",
			len(data)-offset,
		)
	}

	return result, nil
}

func encodeContextBoolean(
	data []byte,
	tagNumber defs.ApplicationTagType,
	value bool,
) int {
	length := codec.EncodeTag(
		data,
		tagNumber,
		true,
		1,
	)

	if value {
		data[length] = 1
	} else {
		data[length] = 0
	}

	return length + 1
}

func decodeContextBoolean(
	data []byte,
	tagNumber byte,
) (int, bool, error) {
	if len(data) == 0 ||
		!codec.IsContextTag(data, tagNumber) {

		return 0, false, fmt.Errorf(
			"missing context tag %d",
			tagNumber,
		)
	}

	tagLength, _, valueLength :=
		codec.DecodeTagNumberAndValue(data)

	if tagLength <= 0 ||
		valueLength != 1 ||
		len(data)-tagLength < 1 {

		return 0, false, fmt.Errorf(
			"invalid Boolean context tag %d",
			tagNumber,
		)
	}

	return tagLength + 1, data[tagLength] != 0, nil
}

func decodeContextUnsigned(
	data []byte,
	tagNumber byte,
) (int, uint32, error) {
	if len(data) == 0 ||
		!codec.IsContextTag(data, tagNumber) {

		return 0, 0, fmt.Errorf(
			"missing context tag %d",
			tagNumber,
		)
	}

	tagLength, _, valueLength :=
		codec.DecodeTagNumberAndValue(data)

	if tagLength <= 0 ||
		valueLength == 0 ||
		valueLength > 4 ||
		int(valueLength) > len(data)-tagLength {

		return 0, 0, fmt.Errorf(
			"invalid unsigned context tag %d",
			tagNumber,
		)
	}

	n, value := codec.DecodeUnsigned(
		data[tagLength:],
		valueLength,
	)
	if n <= 0 {
		return 0, 0, fmt.Errorf(
			"invalid unsigned value for context tag %d",
			tagNumber,
		)
	}

	return tagLength + n, value, nil
}

func decodeContextObjectID(
	data []byte,
	tagNumber byte,
) (int, bactypes.ObjectID, error) {
	var objectID bactypes.ObjectID

	if len(data) == 0 ||
		!codec.IsContextTag(data, tagNumber) {

		return 0, objectID, fmt.Errorf(
			"missing context tag %d",
			tagNumber,
		)
	}

	tagLength, _, valueLength :=
		codec.DecodeTagNumberAndValue(data)

	if tagLength <= 0 ||
		valueLength != 4 ||
		len(data)-tagLength < 4 {

		return 0, objectID, fmt.Errorf(
			"invalid object identifier context tag %d",
			tagNumber,
		)
	}

	n, objectType, instance :=
		codec.DecodeObjectID(data[tagLength:])
	if n <= 0 {
		return 0, objectID, fmt.Errorf(
			"invalid object identifier value for context tag %d",
			tagNumber,
		)
	}

	objectID = bactypes.ObjectID{
		Type:     objectType,
		Instance: instance,
	}

	return tagLength + n, objectID, nil
}

func encodePropertyReference(
	data []byte,
	tagNumber byte,
	ref bactypes.PropertyReference,
) int {
	offset := 0

	offset += codec.EncodeOpeningTag(
		data[offset:],
		tagNumber,
	)

	offset += codec.EncodeContextTaggedEnum(
		data[offset:],
		0,
		ref.PropertyIdentifier,
	)

	if ref.ArrayIndex != nil {
		offset += codec.EncodeContextTaggedUnsigned(
			data[offset:],
			1,
			*ref.ArrayIndex,
		)
	}

	offset += codec.EncodeClosingTag(
		data[offset:],
		tagNumber,
	)

	return offset
}

func decodePropertyReference(
	data []byte,
	tagNumber byte,
) (int, bactypes.PropertyReference, error) {
	var ref bactypes.PropertyReference

	if !codec.IsOpeningTagNumber(data, tagNumber) {
		return 0, ref, fmt.Errorf(
			"missing opening tag %d",
			tagNumber,
		)
	}

	tagLength, _ := codec.DecodeTagNumber(data)
	if tagLength <= 0 {
		return 0, ref, fmt.Errorf(
			"invalid opening tag %d",
			tagNumber,
		)
	}

	offset := tagLength

	n, value, err := decodeContextUnsigned(data[offset:], 0)
	if err != nil {
		return 0, ref, fmt.Errorf(
			"decode property identifier: %w",
			err,
		)
	}
	offset += n
	ref.PropertyIdentifier = value

	if offset < len(data) &&
		codec.IsContextTag(data[offset:], 1) {

		n, value, err := decodeContextUnsigned(data[offset:], 1)
		if err != nil {
			return 0, ref, fmt.Errorf(
				"decode property array index: %w",
				err,
			)
		}

		offset += n
		ref.ArrayIndex = new(uint32)
		*ref.ArrayIndex = value
	}

	if offset >= len(data) ||
		!codec.IsClosingTagNumber(data[offset:], tagNumber) {

		return 0, ref, fmt.Errorf(
			"missing closing tag %d",
			tagNumber,
		)
	}

	tagLength, _ = codec.DecodeTagNumber(data[offset:])
	if tagLength <= 0 {
		return 0, ref, fmt.Errorf(
			"invalid closing tag %d",
			tagNumber,
		)
	}

	offset += tagLength

	return offset, ref, nil
}

func encodeContextReal(
	data []byte,
	tagNumber defs.ApplicationTagType,
	value float32,
) int {
	offset := codec.EncodeTag(
		data,
		tagNumber,
		true,
		4,
	)

	binary.BigEndian.PutUint32(
		data[offset:offset+4],
		math.Float32bits(value),
	)

	return offset + 4
}

func decodeContextReal(
	data []byte,
	tagNumber byte,
) (int, float32, error) {
	if len(data) == 0 ||
		!codec.IsContextTag(data, tagNumber) {

		return 0, 0, fmt.Errorf(
			"missing context tag %d",
			tagNumber,
		)
	}

	tagLength, _, valueLength :=
		codec.DecodeTagNumberAndValue(data)

	if tagLength <= 0 ||
		valueLength != 4 ||
		len(data)-tagLength < 4 {

		return 0, 0, fmt.Errorf(
			"invalid Real context tag %d",
			tagNumber,
		)
	}

	value := math.Float32frombits(
		binary.BigEndian.Uint32(
			data[tagLength : tagLength+4],
		),
	)

	return tagLength + 4, value, nil
}
