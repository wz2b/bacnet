package services

import (
	"bytes"
	"testing"

	"github.com/wz2b/bacnet/apdu"
	"github.com/wz2b/bacnet/bactypes"
	"github.com/wz2b/bacnet/codec"
)

func TestWhoHasRoundTripByObjectID(t *testing.T) {
	original := WhoHas{
		LowLimit:  1000,
		HighLimit: 2000,
		Identifier: bactypes.ObjectID{
			Type:     bactypes.ObjectType(0),
			Instance: 42,
		},
		IsObjectName: false,
	}

	a := original.APDU()

	decoded, err := DecodeWhoHas(&a)
	if err != nil {
		t.Fatalf("DecodeWhoHas() error: %v", err)
	}

	if decoded.LowLimit != original.LowLimit {
		t.Errorf(
			"LowLimit: got %d, want %d",
			decoded.LowLimit,
			original.LowLimit,
		)
	}

	if decoded.HighLimit != original.HighLimit {
		t.Errorf(
			"HighLimit: got %d, want %d",
			decoded.HighLimit,
			original.HighLimit,
		)
	}

	if decoded.IsObjectName {
		t.Errorf("IsObjectName: got true, want false")
	}

	if decoded.Identifier != original.Identifier {
		t.Errorf(
			"Identifier: got %+v, want %+v",
			decoded.Identifier,
			original.Identifier,
		)
	}
}

func TestWhoHasRoundTripByObjectName(t *testing.T) {
	original := WhoHas{
		LowLimit:     -1,
		HighLimit:    -1,
		IsObjectName: true,
		Name: codec.CharacterString{
			Value: []byte("Zone Temperature"),
		},
	}

	a := original.APDU()

	decoded, err := DecodeWhoHas(&a)
	if err != nil {
		t.Fatalf("DecodeWhoHas() error: %v", err)
	}

	if decoded.LowLimit != -1 {
		t.Errorf(
			"LowLimit: got %d, want -1",
			decoded.LowLimit,
		)
	}

	if decoded.HighLimit != -1 {
		t.Errorf(
			"HighLimit: got %d, want -1",
			decoded.HighLimit,
		)
	}

	if !decoded.IsObjectName {
		t.Errorf("IsObjectName: got false, want true")
	}

	if string(decoded.Name.Value) != string(original.Name.Value) {
		t.Errorf(
			"Name: got %q, want %q",
			decoded.Name.Value,
			original.Name.Value,
		)
	}
}

func TestWhoHasLimitRoundTrip(t *testing.T) {
	tests := []struct {
		name      string
		lowLimit  int32
		highLimit int32
	}{
		{
			name:      "no limits",
			lowLimit:  -1,
			highLimit: -1,
		},
		{
			name:      "small range",
			lowLimit:  0,
			highLimit: 100,
		},
		{
			name:      "large range",
			lowLimit:  100000,
			highLimit: 200000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := WhoHas{
				LowLimit:  tt.lowLimit,
				HighLimit: tt.highLimit,
				Identifier: bactypes.ObjectID{
					Type:     bactypes.ObjectType(0),
					Instance: 1,
				},
			}

			a := original.APDU()

			decoded, err := DecodeWhoHas(&a)
			if err != nil {
				t.Fatalf("DecodeWhoHas() error: %v", err)
			}

			if decoded.LowLimit != original.LowLimit {
				t.Errorf(
					"LowLimit: got %d, want %d",
					decoded.LowLimit,
					original.LowLimit,
				)
			}

			if decoded.HighLimit != original.HighLimit {
				t.Errorf(
					"HighLimit: got %d, want %d",
					decoded.HighLimit,
					original.HighLimit,
				)
			}
		})
	}
}

