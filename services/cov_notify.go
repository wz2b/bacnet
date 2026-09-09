package services

import (
	"fmt"

	"github.com/wz2b/bacnet/apdu"
	"github.com/wz2b/bacnet/bactypes"
	"github.com/wz2b/bacnet/codec"
	"github.com/wz2b/bacnet/defs"
)

// COVNotification contains the service data shared by
// ConfirmedCOVNotification and UnconfirmedCOVNotification.
type COVNotification struct {
	SubscriberProcessIdentifier uint32

	// InitiatingDeviceID is the instance number of the Device object
	// that generated the notification.
	InitiatingDeviceID uint32

	MonitoredObjectIdentifier bactypes.ObjectID
	TimeRemaining             uint32
	Values                    []bactypes.PropertyValue
}

type ConfirmedCOVNotification struct {
	COVNotification
}

type UnconfirmedCOVNotification struct {
	COVNotification
}

func (n ConfirmedCOVNotification) APDU() apdu.APDU {
	return apdu.APDU{
		Type:                   defs.PDUTypeConfirmedServiceRequest,
		ConfirmedServiceChoice: defs.ServiceConfirmedCOVNotification,
		Data:                   encodeCOVNotificationData(n.COVNotification),
	}
}

func (n UnconfirmedCOVNotification) APDU() apdu.APDU {
	return apdu.APDU{
		Type:                     defs.PDUTypeUnconfirmedServiceRequest,
		UnconfirmedServiceChoice: defs.ServiceUnconfirmedCOVNotification,
		Data:                     encodeCOVNotificationData(n.COVNotification),
	}
}

func DecodeConfirmedCOVNotification(
	a *apdu.APDU,
) (*ConfirmedCOVNotification, error) {
	if a == nil {
		return nil, fmt.Errorf("nil APDU")
	}

	if a.Type != defs.PDUTypeConfirmedServiceRequest {
		return nil, fmt.Errorf(
			"ConfirmedCOVNotification requires a confirmed service request APDU",
		)
	}

	if a.ConfirmedServiceChoice !=
		defs.ServiceConfirmedCOVNotification {

		return nil, fmt.Errorf(
			"APDU is not a ConfirmedCOVNotification request",
		)
	}

	notification, err := decodeCOVNotificationData(a.Data)
	if err != nil {
		return nil, err
	}

	return &ConfirmedCOVNotification{
		COVNotification: notification,
	}, nil
}

func DecodeUnconfirmedCOVNotification(
	a *apdu.APDU,
) (*UnconfirmedCOVNotification, error) {
	if a == nil {
		return nil, fmt.Errorf("nil APDU")
	}

	if a.Type != defs.PDUTypeUnconfirmedServiceRequest {
		return nil, fmt.Errorf(
			"UnconfirmedCOVNotification requires an unconfirmed service request APDU",
		)
	}

	if a.UnconfirmedServiceChoice !=
		defs.ServiceUnconfirmedCOVNotification {

		return nil, fmt.Errorf(
			"APDU is not an UnconfirmedCOVNotification request",
		)
	}

	notification, err := decodeCOVNotificationData(a.Data)
	if err != nil {
		return nil, err
	}

	return &UnconfirmedCOVNotification{
		COVNotification: notification,
	}, nil
}

func encodeCOVNotificationData(
	notification COVNotification,
) []byte {
	size := 64

	for _, value := range notification.Values {
		size += 32 + len(value.EncodedValue)
	}

	data := make([]byte, size)
	offset := 0

	// [0] subscriberProcessIdentifier
	offset += codec.EncodeContextTaggedUnsigned(
		data[offset:],
		0,
		notification.SubscriberProcessIdentifier,
	)

	// [1] initiatingDeviceIdentifier
	offset += codec.EncodeContextTaggedObjectID(
		data[offset:],
		1,
		bactypes.ObjectID{
			Type:     defs.ObjectDevice,
			Instance: notification.InitiatingDeviceID,
		},
	)

	// [2] monitoredObjectIdentifier
	offset += codec.EncodeContextTaggedObjectID(
		data[offset:],
		2,
		notification.MonitoredObjectIdentifier,
	)

	// [3] timeRemaining
	offset += codec.EncodeContextTaggedUnsigned(
		data[offset:],
		3,
		notification.TimeRemaining,
	)

	// [4] listOfValues
	offset += codec.EncodeOpeningTag(
		data[offset:],
		4,
	)

	for _, value := range notification.Values {
		offset += encodeCOVPropertyValue(
			data[offset:],
			value,
		)
	}

	offset += codec.EncodeClosingTag(
		data[offset:],
		4,
	)

	return data[:offset]
}

