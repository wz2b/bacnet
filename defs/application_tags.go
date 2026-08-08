package defs

// Application tags constants.
const (
	ApplicationTagNull byte = 0x00
	ApplicationTagBoolean byte = 0x01
	ApplicationTagUnsignedInt byte = 0x02
	ApplicationTagSignedInt byte = 0x03
	ApplicationTagReal byte = 0x04
	ApplicationTagDouble byte = 0x05
	ApplicationTagOctetString byte = 0x06
	ApplicationTagCharacterString byte = 0x07
	ApplicationTagBitString byte = 0x08
	ApplicationTagEnumerated byte = 0x09
	ApplicationTagDate byte = 0x0A
	ApplicationTagTime byte = 0x0B
	ApplicationTagObjectID byte = 0x0C
	MaxApplicationTag byte = 0x10
)
