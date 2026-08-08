package codec

import (
	"github.com/wz2b/bacnet/bactypes"
	"github.com/wz2b/bacnet/defs"
)

/* from clause 20.2.14 Encoding of an Object Identifier Value */
/* returns the number of apdu bytes consumed */
func EncodeObjectID(apdu []byte, objectID bactypes.ObjectID) int {
	value :=
		(uint32(objectID.Type)&uint32(defs.MaxObject))<<uint32(defs.InstanceBits) |
			(objectID.Instance & uint32(defs.MaxInstance))

	return EncodeUnsigned32(apdu, value)
}

/* from clause 20.2.14 Encoding of an Object Identifier Value */
/* and 20.2.1 General Rules for Encoding BACnet Tags */
/* returns the number of apdu bytes consumed */
func EncodeApplicationTaggedObjectID(
	apdu []byte,
	objectID bactypes.ObjectID,
) int {
	tagLength := EncodeTag(
		apdu,
		defs.ApplicationTagObjectID,
		false,
		4,
	)

	valueLength := EncodeObjectID(
		apdu[tagLength:],
		objectID,
	)

	return tagLength + valueLength
}

/* from clause 20.2.14 Encoding of an Object Identifier Value */
/* returns the number of apdu bytes consumed */
func DecodeObjectID(apdu []byte) (int, defs.ObjectType, uint32) {
	length, value := DecodeUnsigned32(apdu)

	objectType := defs.ObjectType(
		(value >> uint32(defs.InstanceBits)) &
			uint32(defs.MaxObject),
	)

	instance := value & uint32(defs.MaxInstance)

	return length, objectType, instance
}

/* from clause 20.2.14 Encoding of an Object Identifier Value */
/* and 20.2.1 General Rules for Encoding BACnet Tags */
/* returns the number of apdu bytes consumed */
func EncodeContextTaggedObjectID(
	apdu []byte,
	tagNumber defs.ApplicationTagType,
	objectID bactypes.ObjectID,
) int {
	length := EncodeTag(
		apdu,
		tagNumber,
		true,
		4,
	)

	length += EncodeObjectID(
		apdu[length:],
		objectID,
	)

	return length
}
