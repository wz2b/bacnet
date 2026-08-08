package defs

// Silenced state constants.
const (
	SilencedStateUnsilenced uint16 = 0x00
	SilencedStateAudibleSilenced uint16 = 0x01
	SilencedStateVisibleSilenced uint16 = 0x02
	SilencedStateAllSilenced uint16 = 0x03
	SilencedStateProprietaryMin uint16 = 0x40
	SilencedStateProprietaryMax uint16 = 0xFFFF
)
