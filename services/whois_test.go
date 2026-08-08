package services

import (
	"bytes"
	"testing"

	"github.com/wz2b/bacnet/apdu"
	"github.com/wz2b/bacnet/defs"
)

func checkWhoIsEncoding(t *testing.T, lowLimit, highLimit int32) {
	t.Helper()

	original := WhoIs{
		LowLimit:  lowLimit,
		HighLimit: highLimit,
	}

	data := original.APDU()

	decoded, err := DecodeWhoIs(&apdu.APDU{
		Type:          defs.PDUTypeUnconfirmedServiceRequest,
		ServiceChoice: defs.ServiceUnconfirmedWhoIs,
		Data:          data[2:],
	})
	if err != nil {
		t.Fatalf("DecodeWhoIs() error: %v", err)
	}

	if decoded.LowLimit != original.LowLimit {
		t.Errorf(
			"LowLimit mismatch: got %d, want %d",
			decoded.LowLimit,
			original.LowLimit,
		)
	}

	if decoded.HighLimit != original.HighLimit {
		t.Errorf(
			"HighLimit mismatch: got %d, want %d",
			decoded.HighLimit,
			original.HighLimit,
		)
	}
}

func TestWhoIsEncodingAndDecoding(t *testing.T) {
	checkWhoIsEncoding(t, -1, -1)

	step := int32(defs.MaxInstance / 4)

	for lowLimit := int32(0); lowLimit <= defs.MaxInstance; lowLimit += step {
		for highLimit := int32(0); highLimit <= defs.MaxInstance; highLimit += step {
			checkWhoIsEncoding(t, lowLimit, highLimit)
		}
	}
}

func TestWhoIsEncodingWithoutLimits(t *testing.T) {
	msg := WhoIs{
		LowLimit:  -1,
		HighLimit: -1,
	}

	got := msg.APDU()

	want := []byte{
		// Unconfirmed-Request, Who-Is
		0x10, 0x08,
	}

	if !bytes.Equal(got, want) {
		t.Errorf(
			"encoded Who-Is mismatch:\n got: % X\nwant: % X",
			got,
			want,
		)
	}
}

func TestWhoIsDecodingWithoutLimits(t *testing.T) {
	data := []byte{
		// Unconfirmed-Request, Who-Is
		0x10, 0x08,
	}

	a, err := apdu.Decode(data)
	if err != nil {
		t.Fatalf("apdu.Decode() error: %v", err)
	}

	decoded, err := DecodeWhoIs(&a)
	if err != nil {
		t.Fatalf("DecodeWhoIs() error: %v", err)
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
}

func TestWhoIsEncodingWithLimits(t *testing.T) {
	msg := WhoIs{
		LowLimit:  1000,
		HighLimit: 2000,
	}

	got := msg.APDU()

	want := []byte{
		// Unconfirmed-Request, Who-Is
		0x10, 0x08,

		// [0] deviceInstanceRangeLowLimit = 1000
		// context tag 0, length 2
		0x0A, 0x03, 0xE8,

		// [1] deviceInstanceRangeHighLimit = 2000
		// context tag 1, length 2
		0x1A, 0x07, 0xD0,
	}

	if !bytes.Equal(got, want) {
		t.Errorf(
			"encoded Who-Is mismatch:\n got: % X\nwant: % X",
			got,
			want,
		)
	}
}

func TestWhoIsDecodingWithLimits(t *testing.T) {
	data := []byte{
		// Unconfirmed-Request, Who-Is
		0x10, 0x08,

		// [0] deviceInstanceRangeLowLimit = 1000
		0x0A, 0x03, 0xE8,

		// [1] deviceInstanceRangeHighLimit = 2000
		0x1A, 0x07, 0xD0,
	}

	a, err := apdu.Decode(data)
	if err != nil {
		t.Fatalf("apdu.Decode() error: %v", err)
	}

	decoded, err := DecodeWhoIs(&a)
	if err != nil {
		t.Fatalf("DecodeWhoIs() error: %v", err)
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
}
