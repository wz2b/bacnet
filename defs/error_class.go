package defs

// Error class constants.
const (
	ErrorClassDevice uint16 = 0x00
	ErrorClassObject uint16 = 0x01
	ErrorClassProperty uint16 = 0x02
	ErrorClassResources uint16 = 0x03
	ErrorClassSecurity uint16 = 0x04
	ErrorClassServices uint16 = 0x05
	ErrorClassVT uint16 = 0x06
	ErrorClassCommunication uint16 = 0x07
	MaxErrorClass uint16 = 0x08
	ErrorClassProprietaryFirst uint16 = 0x40
	ErrorClassProprietaryLast uint16 = 0xFFFF
)
