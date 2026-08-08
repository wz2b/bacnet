package services

import (
	"bytes"
	"testing"

	"github.com/wz2b/bacnet/apdu"
	"github.com/wz2b/bacnet/bactypes"
)

func TestReadPropertyACKEncoding(t *testing.T) {
	msg := ReadPropertyACK{
		Object: bactypes.ObjectID{
			Type:     bactypes.ObjectType(2), // Analog Value
			Instance: 1234,
		},
		PropertyIdentifier: 85, // Present Value

		// Application-tagged REAL 72.5:
		//
		// 44             Application Real, length 4
		// 42 91 00 00    IEEE-754 72.5
		EncodedValue: []byte{
			0x44,
			0x42, 0x91, 0x00, 0x00,
		},
	}

	a := msg.APDU()
	a.InvokeID = 42

	got, err := apdu.Encode(a)
	if err != nil {
		t.Fatalf("apdu.Encode() error: %v", err)
	}

	want := []byte{
		// ComplexACK
		0x30,

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

		// [3] propertyValue opening tag
		0x3E,

		// Application Real = 72.5
		0x44,
		0x42, 0x91, 0x00, 0x00,

		// [3] propertyValue closing tag
		0x3F,
	}

	if !bytes.Equal(got, want) {
		t.Errorf(
			"encoded ReadPropertyACK mismatch:\n got: % X\nwant: % X",
			got,
			want,
		)
	}
}

func TestReadPropertyACKDecoding(t *testing.T) {
	data := []byte{
		// ComplexACK
		0x30,

		// Invoke ID = 42
		0x2A,

		// ReadProperty service choice = 12
		0x0C,

		// [0] objectIdentifier = AnalogValue:1234
		0x0C,
		0x00, 0x80, 0x04, 0xD2,

		// [1] propertyIdentifier = Present Value (85)
		0x19, 0x55,

		// [3] propertyValue opening tag
		0x3E,

		// Application Real = 72.5
		0x44,
		0x42, 0x91, 0x00, 0x00,

		// [3] propertyValue closing tag
		0x3F,
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

	decoded, err := DecodeReadPropertyACK(&a)
	if err != nil {
		t.Fatalf("DecodeReadPropertyACK() error: %v", err)
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

	wantValue := []byte{
		0x44,
		0x42, 0x91, 0x00, 0x00,
	}

	if !bytes.Equal(decoded.EncodedValue, wantValue) {
		t.Errorf(
			"EncodedValue mismatch:\n got: % X\nwant: % X",
			decoded.EncodedValue,
			wantValue,
		)
	}
}
