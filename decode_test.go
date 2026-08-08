package bacnet

import (
	"bytes"
	"testing"

	"github.com/wz2b/bacnet/npdu"
	"github.com/wz2b/bacnet/services"
)

func TestDecodeNetworkLayerMessage(t *testing.T) {
	// BACnet/IP Original-Broadcast-NPDU containing:
	//
	//   NPDU version:        1
	//   network-layer bit:   set
	//   message type:        Who-Is-Router-To-Network (0x00)
	//   network number:      7114 (0x1BCA)
	wire := []byte{
		0x81,       // BACnet/IP
		0x0B,       // Original-Broadcast-NPDU
		0x00, 0x09, // BVLC length

		0x01, // NPDU version
		0x80, // network-layer message
		0x00, // Who-Is-Router-To-Network

		0x1B, 0xCA, // network 7114
	}

	msg, err := Decode(wire)
	if err != nil {
		t.Fatalf("Decode() error: %v", err)
	}

	if msg.BVLC == nil {
		t.Fatal("BVLC is nil")
	}

	if msg.NPDU == nil {
		t.Fatal("NPDU is nil")
	}

	if !msg.NPDU.NetworkLayer {
		t.Fatal("NPDU is not marked as network-layer")
	}

	if msg.NPDU.MessageType != 0x00 {
		t.Errorf(
			"MessageType: got 0x%02X, want 0x00",
			msg.NPDU.MessageType,
		)
	}

	if msg.APDU != nil {
		t.Fatal("APDU should be nil for network-layer message")
	}

	value, ok := msg.Value.(*npdu.UnknownNetworkMessage)
	if !ok {
		t.Fatalf(
			"Value type: got %T, want *npdu.UnknownNetworkMessage",
			msg.Value,
		)
	}

	wantData := []byte{0x1B, 0xCA}

	if !bytes.Equal(value.Data, wantData) {
		t.Errorf(
			"network message data mismatch:\n got: % X\nwant: % X",
			value.Data,
			wantData,
		)
	}
}

func TestDecodeWhoIs(t *testing.T) {
	// BACnet/IP Original-Broadcast-NPDU containing an
	// unconfirmed Who-Is with no device-instance limits.
	wire := []byte{
		0x81,       // BACnet/IP
		0x0B,       // Original-Broadcast-NPDU
		0x00, 0x08, // BVLC length

		0x01, // NPDU version
		0x00, // application-layer NPDU

		0x10, // Unconfirmed-Request APDU
		0x08, // Who-Is
	}

	msg, err := Decode(wire)
	if err != nil {
		t.Fatalf("Decode() error: %v", err)
	}

	if msg.BVLC == nil {
		t.Fatal("BVLC is nil")
	}

	if msg.NPDU == nil {
		t.Fatal("NPDU is nil")
	}

	if msg.NPDU.NetworkLayer {
		t.Fatal("NPDU unexpectedly marked as network-layer")
	}

	if msg.APDU == nil {
		t.Fatal("APDU is nil")
	}

	whoIs, ok := msg.Value.(*services.WhoIs)
	if !ok {
		t.Fatalf(
			"Value type: got %T, want *services.WhoIs",
			msg.Value,
		)
	}

	if whoIs == nil {
		t.Fatal("decoded WhoIs is nil")
	}
}
