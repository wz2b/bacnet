package services

import (
	"bytes"
	"testing"

	"github.com/wz2b/bacnet/apdu"
	"github.com/wz2b/bacnet/bactypes"
	"github.com/wz2b/bacnet/codec"
	"github.com/wz2b/bacnet/defs"
)

func TestWhoHasRoundTripByObjectID(t *testing.T) {
	original := WhoHas{
		LowLimit:  uint32Ptr(1000),
		HighLimit: uint32Ptr(2000),
		Identifier: bactypes.ObjectID{
			Type:     defs.ObjectType(0),
			Instance: 42,
		},
		IsObjectName: false,
	}

	a := original.APDU()

	decoded, err := DecodeWhoHas(&a)
	if err != nil {
		t.Fatalf("DecodeWhoHas() error: %v", err)
	}

	if decoded.LowLimit == nil ||
		original.LowLimit == nil ||
		*decoded.LowLimit != *original.LowLimit {
		t.Errorf(
			"LowLimit: got %v, want %v",
			decoded.LowLimit,
			original.LowLimit,
		)
	}

	if decoded.HighLimit == nil ||
		original.HighLimit == nil ||
		*decoded.HighLimit != *original.HighLimit {
		t.Errorf(
			"HighLimit: got %v, want %v",
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
		LowLimit:     nil,
		HighLimit:    nil,
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

	if decoded.LowLimit != nil {
		t.Errorf(
			"LowLimit: got %d, want nil",
			*decoded.LowLimit,
		)
	}

	if decoded.HighLimit != nil {
		t.Errorf(
			"HighLimit: got %d, want nil",
			*decoded.HighLimit,
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
		lowLimit  *uint32
		highLimit *uint32
	}{
		{
			name: "no limits",
		},
		{
			name:      "small range",
			lowLimit:  uint32Ptr(0),
			highLimit: uint32Ptr(100),
		},
		{
			name:      "large range",
			lowLimit:  uint32Ptr(100000),
			highLimit: uint32Ptr(200000),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := WhoHas{
				LowLimit:  tt.lowLimit,
				HighLimit: tt.highLimit,
				Identifier: bactypes.ObjectID{
					Type:     defs.ObjectType(0),
					Instance: 1,
				},
			}

			a := original.APDU()

			decoded, err := DecodeWhoHas(&a)
			if err != nil {
				t.Fatalf("DecodeWhoHas() error: %v", err)
			}

			if tt.lowLimit == nil {
				if decoded.LowLimit != nil {
					t.Errorf("LowLimit: got %d, want nil", *decoded.LowLimit)
				}
			} else if decoded.LowLimit == nil ||
				*decoded.LowLimit != *tt.lowLimit {
				t.Errorf(
					"LowLimit: got %v, want %v",
					decoded.LowLimit,
					tt.lowLimit,
				)
			}

			if tt.highLimit == nil {
				if decoded.HighLimit != nil {
					t.Errorf("HighLimit: got %d, want nil", *decoded.HighLimit)
				}
			} else if decoded.HighLimit == nil ||
				*decoded.HighLimit != *tt.highLimit {
				t.Errorf(
					"HighLimit: got %v, want %v",
					decoded.HighLimit,
					tt.highLimit,
				)
			}
		})
	}
}

func TestWhoHasEncodingByObjectID(t *testing.T) {
	msg := WhoHas{
		LowLimit:  uint32Ptr(1000),
		HighLimit: uint32Ptr(2000),
		Identifier: bactypes.ObjectID{
			Type:     defs.ObjectType(0),
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
		0x10, 0x07,
		0x0A, 0x03, 0xE8,
		0x1A, 0x07, 0xD0,
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
		0x10, 0x07,
		0x0A, 0x03, 0xE8,
		0x1A, 0x07, 0xD0,
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

	if decoded.LowLimit == nil || *decoded.LowLimit != 1000 {
		t.Errorf("LowLimit: got %v, want 1000", decoded.LowLimit)
	}

	if decoded.HighLimit == nil || *decoded.HighLimit != 2000 {
		t.Errorf("HighLimit: got %v, want 2000", decoded.HighLimit)
	}

	wantObjectID := bactypes.ObjectID{
		Type:     defs.ObjectType(0),
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
		LowLimit:     nil,
		HighLimit:    nil,
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
		0x10, 0x07,
		0x3D, 0x11,
		0x00,
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
		0x10, 0x07,
		0x3D, 0x11,
		0x00,
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

	if decoded.LowLimit != nil {
		t.Errorf("LowLimit: got %d, want nil", *decoded.LowLimit)
	}

	if decoded.HighLimit != nil {
		t.Errorf("HighLimit: got %d, want nil", *decoded.HighLimit)
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
