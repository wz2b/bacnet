package services

import (
	"bytes"
	"testing"

	"github.com/wz2b/bacnet/apdu"
	"github.com/wz2b/bacnet/defs"
)

func checkWhoIsEncoding(
	t *testing.T,
	lowLimit *uint32,
	highLimit *uint32,
) {
	t.Helper()

	original := WhoIs{
		LowLimit:  lowLimit,
		HighLimit: highLimit,
	}

	data := original.APDU()

	decoded, err := DecodeWhoIs(&apdu.APDU{
		Type:                     defs.PDUTypeUnconfirmedServiceRequest,
		UnconfirmedServiceChoice: defs.ServiceUnconfirmedWhoIs,
		Data:                     data[2:],
	})
	if err != nil {
		t.Fatalf("DecodeWhoIs() error: %v", err)
	}

	if lowLimit == nil {
		if decoded.LowLimit != nil {
			t.Errorf(
				"LowLimit mismatch: got %d, want nil",
				*decoded.LowLimit,
			)
		}
	} else if decoded.LowLimit == nil ||
		*decoded.LowLimit != *lowLimit {
		t.Errorf(
			"LowLimit mismatch: got %v, want %v",
			decoded.LowLimit,
			lowLimit,
		)
	}

	if highLimit == nil {
		if decoded.HighLimit != nil {
			t.Errorf(
				"HighLimit mismatch: got %d, want nil",
				*decoded.HighLimit,
			)
		}
	} else if decoded.HighLimit == nil ||
		*decoded.HighLimit != *highLimit {
		t.Errorf(
			"HighLimit mismatch: got %v, want %v",
			decoded.HighLimit,
			highLimit,
		)
	}
}

func TestWhoIsEncodingAndDecoding(t *testing.T) {
	checkWhoIsEncoding(t, nil, nil)

	step := uint32(defs.MaxInstance / 4)

	for lowLimit := uint32(0); lowLimit <= uint32(defs.MaxInstance); lowLimit += step {
		for highLimit := uint32(0); highLimit <= uint32(defs.MaxInstance); highLimit += step {
			checkWhoIsEncoding(
				t,
				uint32Ptr(lowLimit),
				uint32Ptr(highLimit),
			)
		}
	}
}

func TestWhoIsEncodingWithoutLimits(t *testing.T) {
	msg := WhoIs{
		LowLimit:  nil,
		HighLimit: nil,
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
}

func TestWhoIsEncodingWithLimits(t *testing.T) {
	msg := WhoIs{
		LowLimit:  uint32Ptr(1000),
		HighLimit: uint32Ptr(2000),
	}

	got := msg.APDU()

	want := []byte{
		// Unconfirmed-Request, Who-Is
		0x10, 0x08,

		// [0] deviceInstanceRangeLowLimit = 1000
		0x0A, 0x03, 0xE8,

		// [1] deviceInstanceRangeHighLimit = 2000
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

	if decoded.LowLimit == nil ||
		*decoded.LowLimit != 1000 {
		t.Errorf(
			"LowLimit: got %v, want 1000",
			decoded.LowLimit,
		)
	}

	if decoded.HighLimit == nil ||
		*decoded.HighLimit != 2000 {
		t.Errorf(
			"HighLimit: got %v, want 2000",
			decoded.HighLimit,
		)
	}
}
