package defs

type SilenceState uint16

// Silenced state constants.
const (
	SilencedStateUnsilenced      SilenceState = 0x00
	SilencedStateAudibleSilenced SilenceState = 0x01
	SilencedStateVisibleSilenced SilenceState = 0x02
	SilencedStateAllSilenced     SilenceState = 0x03
	SilencedStateProprietaryMin  SilenceState = 0x40
	SilencedStateProprietaryMax  SilenceState = 0xFFFF
)
