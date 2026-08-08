package defs

type CharacterStringEncoding byte

const (
	CharacterANSIX34  CharacterStringEncoding = 0x00
	CharacterUTF8     CharacterStringEncoding = 0x00
	CharacterMSDBCS   CharacterStringEncoding = 0x01
	CharacterJISC6226 CharacterStringEncoding = 0x02
	CharacterUCS4     CharacterStringEncoding = 0x03
	CharacterUCS2     CharacterStringEncoding = 0x04
	CharacterISO8859  CharacterStringEncoding = 0x05
)
