package defs

type ApplicationTagType byte

// Application tags constants.
const (
	ApplicationTagNull            ApplicationTagType = 0x00
	ApplicationTagBoolean         ApplicationTagType = 0x01
	ApplicationTagUnsignedInt     ApplicationTagType = 0x02
	ApplicationTagSignedInt       ApplicationTagType = 0x03
	ApplicationTagReal            ApplicationTagType = 0x04
	ApplicationTagDouble          ApplicationTagType = 0x05
	ApplicationTagOctetString     ApplicationTagType = 0x06
	ApplicationTagCharacterString ApplicationTagType = 0x07
	ApplicationTagBitString       ApplicationTagType = 0x08
	ApplicationTagEnumerated      ApplicationTagType = 0x09
	ApplicationTagDate            ApplicationTagType = 0x0A
	ApplicationTagTime            ApplicationTagType = 0x0B
	ApplicationTagObjectID        ApplicationTagType = 0x0C
	MaxApplicationTag             ApplicationTagType = 0x10
)
