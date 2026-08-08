package services

import (
	"testing"

	"github.com/wz2b/bacnet/bactypes"
	"github.com/wz2b/bacnet/codec"
	"github.com/wz2b/bacnet/defs"
)

func TestACKAlarmEncodingAndDecoding(t *testing.T) {
	original := ACKAlarm{
		AckProcessIdentifier: 1234,
		EventObjectIdentifier: bactypes.ObjectID{
			Type:     defs.ObjectType(2),
			Instance: 1001,
		},
		EventStateAcked: 1,
		AckSource: codec.CharacterString{
			Value: []byte("operator"),
		},
	}

	a := original.APDU()

	decoded, err := DecodeACKAlarm(&a)
	if err != nil {
		t.Fatalf("DecodeACKAlarm() error: %v", err)
	}

	if decoded.AckProcessIdentifier != original.AckProcessIdentifier {
		t.Errorf(
			"AckProcessIdentifier: got %d, want %d",
			decoded.AckProcessIdentifier,
			original.AckProcessIdentifier,
		)
	}

	if decoded.EventObjectIdentifier != original.EventObjectIdentifier {
		t.Errorf(
			"EventObjectIdentifier: got %+v, want %+v",
			decoded.EventObjectIdentifier,
			original.EventObjectIdentifier,
		)
	}

	if decoded.EventStateAcked != original.EventStateAcked {
		t.Errorf(
			"EventStateAcked: got %d, want %d",
			decoded.EventStateAcked,
			original.EventStateAcked,
		)
	}
}
