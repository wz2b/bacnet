package bvlc

import (
	"bytes"
	"testing"

	"github.com/wz2b/bacnet/defs"
)

func TestReadBroadcastDistributionTableEncoding(t *testing.T) {
	request := &ReadBroadcastDistributionTable{}

	got, err := request.Encode()
	if err != nil {
		t.Fatalf("Encode() error: %v", err)
	}

	want := []byte{
		0x81, // BACnet/IP
		0x02, // Read-Broadcast-Distribution-Table
		0x00, 0x04,
	}

	if !bytes.Equal(got, want) {
		t.Errorf(
			"encoded BVLC mismatch:\n got: % X\nwant: % X",
			got,
			want,
		)
	}
}

func TestReadBroadcastDistributionTableAckDecoding(t *testing.T) {
	b := &BVLC{
		BVLLType: BVLCTypeBACnetIP,
		Function: defs.BVLCFunctionReadBroadcastDistributionTableACK,
		Payload: []byte{
			// 10.7.6.148:47808, mask 255.255.255.255
			0x0A, 0x07, 0x06, 0x94,
			0xBA, 0xC0,
			0xFF, 0xFF, 0xFF, 0xFF,

			// 10.7.5.101:47808, mask 255.255.255.0
			0x0A, 0x07, 0x05, 0x65,
			0xBA, 0xC0,
			0xFF, 0xFF, 0xFF, 0x00,
		},
	}

	ack, err := DecodeReadBroadcastDistributionTableAck(b)
	if err != nil {
		t.Fatalf(
			"DecodeReadBroadcastDistributionTableAck() error: %v",
			err,
		)
	}

	if len(ack.Entries) != 2 {
		t.Fatalf(
			"entry count: got %d, want 2",
			len(ack.Entries),
		)
	}

	if got := ack.Entries[0].Address.String(); got != "10.7.6.148:47808" {
		t.Errorf(
			"entry 0 address: got %q, want %q",
			got,
			"10.7.6.148:47808",
		)
	}

	if ack.Entries[0].Mask != 0xFFFFFFFF {
		t.Errorf(
			"entry 0 mask: got 0x%08X, want 0xFFFFFFFF",
			ack.Entries[0].Mask,
		)
	}

	if got := ack.Entries[1].Address.String(); got != "10.7.5.101:47808" {
		t.Errorf(
			"entry 1 address: got %q, want %q",
			got,
			"10.7.5.101:47808",
		)
	}

	if ack.Entries[1].Mask != 0xFFFFFF00 {
		t.Errorf(
			"entry 1 mask: got 0x%08X, want 0xFFFFFF00",
			ack.Entries[1].Mask,
		)
	}
}
