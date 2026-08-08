package defs

type ErrorClass uint16

// Error class constants.
const (
	ErrorClassDevice           ErrorClass = 0x00
	ErrorClassObject           ErrorClass = 0x01
	ErrorClassProperty         ErrorClass = 0x02
	ErrorClassResources        ErrorClass = 0x03
	ErrorClassSecurity         ErrorClass = 0x04
	ErrorClassServices         ErrorClass = 0x05
	ErrorClassVT               ErrorClass = 0x06
	ErrorClassCommunication    ErrorClass = 0x07
	MaxErrorClass              ErrorClass = 0x08
	ErrorClassProprietaryFirst ErrorClass = 0x40
	ErrorClassProprietaryLast  ErrorClass = 0xFFFF
)
