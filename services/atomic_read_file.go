// atomic_read_file.go

package services

import (
	"encoding/binary"
	"fmt"

	"github.com/wz2b/bacnet/apdu"
	"github.com/wz2b/bacnet/bactypes"
	"github.com/wz2b/bacnet/codec"
	"github.com/wz2b/bacnet/defs"
)

// AtomicReadFileRequest represents the stream-access form of an
// AtomicReadFile request.
//
// Record access can be added later if we need it.
type AtomicReadFileRequest struct {
	Object              bactypes.ObjectID
	StartPosition       int32
	RequestedOctetCount uint32
}

// AtomicReadFileACK represents the stream-access form of an
// AtomicReadFile acknowledgement.
type AtomicReadFileACK struct {
	EndOfFile     bool
	StartPosition int32
	Data          []byte
}

func (r AtomicReadFileRequest) APDU() apdu.APDU {
	data := make([]byte, 32)

	offset := 0

	//
	// file-identifier
	//
	offset += codec.EncodeApplicationTaggedObjectID(
		data[offset:],
		r.Object,
	)

	//
	// stream-access [0]
	//
	offset += codec.EncodeOpeningTag(
		data[offset:],
		0,
	)

	//
	// file-start-position
	//
	offset += encodeApplicationTaggedSigned(
		data[offset:],
		r.StartPosition,
	)

	//
	// requested-octet-count
	//
	unsigned := codec.EncodeApplicationTaggedUnsigned(
		r.RequestedOctetCount,
	)

	copy(
		data[offset:],
		unsigned,
	)
	offset += len(unsigned)

	offset += codec.EncodeClosingTag(
		data[offset:],
		0,
	)

	return apdu.APDU{
		Type:                   defs.PDUTypeConfirmedServiceRequest,
		ConfirmedServiceChoice: defs.ServiceConfirmedAtomicReadFile,
		Data:                   data[:offset],
	}
}

func (a AtomicReadFileACK) APDU() apdu.APDU {
	data := make(
		[]byte,
		32+len(a.Data),
	)

	offset := 0

	//
	// end-of-file BOOLEAN
	//
	data[offset] =
		byte(defs.ApplicationTagBoolean) << 4

	if a.EndOfFile {
		data[offset] |= 1
	}
	offset++

	//
	// stream-access [0]
	//
	offset += codec.EncodeOpeningTag(
		data[offset:],
		0,
	)

	offset += encodeApplicationTaggedSigned(
		data[offset:],
		a.StartPosition,
	)

	//
	// file-data OCTET STRING
	//
	offset += codec.EncodeTag(
		data[offset:],
		defs.ApplicationTagOctetString,
		false,
		uint32(len(a.Data)),
	)

	copy(
		data[offset:],
		a.Data,
	)
	offset += len(a.Data)

	offset += codec.EncodeClosingTag(
		data[offset:],
		0,
	)

	return apdu.APDU{
		Type:                   defs.PDUTypeComplexACK,
		ConfirmedServiceChoice: defs.ServiceConfirmedAtomicReadFile,
		Data:                   data[:offset],
	}
}

func DecodeAtomicReadFileRequest(
	a *apdu.APDU,
) (*AtomicReadFileRequest, error) {
	if a == nil {
		return nil, fmt.Errorf(
			"nil APDU",
		)
	}

	if a.Type != defs.PDUTypeConfirmedServiceRequest {
		return nil, fmt.Errorf(
			"AtomicReadFile requires a confirmed service request APDU",
		)
	}

	if a.ConfirmedServiceChoice !=
		defs.ServiceConfirmedAtomicReadFile {
		return nil, fmt.Errorf(
			"APDU is not an AtomicReadFile request",
		)
	}

	result := &AtomicReadFileRequest{}

	data := a.Data
	offset := 0

	//
	// file-identifier
	//
	object, n, err :=
		decodeApplicationTaggedObjectID(
			data[offset:],
		)
	if err != nil {
		return nil, err
	}

	result.Object = object
	offset += n

	//
	// stream-access [0]
	//
	n, err = decodeOpeningTag(
		data[offset:],
		0,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"AtomicReadFile stream access: %w",
			err,
		)
	}
	offset += n

	//
	// file-start-position
	//
	position, n, err :=
		decodeApplicationTaggedSigned(
			data[offset:],
		)
	if err != nil {
		return nil, fmt.Errorf(
			"AtomicReadFile start position: %w",
			err,
		)
	}

	result.StartPosition = position
	offset += n

	//
	// requested-octet-count
	//
	count, n, err :=
		decodeApplicationTaggedUnsigned(
			data[offset:],
		)
	if err != nil {
		return nil, fmt.Errorf(
			"AtomicReadFile octet count: %w",
			err,
		)
	}

	result.RequestedOctetCount = count
	offset += n

	n, err = decodeClosingTag(
		data[offset:],
		0,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"AtomicReadFile stream access: %w",
			err,
		)
	}
	offset += n

	if offset != len(data) {
		return nil, fmt.Errorf(
			"unexpected trailing AtomicReadFile request data: %d bytes",
			len(data)-offset,
		)
	}

	return result, nil
}

