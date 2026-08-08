package npdu

import (
	"bytes"
	"testing"
)

func TestNPDUEncodingAndDecoding(t *testing.T) {
	original := &NPDU{
		ProtocolVersion: 0x01,
		Priority:        2,
		ExpectingReply:  true,
		NetworkLayer:    false,
		PDU: []byte{
			0x10, 0x08,
			0x0A, 0x03, 0xE8,
		},
	}

	wire, err := original.Encode()
	if err != nil {
		t.Fatalf("Encode() error: %v", err)
	}

	decoded, err := Decode(wire)
	if err != nil {
		t.Fatalf("Decode() error: %v", err)
	}

	if decoded.ProtocolVersion != original.ProtocolVersion {
		t.Errorf(
			"ProtocolVersion: got %d, want %d",
			decoded.ProtocolVersion,
			original.ProtocolVersion,
		)
	}

	if decoded.Priority != original.Priority {
		t.Errorf(
			"Priority: got %d, want %d",
			decoded.Priority,
			original.Priority,
		)
	}

	if decoded.ExpectingReply != original.ExpectingReply {
		t.Errorf(
			"ExpectingReply: got %t, want %t",
			decoded.ExpectingReply,
			original.ExpectingReply,
		)
	}

	if decoded.NetworkLayer != original.NetworkLayer {
		t.Errorf(
			"NetworkLayer: got %t, want %t",
			decoded.NetworkLayer,
			original.NetworkLayer,
		)
	}

	if !bytes.Equal(decoded.Payload(), original.PDU) {
		t.Errorf(
			"Payload mismatch:\n got: % X\nwant: % X",
			decoded.Payload(),
			original.PDU,
		)
	}
}
