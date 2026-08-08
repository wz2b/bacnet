// Package bacnet
// File bac_property.go
package bacnet

type PropertyReference struct {
	PropertyIdentifier uint32
	ArrayIndex         *uint32
}

type PropertyValue struct {
	PropertyIdentifier uint32
	ArrayIndex         *uint32

	// Complete BACnet application-tagged value.
	//
	// For example:
	//   REAL      -> 0x44 followed by four IEEE-754 bytes
	//   BIT STRING -> application tag plus encoded bit-string content
	EncodedValue []byte

	// Used by services that support an optional write priority.
	// It will ordinarily be nil for COV notifications.
	Priority *uint8
}