func TestWhoHasEncodingByObjectID(t *testing.T) {
	msg := WhoHas{
		LowLimit:  1000,
		HighLimit: 2000,
		Identifier: bactypes.ObjectID{
			Type:     bactypes.ObjectType(0), // Analog Input
			Instance: 42,
		},
		IsObjectName: false,
	}

	a := msg.APDU()

	got, err := apdu.Encode(a)
	if err != nil {
		t.Fatalf("apdu.Encode() error: %v", err)
	}

	want := []byte{
		// Unconfirmed-Request, Who-Has
		0x10, 0x07,

		// [0] lowLimit = 1000
		// context tag 0, length 2
		0x0A, 0x03, 0xE8,

		// [1] highLimit = 2000
		// context tag 1, length 2
		0x1A, 0x07, 0xD0,

		// [2] objectIdentifier = AnalogInput:42
		// context tag 2, length 4
		0x2C,
		0x00, 0x00, 0x00, 0x2A,
	}

	if !bytes.Equal(got, want) {
		t.Errorf(
			"encoded Who-Has mismatch:\n got: % X\nwant: % X",
			got,
			want,
		)
	}
}

func TestWhoHasDecodingByObjectID(t *testing.T) {
	data := []byte{
		// Unconfirmed-Request, Who-Has
		0x10, 0x07,

		// [0] lowLimit = 1000
		0x0A, 0x03, 0xE8,

		// [1] highLimit = 2000
		0x1A, 0x07, 0xD0,

		// [2] objectIdentifier = AnalogInput:42
		0x2C,
		0x00, 0x00, 0x00, 0x2A,
	}

	a, err := apdu.Decode(data)
	if err != nil {
		t.Fatalf("apdu.Decode() error: %v", err)
	}

	decoded, err := DecodeWhoHas(&a)
	if err != nil {
		t.Fatalf("DecodeWhoHas() error: %v", err)
	}

	if decoded.LowLimit != 1000 {
		t.Errorf(
			"LowLimit: got %d, want %d",
			decoded.LowLimit,
			1000,
		)
	}

	if decoded.HighLimit != 2000 {
		t.Errorf(
			"HighLimit: got %d, want %d",
			decoded.HighLimit,
			2000,
		)
	}

	if decoded.IsObjectName {
		t.Errorf("IsObjectName: got true, want false")
	}

	wantObjectID := bactypes.ObjectID{
		Type:     bactypes.ObjectType(0),
		Instance: 42,
	}

	if decoded.Identifier != wantObjectID {
		t.Errorf(
			"Identifier: got %+v, want %+v",
			decoded.Identifier,
			wantObjectID,
		)
	}
}

func TestWhoHasEncodingByObjectName(t *testing.T) {
	msg := WhoHas{
		LowLimit:     -1,
		HighLimit:    -1,
		IsObjectName: true,
		Name: codec.CharacterString{
			Value: []byte("Zone Temperature"),
		},
	}

	a := msg.APDU()

	got, err := apdu.Encode(a)
	if err != nil {
		t.Fatalf("apdu.Encode() error: %v", err)
	}

	want := []byte{
		// Unconfirmed-Request, Who-Has
		0x10, 0x07,

		// [3] objectName
		// context tag 3, extended length 17
		0x3D, 0x11,

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
			"encoded Who-Has mismatch:\n got: % X\nwant: % X",
			got,
			want,
		)
	}
}

func TestWhoHasDecodingByObjectName(t *testing.T) {
	data := []byte{
		// Unconfirmed-Request, Who-Has
		0x10, 0x07,

		// [3] objectName
		// context tag 3, extended length 17
		0x3D, 0x11,

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

	decoded, err := DecodeWhoHas(&a)
	if err != nil {
		t.Fatalf("DecodeWhoHas() error: %v", err)
	}

	if decoded.LowLimit != -1 {
		t.Errorf(
			"LowLimit: got %d, want -1",
			decoded.LowLimit,
		)
	}

	if decoded.HighLimit != -1 {
		t.Errorf(
			"HighLimit: got %d, want -1",
			decoded.HighLimit,
		)
	}

	if !decoded.IsObjectName {
		t.Errorf("IsObjectName: got false, want true")
	}

	if decoded.Name.Encoding != 0 {
		t.Errorf(
			"Name.Encoding: got %d, want 0",
			decoded.Name.Encoding,
		)
	}

	if string(decoded.Name.Value) != "Zone Temperature" {
		t.Errorf(
			"Name: got %q, want %q",
			decoded.Name.Value,
			"Zone Temperature",
		)
	}
}
