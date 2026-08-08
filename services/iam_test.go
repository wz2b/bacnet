package services

import (
	"bytes"
	"testing"

	"github.com/wz2b/bacnet/apdu"
	"github.com/wz2b/bacnet/bactypes"
)

func TestIAmRoundTrip(t *testing.T) {
	original := IAm{
		DeviceID:     12345,
		MaxAPDU:      1476,
		Segmentation: 3,
		VendorID:     15,
	}

	encoded := original.APDU()

	decoded, err := DecodeIAm(&encoded)
	if err != nil {
		t.Fatalf("DecodeIAm() error: %v", err)
	}

	if decoded.DeviceID != original.DeviceID {
		t.Errorf("DeviceID: got %d, want %d", decoded.DeviceID, original.DeviceID)
	}

	if decoded.MaxAPDU != original.MaxAPDU {
		t.Errorf("MaxAPDU: got %d, want %d", decoded.MaxAPDU, original.MaxAPDU)
	}

	if decoded.Segmentation != original.Segmentation {
		t.Errorf("Segmentation: got %d, want %d", decoded.Segmentation, original.Segmentation)
	}

	if decoded.VendorID != original.VendorID {
		t.Errorf("VendorID: got %d, want %d", decoded.VendorID, original.VendorID)
	}

	_ = bactypes.ObjectType(0) // remove if import isn't needed
}

func TestIAmEncoding(t *testing.T) {
	msg := IAm{
		DeviceID:     12345,
		MaxAPDU:      1476,
		Segmentation: 3,
		VendorID:     15,
	}

	a := msg.APDU()

	got, err := apdu.Encode(a)
	if err != nil {
		t.Fatalf("apdu.Encode() error: %v", err)
	}

	want := []byte{
		0x10, 0x00,
		0xC4, 0x02, 0x00, 0x30, 0x39,
		0x22, 0x05, 0xC4,
		0x91, 0x03,
		0x21, 0x0F,
	}

	if !bytes.Equal(got, want) {
		t.Errorf(
			"encoded I-Am mismatch:\n got: % X\nwant: % X",
			got,
			want,
		)
	}
}

func TestIAmDecoding(t *testing.T) {
	data := []byte{
		0x10, 0x00,
		0xC4, 0x02, 0x00, 0x30, 0x39,
		0x22, 0x05, 0xC4,
		0x91, 0x03,
		0x21, 0x0F,
	}

	a, err := apdu.Decode(data)
	if err != nil {
		t.Fatalf("apdu.Decode() error: %v", err)
	}

	decoded, err := DecodeIAm(&a)
	if err != nil {
		t.Fatalf("DecodeIAm() error: %v", err)
	}

	if decoded.DeviceID != 12345 {
		t.Errorf("DeviceID: got %d, want %d", decoded.DeviceID, 12345)
	}

	if decoded.MaxAPDU != 1476 {
		t.Errorf("MaxAPDU: got %d, want %d", decoded.MaxAPDU, 1476)
	}

	if decoded.Segmentation != 3 {
		t.Errorf("Segmentation: got %d, want %d", decoded.Segmentation, 3)
	}

	if decoded.VendorID != 15 {
		t.Errorf("VendorID: got %d, want %d", decoded.VendorID, 15)
	}
}
