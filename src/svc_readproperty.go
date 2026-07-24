package bacnet

import (
	"encoding/binary"
	"fmt"
	"math"
)

func NewReadPropertyACKEncoded(
	pdu []byte,
	invokeID byte,
	objectType uint16,
	objectInstance uint32,
	propertyIdentifier uint32,
	arrayIndex *uint32,
	encodedValue []byte,
) *ReadPropertyACK {
	return &ReadPropertyACK{
		PDU:                pdu,
		InvokeID:           invokeID,
		ObjectType:         objectType,
		ObjectInstance:     objectInstance,
		PropertyIdentifier: propertyIdentifier,
		ArrayIndex:         arrayIndex,
		EncodedValue:       append([]byte(nil), encodedValue...),
	}
}

// ReadPropertyRequest is the service-specific portion of a BACnet
// ReadProperty confirmed request.
//
// PDU contains only the service data, not the confirmed-request APDU header.
type ReadPropertyRequest struct {
	PDU []byte

	ObjectType     uint16
	ObjectInstance uint32

	PropertyIdentifier uint32
	ArrayIndex         *uint32

	Length int
}

func NewReadPropertyRequest(serviceData []byte) *ReadPropertyRequest {
	return &ReadPropertyRequest{
		PDU: serviceData,
	}
}

// Decode decodes:
//
//	[0] objectIdentifier
//	[1] propertyIdentifier
//	[2] propertyArrayIndex OPTIONAL
func (r *ReadPropertyRequest) Decode() (*ReadPropertyRequest, error) {
	if r == nil {
		return nil, fmt.Errorf("decode ReadProperty: nil request")
	}

	data := r.PDU
	offset := 0

	// [0] objectIdentifier
	tag, valueBytes, consumed, err := decodeContextValue(data[offset:])
	if err != nil {
		return nil, fmt.Errorf(
			"decode ReadProperty object identifier: %w",
			err,
		)
	}

	if tag != 0 {
		return nil, fmt.Errorf(
			"decode ReadProperty: expected context tag 0, got %d",
			tag,
		)
	}

	if len(valueBytes) != 4 {
		return nil, fmt.Errorf(
			"decode ReadProperty: object identifier is %d bytes, expected 4",
			len(valueBytes),
		)
	}

	objectID := binary.BigEndian.Uint32(valueBytes)

	r.ObjectType = uint16(
		(objectID >> uint32(BACNET_INSTANCE_BITS)) &
			uint32(BACNET_MAX_OBJECT),
	)

	r.ObjectInstance =
		objectID & uint32(BACNET_MAX_INSTANCE)

	offset += consumed

	// [1] propertyIdentifier
	tag, valueBytes, consumed, err =
		decodeContextValue(data[offset:])
	if err != nil {
		return nil, fmt.Errorf(
			"decode ReadProperty property identifier: %w",
			err,
		)
	}

	if tag != 1 {
		return nil, fmt.Errorf(
			"decode ReadProperty: expected context tag 1, got %d",
			tag,
		)
	}

	propertyID, err := decodeUnsignedValue(valueBytes)
	if err != nil {
		return nil, fmt.Errorf(
			"decode ReadProperty property identifier: %w",
			err,
		)
	}

	r.PropertyIdentifier = propertyID
	offset += consumed

	// [2] propertyArrayIndex OPTIONAL
	if offset < len(data) {
		tag, valueBytes, consumed, err =
			decodeContextValue(data[offset:])
		if err != nil {
			return nil, fmt.Errorf(
				"decode ReadProperty array index: %w",
				err,
			)
		}

		if tag != 2 {
			return nil, fmt.Errorf(
				"decode ReadProperty: expected context tag 2, got %d",
				tag,
			)
		}

		arrayIndex, err := decodeUnsignedValue(valueBytes)
		if err != nil {
			return nil, fmt.Errorf(
				"decode ReadProperty array index: %w",
				err,
			)
		}

		r.ArrayIndex = &arrayIndex
		offset += consumed
	}

	if offset != len(data) {
		return nil, fmt.Errorf(
			"decode ReadProperty: %d trailing bytes",
			len(data)-offset,
		)
	}

	r.Length = offset
	return r, nil
}

// ReadPropertyACK represents an unsegmented ComplexACK carrying one
// ReadProperty result.
//
// EncodedValue, when non-empty, must contain exactly one BACnet
// application-tagged value. It must not include the surrounding
// context tag 3 opening and closing tags.
//
// RealValue is retained for backward compatibility. It is used only
// when EncodedValue is empty.
type ReadPropertyACK struct {
	PDU []byte

	InvokeID byte

	ObjectType     uint16
	ObjectInstance uint32

	PropertyIdentifier uint32
	ArrayIndex         *uint32

	// Encoded BACnet application-tagged value.
	EncodedValue []byte

	// Retained temporarily for compatibility.
	RealValue float32

	Length int
}

func NewReadPropertyACK(
	pdu []byte,
	invokeID byte,
	objectType uint16,
	objectInstance uint32,
	propertyIdentifier uint32,
	arrayIndex *uint32,
	value float32,
) *ReadPropertyACK {
	if pdu == nil {
		pdu = make([]byte, 100)
	}

	return &ReadPropertyACK{
		PDU:                pdu,
		InvokeID:           invokeID,
		ObjectType:         objectType,
		ObjectInstance:     objectInstance,
		PropertyIdentifier: propertyIdentifier,
		ArrayIndex:         arrayIndex,
		RealValue:          value,
	}
}

