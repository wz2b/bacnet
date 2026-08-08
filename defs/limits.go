package defs

// todo: start find right value for these
const (
	MaxAPDU                 = 1476
	MaxCharacterStringBytes = MaxAPDU - 6
	CharacterStringCapacity = MaxCharacterStringBytes - 1
)