func DecodeAtomicReadFileACK(
	a *apdu.APDU,
) (*AtomicReadFileACK, error) {
	if a == nil {
		return nil, fmt.Errorf(
			"nil APDU",
		)
	}

	if a.Type != defs.PDUTypeComplexACK {
		return nil, fmt.Errorf(
			"AtomicReadFile ACK requires a ComplexACK APDU",
		)
	}

	if a.ConfirmedServiceChoice !=
		defs.ServiceConfirmedAtomicReadFile {
		return nil, fmt.Errorf(
			"APDU is not an AtomicReadFile ACK",
		)
	}

	result := &AtomicReadFileACK{}

	data := a.Data
	offset := 0

	//
	// end-of-file BOOLEAN
	//
	if offset >= len(data) {
		return nil, fmt.Errorf(
			"AtomicReadFile ACK missing end-of-file",
		)
	}

	if codec.IsContextSpecific(data[offset]) {
		return nil, fmt.Errorf(
			"AtomicReadFile ACK end-of-file is context tagged",
		)
	}

	tagLength, tag, value :=
		codec.DecodeTagNumberAndValue(
			data[offset:],
		)

	if tagLength <= 0 ||
		tag != defs.ApplicationTagBoolean ||
		value > 1 {
		return nil, fmt.Errorf(
			"invalid AtomicReadFile ACK end-of-file",
		)
	}

	result.EndOfFile = value != 0
	offset += tagLength

	//
	// stream-access [0]
	//
	n, err := decodeOpeningTag(
		data[offset:],
		0,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"AtomicReadFile ACK stream access: %w",
			err,
		)
	}
	offset += n

	//
	// file-start-position
	//
	position, n, err :=
		decodeApplicationTaggedSigned(
			data[offset:],
		)
	if err != nil {
		return nil, fmt.Errorf(
			"AtomicReadFile ACK start position: %w",
			err,
		)
	}

	result.StartPosition = position
	offset += n

	//
	// file-data OCTET STRING
	//
	if offset >= len(data) {
		return nil, fmt.Errorf(
			"AtomicReadFile ACK missing file data",
		)
	}

	if codec.IsContextSpecific(data[offset]) {
		return nil, fmt.Errorf(
			"AtomicReadFile ACK file data is context tagged",
		)
	}

	tagLength, tag, valueLength :=
		codec.DecodeTagNumberAndValue(
			data[offset:],
		)

	if tagLength <= 0 ||
		tag != defs.ApplicationTagOctetString {
		return nil, fmt.Errorf(
			"invalid AtomicReadFile ACK file data",
		)
	}

	offset += tagLength

	if int(valueLength) > len(data)-offset {
		return nil, fmt.Errorf(
			"truncated AtomicReadFile ACK file data",
		)
	}

	result.Data = append(
		[]byte(nil),
		data[offset:offset+int(valueLength)]...,
	)

	offset += int(valueLength)

	n, err = decodeClosingTag(
		data[offset:],
		0,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"AtomicReadFile ACK stream access: %w",
			err,
		)
	}
	offset += n

	if offset != len(data) {
		return nil, fmt.Errorf(
			"unexpected trailing AtomicReadFile ACK data: %d bytes",
			len(data)-offset,
		)
	}

	return result, nil
}

func encodeApplicationTaggedSigned(
	data []byte,
	value int32,
) int {
	var valueLength int

	switch {
	case value >= -128 && value <= 127:
		valueLength = 1

	case value >= -32768 && value <= 32767:
		valueLength = 2

	case value >= -8388608 && value <= 8388607:
		valueLength = 3

	default:
		valueLength = 4
	}

	offset := codec.EncodeTag(
		data,
		defs.ApplicationTagSignedInt,
		false,
		uint32(valueLength),
	)

	switch valueLength {
	case 1:
		offset += codec.EncodeSigned8(
			data[offset:],
			int8(value),
		)

	case 2:
		offset += codec.EncodeSigned16(
			data[offset:],
			value,
		)

	case 3:
		offset += codec.EncodeSigned24(
			data[offset:],
			value,
		)

	case 4:
		offset += codec.EncodeSigned32(
			data[offset:],
			int64(value),
		)
	}

	return offset
}