// Encode produces:
//
//	ComplexACK
//	invoke ID
//	ReadProperty service choice
//	[0] objectIdentifier
//	[1] propertyIdentifier
//	[2] propertyArrayIndex OPTIONAL
//	[3] propertyValue opening tag
//	    application-tagged value
//	[3] propertyValue closing tag
func (a *ReadPropertyACK) Encode() *ReadPropertyACK {
	offset := 0

	requiredCapacity := 64 + len(a.EncodedValue)

	if len(a.PDU) < requiredCapacity {
		newPDU := make([]byte, requiredCapacity)
		copy(newPDU, a.PDU)
		a.PDU = newPDU
	}

	a.PDU[offset] = PDU_TYPE_COMPLEX_ACK
	offset++

	a.PDU[offset] = a.InvokeID
	offset++

	a.PDU[offset] = SERVICE_CONFIRMED_READ_PROPERTY
	offset++

	offset += encode_context_object_id(
		a.PDU[offset:],
		0,
		int(a.ObjectType),
		a.ObjectInstance,
	)

	offset += encode_context_enumerated(
		a.PDU[offset:],
		1,
		a.PropertyIdentifier,
	)

	if a.ArrayIndex != nil {
		offset += encode_context_unsigned(
			a.PDU[offset:],
			2,
			*a.ArrayIndex,
		)
	}

	offset += encode_opening_tag(
		a.PDU[offset:],
		3,
	)

	if len(a.EncodedValue) > 0 {
		copy(a.PDU[offset:], a.EncodedValue)
		offset += len(a.EncodedValue)
	} else {
		offset += encodeApplicationReal(
			a.PDU[offset:],
			a.RealValue,
		)
	}

	offset += encode_closing_tag(
		a.PDU[offset:],
		3,
	)

	a.PDU = a.PDU[:offset]
	a.Length = offset

	return a
}

// decodeContextValue decodes one primitive context-tagged value and returns
// the context tag number, the encoded value bytes, and total bytes consumed.
func decodeContextValue(
	data []byte,
) (
	tagNumber byte,
	value []byte,
	consumed int,
	err error,
) {
	if len(data) < 1 {
		return 0, nil, 0, fmt.Errorf("missing context tag")
	}

	first := data[0]

	// Bit 3 distinguishes context-specific from application tags.
	if first&0x08 == 0 {
		return 0, nil, 0, fmt.Errorf(
			"expected context-specific tag, got 0x%02X",
			first,
		)
	}

	offset := 1
	tagNumber = first >> 4

	// Extended tag number.
	if tagNumber == 0x0F {
		if len(data) < offset+1 {
			return 0, nil, 0, fmt.Errorf(
				"truncated extended tag number",
			)
		}

		tagNumber = data[offset]
		offset++
	}

	lengthValueType := first & 0x07

	switch lengthValueType {
	case 6:
		return 0, nil, 0, fmt.Errorf(
			"opening tag is not a primitive value",
		)

	case 7:
		return 0, nil, 0, fmt.Errorf(
			"closing tag is not a primitive value",
		)
	}

	valueLength := uint32(lengthValueType)

	// Extended value length.
	if lengthValueType == 5 {
		if len(data) < offset+1 {
			return 0, nil, 0, fmt.Errorf(
				"truncated extended value length",
			)
		}

		extended := data[offset]
		offset++

		switch extended {
		case 254:
			if len(data) < offset+2 {
				return 0, nil, 0, fmt.Errorf(
					"truncated 16-bit value length",
				)
			}

			valueLength =
				uint32(binary.BigEndian.Uint16(data[offset:]))

			offset += 2

		case 255:
			if len(data) < offset+4 {
				return 0, nil, 0, fmt.Errorf(
					"truncated 32-bit value length",
				)
			}

			valueLength =
				binary.BigEndian.Uint32(data[offset:])

			offset += 4

		default:
			valueLength = uint32(extended)
		}
	}

	end := offset + int(valueLength)

	if end > len(data) {
		return 0, nil, 0, fmt.Errorf(
			"tag value truncated: need %d bytes, have %d",
			valueLength,
			len(data)-offset,
		)
	}

	return tagNumber, data[offset:end], end, nil
}

func decodeUnsignedValue(data []byte) (uint32, error) {
	switch len(data) {
	case 1:
		return uint32(data[0]), nil

	case 2:
		return uint32(binary.BigEndian.Uint16(data)), nil

	case 3:
		return uint32(data[0])<<16 |
			uint32(data[1])<<8 |
			uint32(data[2]), nil

	case 4:
		return binary.BigEndian.Uint32(data), nil

	default:
		return 0, fmt.Errorf(
			"unsigned value has invalid length %d",
			len(data),
		)
	}
}

// encodeApplicationReal encodes a BACnet application-tagged REAL.
//
// A BACnet REAL is a four-byte IEEE-754 single-precision value in network
// byte order. The first byte is application tag 4 with a value length of 4.
func encodeApplicationReal(
	apdu []byte,
	value float32,
) int {
	apdu[0] =
		(BACNET_APPLICATION_TAG_REAL << 4) | 4

	binary.BigEndian.PutUint32(
		apdu[1:5],
		math.Float32bits(value),
	)

	return 5
}
