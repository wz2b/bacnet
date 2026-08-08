package defs

type VTClass uint16

// Vt class constants.
const (
	VTClassDefault        VTClass = 0x00
	VTClassProprietaryMin VTClass = 0x40
	VTClassProprietaryMax VTClass = 0xFFFF
)