func decodeCOVNotificationData(
	data []byte,
) (COVNotification, error) {
	var result COVNotification
	offset := 0

	// [0] subscriberProcessIdentifier
	n, value, err := decodeContextUnsigned(data[offset:], 0)
	if err != nil {
		return result, fmt.Errorf(
			"decode COV subscriber process identifier: %w",
			err,
		)
	}
	offset += n
	result.SubscriberProcessIdentifier = value

	// [1] initiatingDeviceIdentifier
	n, objectID, err := decodeContextObjectID(data[offset:], 1)
	if err != nil {
		return result, fmt.Errorf(
			"decode COV initiating device identifier: %w",
			err,
		)
	}
	offset += n

	if objectID.Type != defs.ObjectDevice {
		return result, fmt.Errorf(
			"COV initiating device identifier is not a Device object: %d",
			objectID.Type,
		)
	}

	result.InitiatingDeviceID = objectID.Instance

	// [2] monitoredObjectIdentifier
	n, objectID, err = decodeContextObjectID(data[offset:], 2)
	if err != nil {
		return result, fmt.Errorf(
			"decode COV monitored object identifier: %w",
			err,
		)
	}
	offset += n
	result.MonitoredObjectIdentifier = objectID

	// [3] timeRemaining
	n, value, err = decodeContextUnsigned(data[offset:], 3)
	if err != nil {
		return result, fmt.Errorf(
			"decode COV time remaining: %w",
			err,
		)
	}
	offset += n
	result.TimeRemaining = value

	// [4] listOfValues
	if offset >= len(data) ||
		!codec.IsOpeningTagNumber(data[offset:], 4) {

		return result, fmt.Errorf(
			"COV notification missing list-of-values opening tag",
		)
	}

	tagLength, _ := codec.DecodeTagNumber(data[offset:])
	if tagLength <= 0 {
		return result, fmt.Errorf(
			"invalid COV list-of-values opening tag",
		)
	}
	offset += tagLength

	for {
		if offset >= len(data) {
			return result, fmt.Errorf(
				"COV notification missing list-of-values closing tag",
			)
		}

		if codec.IsClosingTagNumber(data[offset:], 4) {
			tagLength, _ := codec.DecodeTagNumber(data[offset:])
			if tagLength <= 0 {
				return result, fmt.Errorf(
					"invalid COV list-of-values closing tag",
				)
			}

			offset += tagLength
			break
		}

		n, propertyValue, err :=
			decodeCOVPropertyValue(data[offset:])
		if err != nil {
			return result, fmt.Errorf(
				"decode COV property value: %w",
				err,
			)
		}

		offset += n

		result.Values = append(
			result.Values,
			propertyValue,
		)
	}

	if offset != len(data) {
		return result, fmt.Errorf(
			"unexpected trailing data in COV notification: %d bytes",
			len(data)-offset,
		)
	}

	return result, nil
}

func encodeCOVPropertyValue(
	data []byte,
	value bactypes.PropertyValue,
) int {
	offset := 0

	// [0] propertyIdentifier
	offset += codec.EncodeContextTaggedEnum(
		data[offset:],
		0,
		value.PropertyIdentifier,
	)

	// [1] propertyArrayIndex OPTIONAL
	if value.ArrayIndex != nil {
		offset += codec.EncodeContextTaggedUnsigned(
			data[offset:],
			1,
			*value.ArrayIndex,
		)
	}

	// [2] value
	offset += codec.EncodeOpeningTag(
		data[offset:],
		2,
	)

	copy(
		data[offset:],
		value.EncodedValue,
	)
	offset += len(value.EncodedValue)

	offset += codec.EncodeClosingTag(
		data[offset:],
		2,
	)

	// [3] priority OPTIONAL
	if value.Priority != nil {
		offset += codec.EncodeContextTaggedUnsigned(
			data[offset:],
			3,
			uint32(*value.Priority),
		)
	}

	return offset
}