func decodeApplicationTaggedSigned(
	data []byte,
) (int32, int, error) {
	if len(data) == 0 {
		return 0, 0, fmt.Errorf(
			"missing signed integer",
		)
	}

	if codec.IsContextSpecific(data[0]) {
		return 0, 0, fmt.Errorf(
			"signed integer is context tagged",
		)
	}

	tagLength, tag, valueLength :=
		codec.DecodeTagNumberAndValue(data)

	if tagLength <= 0 ||
		tag != defs.ApplicationTagSignedInt ||
		valueLength < 1 ||
		valueLength > 4 {
		return 0, 0, fmt.Errorf(
			"invalid signed integer",
		)
	}

	if int(valueLength) > len(data)-tagLength {
		return 0, 0, fmt.Errorf(
			"truncated signed integer",
		)
	}

	valueData :=
		data[tagLength : tagLength+int(valueLength)]

	var value int32

	switch valueLength {
	case 1:
		value = int32(
			int8(valueData[0]),
		)

	case 2:
		value = int32(
			int16(
				binary.BigEndian.Uint16(
					valueData,
				),
			),
		)

	case 3:
		u :=
			uint32(valueData[0])<<16 |
				uint32(valueData[1])<<8 |
				uint32(valueData[2])

		if u&0x00800000 != 0 {
			u |= 0xFF000000
		}

		value = int32(u)

	case 4:
		value = int32(
			binary.BigEndian.Uint32(
				valueData,
			),
		)
	}

	return value,
		tagLength + int(valueLength),
		nil
}

func decodeApplicationTaggedUnsigned(
	data []byte,
) (uint32, int, error) {
	if len(data) == 0 {
		return 0, 0, fmt.Errorf(
			"missing unsigned integer",
		)
	}

	if codec.IsContextSpecific(data[0]) {
		return 0, 0, fmt.Errorf(
			"unsigned integer is context tagged",
		)
	}

	tagLength, tag, valueLength :=
		codec.DecodeTagNumberAndValue(data)

	if tagLength <= 0 ||
		tag != defs.ApplicationTagUnsignedInt ||
		valueLength < 1 ||
		valueLength > 4 {
		return 0, 0, fmt.Errorf(
			"invalid unsigned integer",
		)
	}

	if int(valueLength) > len(data)-tagLength {
		return 0, 0, fmt.Errorf(
			"truncated unsigned integer",
		)
	}

	n, value := codec.DecodeUnsigned(
		data[tagLength:],
		valueLength,
	)

	if n <= 0 {
		return 0, 0, fmt.Errorf(
			"invalid unsigned integer value",
		)
	}

	return value,
		tagLength + n,
		nil
}

func decodeApplicationTaggedObjectID(
	data []byte,
) (bactypes.ObjectID, int, error) {
	var result bactypes.ObjectID

	if len(data) == 0 {
		return result, 0, fmt.Errorf(
			"missing object identifier",
		)
	}

	if codec.IsContextSpecific(data[0]) {
		return result, 0, fmt.Errorf(
			"object identifier is context tagged",
		)
	}

	tagLength, tag, valueLength :=
		codec.DecodeTagNumberAndValue(data)

	if tagLength <= 0 ||
		tag != defs.ApplicationTagObjectID ||
		valueLength != 4 {
		return result, 0, fmt.Errorf(
			"invalid object identifier",
		)
	}

	if len(data)-tagLength < 4 {
		return result, 0, fmt.Errorf(
			"truncated object identifier",
		)
	}

	n, objectType, instance :=
		codec.DecodeObjectID(
			data[tagLength:],
		)

	if n != 4 {
		return result, 0, fmt.Errorf(
			"invalid object identifier value",
		)
	}

	result = bactypes.ObjectID{
		Type:     objectType,
		Instance: instance,
	}

	return result,
		tagLength + n,
		nil
}

func decodeOpeningTag(
	data []byte,
	tagNumber byte,
) (int, error) {
	if !codec.IsOpeningTagNumber(
		data,
		tagNumber,
	) {
		return 0, fmt.Errorf(
			"missing opening tag %d",
			tagNumber,
		)
	}

	n, _ := codec.DecodeTagNumber(data)

	if n <= 0 {
		return 0, fmt.Errorf(
			"invalid opening tag %d",
			tagNumber,
		)
	}

	return n, nil
}

func decodeClosingTag(
	data []byte,
	tagNumber byte,
) (int, error) {
	if !codec.IsClosingTagNumber(
		data,
		tagNumber,
	) {
		return 0, fmt.Errorf(
			"missing closing tag %d",
			tagNumber,
		)
	}

	n, _ := codec.DecodeTagNumber(data)

	if n <= 0 {
		return 0, fmt.Errorf(
			"invalid closing tag %d",
			tagNumber,
		)
	}

	return n, nil
}
