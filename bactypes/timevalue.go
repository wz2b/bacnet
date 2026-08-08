package bactypes

type TimeValue struct {
	Time Time

	// Complete BACnet application-tagged value.
	EncodedValue []byte
}