func decodeCOVPropertyValue(
	data []byte,
) (int, bactypes.PropertyValue, error) {
	var result bactypes.PropertyValue
	offset := 0

	// [0] propertyIdentifier
	n, value, err := decodeContextUnsigned(data[offset:], 0)
	if err != nil {
		return 0, result, fmt.Errorf(
			"decode property identifier: %w",
			err,
		)
	}
	offset += n
	result.PropertyIdentifier = value

	// [1] propertyArrayIndex OPTIONAL
	if offset < len(data) &&
		codec.IsContextTag(data[offset:], 1) {

		n, value, err := decodeContextUnsigned(data[offset:], 1)
		if err != nil {
			return 0, result, fmt.Errorf(
				"decode property array index: %w",
				err,
			)
		}

		offset += n
		result.ArrayIndex = new(uint32)
		*result.ArrayIndex = value
	}

	// [2] value
	n, encodedValue, err :=
		decodeConstructedValue(data[offset:], 2)
	if err != nil {
		return 0, result, fmt.Errorf(
			"decode property value: %w",
			err,
		)
	}
	offset += n
	result.EncodedValue = encodedValue

	// [3] priority OPTIONAL
	if offset < len(data) &&
		codec.IsContextTag(data[offset:], 3) {

		n, value, err := decodeContextUnsigned(data[offset:], 3)
		if err != nil {
			return 0, result, fmt.Errorf(
				"decode property priority: %w",
				err,
			)
		}

		if value < 1 || value > 16 {
			return 0, result, fmt.Errorf(
				"invalid property priority: %d",
				value,
			)
		}

		offset += n

		priority := uint8(value)
		result.Priority = &priority
	}

	return offset, result, nil
}

// decodeConstructedValue returns the encoded contents enclosed by
// an opening/closing context tag pair.
//
// The returned EncodedValue does not include the outer opening and
// closing tags. Nested constructed values are preserved.
func decodeConstructedValue(
	data []byte,
	tagNumber byte,
) (int, []byte, error) {
	if !codec.IsOpeningTagNumber(data, tagNumber) {
		return 0, nil, fmt.Errorf(
			"missing opening tag %d",
			tagNumber,
		)
	}

	tagLength, actualTag := codec.DecodeTagNumber(data)
	if tagLength <= 0 {
		return 0, nil, fmt.Errorf(
			"invalid opening tag %d",
			tagNumber,
		)
	}

	offset := tagLength
	valueStart := offset

	tagStack := []byte{actualTag}

	for offset < len(data) {
		current := data[offset:]

		if codec.IsOpeningTag(current) {
			tagLength, actualTag = codec.DecodeTagNumber(current)
			if tagLength <= 0 {
				return 0, nil, fmt.Errorf(
					"invalid nested opening tag",
				)
			}

			tagStack = append(
				tagStack,
				actualTag,
			)

			offset += tagLength
			continue
		}

		if codec.IsClosingTag(current) {
			tagLength, actualTag = codec.DecodeTagNumber(current)
			if tagLength <= 0 {
				return 0, nil, fmt.Errorf(
					"invalid closing tag",
				)
			}

			if len(tagStack) == 0 ||
				tagStack[len(tagStack)-1] != actualTag {

				return 0, nil, fmt.Errorf(
					"mismatched closing tag %d",
					actualTag,
				)
			}

			tagStack = tagStack[:len(tagStack)-1]

			if len(tagStack) == 0 {
				encodedValue := append(
					[]byte(nil),
					data[valueStart:offset]...,
				)

				return offset + tagLength, encodedValue, nil
			}

			offset += tagLength
			continue
		}

		tagLength, applicationTag, valueLength :=
			codec.DecodeTagNumberAndValue(current)

		if tagLength <= 0 {
			return 0, nil, fmt.Errorf(
				"invalid BACnet value tag",
			)
		}

		payloadLength := int(valueLength)

		//
		// Application-tagged Boolean stores the Boolean value
		// in the tag's LVT field and has no contents octet.
		//
		if !codec.IsContextSpecific(current[0]) &&
			applicationTag == defs.ApplicationTagBoolean {

			payloadLength = 0
		}

		if payloadLength > len(current)-tagLength {
			return 0, nil, fmt.Errorf(
				"truncated BACnet value",
			)
		}

		offset += tagLength + payloadLength
	}

	return 0, nil, fmt.Errorf(
		"missing closing tag %d",
		tagNumber,
	)
}
