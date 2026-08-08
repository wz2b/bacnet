package services

import (
	"bytes"
	"testing"

	"github.com/wz2b/bacnet/apdu"
	"github.com/wz2b/bacnet/bactypes"
)

func TestReadPropertyRequestEncoding(t *testing.T) {
	msg := ReadPropertyRequest{
		Object: bactypes.ObjectID{
			Type:     bactypes.ObjectType(2), // Analog Value
			Instance: 1234,
		},
		PropertyIdentifier: 85, // Present Value
	}

	a := msg.APDU()

	a.InvokeID = 42
	a.MaxAPDU = 1476
	a.MaxSegments = 0

	got, err := apdu.Encode(a)
	if err != nil {
		t.Fatalf("apdu.Encode() error: %v", err)
	}

	want := []byte{
		// Confirmed-Request
		0x00,

		// Max Segments = unspecified, Max APDU = 1476
		0x05,

		// Invoke ID = 42
		0x2A,

		// ReadProperty service choice = 12
		0x0C,

		// [0] objectIdentifier = AnalogValue:1234
		//
		// (2 << 22) | 1234 = 0x008004D2
		0x0C,
		0x00, 0x80, 0x04, 0xD2,

		// [1] propertyIdentifier = Present Value (85)
		0x19, 0x55,
	}

	if !bytes.Equal(got, want) {
		t.Errorf(
			"encoded ReadPropertyRequest mismatch:\n got: % X\nwant: % X",
			got,
			want,
		)
	}
}

func TestReadPropertyRequestDecoding(t *testing.T) {
	data := []byte{
		// Confirmed-Request
		0x00,

		// Max Segments = unspecified, Max APDU = 1476
		0x05,

		// Invoke ID = 42
		0x2A,

		// ReadProperty service choice = 12
		0x0C,

		// [0] objectIdentifier = AnalogValue:1234
		0x0C,
		0x00, 0x80, 0x04, 0xD2,

		// [1] propertyIdentifier = Present Value (85)
		0x19, 0x55,
	}

	a, err := apdu.Decode(data)
	if err != nil {
		t.Fatalf("apdu.Decode() error: %v", err)
	}

	if a.InvokeID != 42 {
		t.Errorf(
			"InvokeID: got %d, want %d",
			a.InvokeID,
			42,
		)
	}

	if a.MaxAPDU != 1476 {
		t.Errorf(
			"MaxAPDU: got %d, want %d",
			a.MaxAPDU,
			1476,
		)
	}

	decoded, err := DecodeReadPropertyRequest(&a)
	if err != nil {
		t.Fatalf("DecodeReadPropertyRequest() error: %v", err)
	}

	wantObject := bactypes.ObjectID{
		Type:     bactypes.ObjectType(2),
		Instance: 1234,
	}

	if decoded.Object != wantObject {
		t.Errorf(
			"Object: got %+v, want %+v",
			decoded.Object,
			wantObject,
		)
	}

	if decoded.PropertyIdentifier != 85 {
		t.Errorf(
			"PropertyIdentifier: got %d, want %d",
			decoded.PropertyIdentifier,
			85,
		)
	}

	if decoded.ArrayIndex != nil {
		t.Errorf(
			"ArrayIndex: got %d, want nil",
			*decoded.ArrayIndex,
		)
	}
}
