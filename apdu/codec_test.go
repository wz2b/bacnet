package apdu

import (
	"bytes"
	"testing"

	"github.com/wz2b/bacnet/defs"
)

func TestEncodeUnconfirmedServiceRequest(t *testing.T) {
	a := APDU{
		Type:                     defs.PDUTypeUnconfirmedServiceRequest,
		UnconfirmedServiceChoice: defs.ServiceUnconfirmedWhoIs,
		Data: []byte{
			0x0A, 0x03, 0xE8,
		},
	}

	got, err := Encode(a)
	if err != nil {
		t.Fatalf("Encode() error: %v", err)
	}

	want := []byte{
		0x10,
		0x08,
		0x0A, 0x03, 0xE8,
	}

	if !bytes.Equal(got, want) {
		t.Errorf(
			"encoded APDU mismatch:\n got: % X\nwant: % X",
			got,
			want,
		)
	}
}

func TestDecodeUnconfirmedServiceRequest(t *testing.T) {
	data := []byte{
		0x10,
		0x08,
		0x0A, 0x03, 0xE8,
	}

	a, err := Decode(data)
	if err != nil {
		t.Fatalf("Decode() error: %v", err)
	}

	if a.Type != defs.PDUTypeUnconfirmedServiceRequest {
		t.Errorf(
			"Type: got 0x%02X, want 0x%02X",
			byte(a.Type),
			byte(defs.PDUTypeUnconfirmedServiceRequest),
		)
	}

	if a.UnconfirmedServiceChoice != defs.ServiceUnconfirmedWhoIs {
		t.Errorf(
			"UnconfirmedServiceChoice: got %d, want %d",
			a.UnconfirmedServiceChoice,
			defs.ServiceUnconfirmedWhoIs,
		)
	}

	wantData := []byte{0x0A, 0x03, 0xE8}

	if !bytes.Equal(a.Data, wantData) {
		t.Errorf(
			"Data mismatch:\n got: % X\nwant: % X",
			a.Data,
			wantData,
		)
	}
}

func TestEncodeConfirmedServiceRequest(t *testing.T) {
	a := APDU{
		Type:                   defs.PDUTypeConfirmedServiceRequest,
		ConfirmedServiceChoice: defs.ServiceConfirmedReadProperty,
		InvokeID:               42,
		MaxSegments:            0,
		MaxAPDU:                1476,
		Data: []byte{
			0x0C,
			0x00, 0x80, 0x04, 0xD2,
			0x19, 0x55,
		},
	}

	got, err := Encode(a)
	if err != nil {
		t.Fatalf("Encode() error: %v", err)
	}

	want := []byte{
		0x00,
		0x05,
		0x2A,
		0x0C,

		0x0C,
		0x00, 0x80, 0x04, 0xD2,
		0x19, 0x55,
	}

	if !bytes.Equal(got, want) {
		t.Errorf(
			"encoded APDU mismatch:\n got: % X\nwant: % X",
			got,
			want,
		)
	}
}

func TestDecodeConfirmedServiceRequest(t *testing.T) {
	data := []byte{
		0x00,
		0x05,
		0x2A,
		0x0C,

		0x0C,
		0x00, 0x80, 0x04, 0xD2,
		0x19, 0x55,
	}

	a, err := Decode(data)
	if err != nil {
		t.Fatalf("Decode() error: %v", err)
	}

	if a.Type != defs.PDUTypeConfirmedServiceRequest {
		t.Errorf(
			"Type: got 0x%02X, want 0x%02X",
			byte(a.Type),
			byte(defs.PDUTypeConfirmedServiceRequest),
		)
	}

	if a.InvokeID != 42 {
		t.Errorf(
			"InvokeID: got %d, want %d",
			a.InvokeID,
			42,
		)
	}

	if a.ConfirmedServiceChoice != defs.ServiceConfirmedReadProperty {
		t.Errorf(
			"ConfirmedServiceChoice: got %d, want %d",
			a.ConfirmedServiceChoice,
			defs.ServiceConfirmedReadProperty,
		)
	}

	if a.MaxSegments != 0 {
		t.Errorf(
			"MaxSegments: got %d, want %d",
			a.MaxSegments,
			0,
		)
	}

	if a.MaxAPDU != 1476 {
		t.Errorf(
			"MaxAPDU: got %d, want %d",
			a.MaxAPDU,
			1476,
		)
	}

	wantData := []byte{
		0x0C,
		0x00, 0x80, 0x04, 0xD2,
		0x19, 0x55,
	}

	if !bytes.Equal(a.Data, wantData) {
		t.Errorf(
			"Data mismatch:\n got: % X\nwant: % X",
			a.Data,
			wantData,
		)
	}
}

func TestEncodeComplexACK(t *testing.T) {
	a := APDU{
		Type:                   defs.PDUTypeComplexACK,
		InvokeID:               42,
		ConfirmedServiceChoice: defs.ServiceConfirmedReadProperty,
		Data: []byte{
			0x0C,
			0x00, 0x80, 0x04, 0xD2,
			0x19, 0x55,
			0x3E,
			0x44, 0x42, 0x91, 0x00, 0x00,
			0x3F,
		},
	}

	got, err := Encode(a)
	if err != nil {
		t.Fatalf("Encode() error: %v", err)
	}

	want := []byte{
		0x30,
		0x2A,
		0x0C,

		0x0C,
		0x00, 0x80, 0x04, 0xD2,
		0x19, 0x55,
		0x3E,
		0x44, 0x42, 0x91, 0x00, 0x00,
		0x3F,
	}

	if !bytes.Equal(got, want) {
		t.Errorf(
			"encoded APDU mismatch:\n got: % X\nwant: % X",
			got,
			want,
		)
	}
}

func TestDecodeComplexACK(t *testing.T) {
	data := []byte{
		0x30,
		0x2A,
		0x0C,

		0x0C,
		0x00, 0x80, 0x04, 0xD2,
		0x19, 0x55,
		0x3E,
		0x44, 0x42, 0x91, 0x00, 0x00,
		0x3F,
	}

	a, err := Decode(data)
	if err != nil {
		t.Fatalf("Decode() error: %v", err)
	}

	if a.Type != defs.PDUTypeComplexACK {
		t.Errorf(
			"Type: got 0x%02X, want 0x%02X",
			byte(a.Type),
			byte(defs.PDUTypeComplexACK),
		)
	}

	if a.InvokeID != 42 {
		t.Errorf(
			"InvokeID: got %d, want %d",
			a.InvokeID,
			42,
		)
	}

	if a.ConfirmedServiceChoice != defs.ServiceConfirmedReadProperty {
		t.Errorf(
			"ConfirmedServiceChoice: got %d, want %d",
			a.ConfirmedServiceChoice,
			defs.ServiceConfirmedReadProperty,
		)
	}

	wantData := []byte{
		0x0C,
		0x00, 0x80, 0x04, 0xD2,
		0x19, 0x55,
		0x3E,
		0x44, 0x42, 0x91, 0x00, 0x00,
		0x3F,
	}

	if !bytes.Equal(a.Data, wantData) {
		t.Errorf(
			"Data mismatch:\n got: % X\nwant: % X",
			a.Data,
			wantData,
		)
	}
}
