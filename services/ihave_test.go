package services

import (
	"bytes"
	"testing"

	"github.com/wz2b/bacnet/apdu"
	"github.com/wz2b/bacnet/bactypes"
	"github.com/wz2b/bacnet/codec"
	"github.com/wz2b/bacnet/defs"
)

func TestIHaveRoundTrip(t *testing.T) {
	original := IHave{
		DeviceID: 12345,
		ObjectID: bactypes.ObjectID{
			Type:     defs.ObjectType(0), // Analog Input
			Instance: 42,
		},
		ObjectName: codec.CharacterString{
			Value: []byte("Zone Temperature"),
		},
	}

	a := original.APDU()

	decoded, err := DecodeIHave(&a)
	if err != nil {
		t.Fatalf("DecodeIHave() error: %v", err)
	}

	if decoded.DeviceID != original.DeviceID {
		t.Errorf(
			"DeviceID: got %d, want %d",
			decoded.DeviceID,
			original.DeviceID,
		)
	}

	if decoded.ObjectID != original.ObjectID {
		t.Errorf(
			"ObjectID: got %+v, want %+v",
			decoded.ObjectID,
			original.ObjectID,
		)
	}

	if string(decoded.ObjectName.Value) != string(original.ObjectName.Value) {
		t.Errorf(
			"ObjectName: got %q, want %q",
			decoded.ObjectName.Value,
			original.ObjectName.Value,
		)
	}
}

func TestIHaveEncoding(t *testing.T) {
	msg := IHave{
		DeviceID: 12345,
		ObjectID: bactypes.ObjectID{
			Type:     defs.ObjectType(0), // Analog Input
			Instance: 42,
		},
		ObjectName: codec.CharacterString{
			Value: []byte("Zone Temperature"),
		},
	}

	a := msg.APDU()

	got, err := apdu.Encode(a)
	if err != nil {
		t.Fatalf("apdu.Encode() error: %v", err)
	}

	want := []byte{
		// Unconfirmed-Request, I-Have
		0x10, 0x01,

		// Device Identifier:
		// application ObjectIdentifier, Device:12345
		0xC4,
		0x02, 0x00, 0x30, 0x39,

		// Object Identifier:
		// application ObjectIdentifier, AnalogInput:42
		0xC4,
		0x00, 0x00, 0x00, 0x2A,

		// Object Name:
		// application CharacterString, length 17
		0x75, 0x11,

		// Character encoding: ANSI X3.4
		0x00,

		// "Zone Temperature"
		0x5A, 0x6F, 0x6E, 0x65,
		0x20,
		0x54, 0x65, 0x6D, 0x70,
		0x65, 0x72, 0x61, 0x74,
		0x75, 0x72, 0x65,
	}

	if !bytes.Equal(got, want) {
		t.Errorf(
			"encoded I-Have mismatch:\n got: % X\nwant: % X",
			got,
			want,
		)
	}
}

func TestIHaveDecoding(t *testing.T) {
	data := []byte{
		// Unconfirmed-Request, I-Have
		0x10, 0x01,

		// Device Identifier:
		// application ObjectIdentifier, Device:12345
		0xC4,
		0x02, 0x00, 0x30, 0x39,

		// Object Identifier:
		// application ObjectIdentifier, AnalogInput:42
		0xC4,
		0x00, 0x00, 0x00, 0x2A,

		// Object Name:
		// application CharacterString, length 17
		0x75, 0x11,

		// Character encoding: ANSI X3.4
		0x00,

		// "Zone Temperature"
		0x5A, 0x6F, 0x6E, 0x65,
		0x20,
		0x54, 0x65, 0x6D, 0x70,
		0x65, 0x72, 0x61, 0x74,
		0x75, 0x72, 0x65,
	}

	a, err := apdu.Decode(data)
	if err != nil {
		t.Fatalf("apdu.Decode() error: %v", err)
	}

	decoded, err := DecodeIHave(&a)
	if err != nil {
		t.Fatalf("DecodeIHave() error: %v", err)
	}

	if decoded.DeviceID != 12345 {
		t.Errorf(
			"DeviceID: got %d, want %d",
			decoded.DeviceID,
			12345,
		)
	}

	wantObjectID := bactypes.ObjectID{
		Type:     defs.ObjectType(0),
		Instance: 42,
	}

	if decoded.ObjectID != wantObjectID {
		t.Errorf(
			"ObjectID: got %+v, want %+v",
			decoded.ObjectID,
			wantObjectID,
		)
	}

	if decoded.ObjectName.Encoding != 0 {
		t.Errorf(
			"ObjectName.Encoding: got %d, want 0",
			decoded.ObjectName.Encoding,
		)
	}

	if string(decoded.ObjectName.Value) != "Zone Temperature" {
		t.Errorf(
			"ObjectName: got %q, want %q",
			decoded.ObjectName.Value,
			"Zone Temperature",
		)
	}
}
