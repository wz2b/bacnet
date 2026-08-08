package bvlc

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/wz2b/bacnet/bactypes"
	"github.com/wz2b/bacnet/defs"
)

const BVLCTypeBACnetIP byte = 0x81

type BVLC struct {
	BVLLType byte
	Function defs.BVLCFunction
	Length   uint16
	Payload  []byte
}

type BVLCResult struct {
	Code uint16
}

type ForwardedNPDU struct {
	OriginatingAddress IPAddress
	NPDU               []byte
}

type RegisterForeignDevice struct {
	TTL uint16
}

type ReadForeignDeviceTable struct{}

type ForeignDeviceTableEntry struct {
	Address       IPAddress
	TTL           uint16
	RemainingTime uint16
}

type ReadForeignDeviceTableAck struct {
	Entries []ForeignDeviceTableEntry
}

type DeleteForeignDeviceTableEntry struct {
	Address bactypes.Address
}

type DistributeBroadcastToNetwork struct {
	NPDU []byte
}

type OriginalUnicastNPDU struct {
	NPDU []byte
}

type OriginalBroadcastNPDU struct {
	NPDU []byte
}

type BroadcastDistributionTableEntry struct {
	Address IPAddress
	Mask    uint32
}

func EncodeBVLC(
	function defs.BVLCFunction,
	payload []byte,
) ([]byte, error) {
	length := 4 + len(payload)
	if length > 0xffff {
		return nil, fmt.Errorf(
			"BVLC packet too large: %d bytes",
			length,
		)
	}

	packet := make([]byte, length)

	packet[0] = BVLCTypeBACnetIP
	packet[1] = byte(function)
	binary.BigEndian.PutUint16(
		packet[2:4],
		uint16(length),
	)
	copy(packet[4:], payload)

	return packet, nil
}

func Decode(packet []byte) (*BVLC, error) {
	if len(packet) < 4 {
		return nil, fmt.Errorf(
			"BVLC packet too short: %d bytes",
			len(packet),
		)
	}

	if packet[0] != BVLCTypeBACnetIP {
		return nil, fmt.Errorf(
			"invalid BVLC type: 0x%02X",
			packet[0],
		)
	}

	length := binary.BigEndian.Uint16(packet[2:4])

	if int(length) != len(packet) {
		return nil, fmt.Errorf(
			"BVLC length mismatch: header=%d packet=%d",
			length,
			len(packet),
		)
	}

	return &BVLC{
		BVLLType: packet[0],
		Function: defs.BVLCFunction(packet[1]),
		Length:   length,
		Payload:  append([]byte(nil), packet[4:]...),
	}, nil
}

func (b *BVLC) Encode() ([]byte, error) {
	if b == nil {
		return nil, errors.New("nil BVLC")
	}

	return EncodeBVLC(
		b.Function,
		b.Payload,
	)
}

func DecodeResult(b *BVLC) (*BVLCResult, error) {
	if b == nil {
		return nil, errors.New("nil BVLC")
	}

	if b.Function != defs.BVLCFunctionResult {
		return nil, fmt.Errorf(
			"BVLC function is %v, not BVLC-Result",
			b.Function,
		)
	}

	if len(b.Payload) != 2 {
		return nil, fmt.Errorf(
			"invalid BVLC-Result payload length: %d",
			len(b.Payload),
		)
	}

	return &BVLCResult{
		Code: binary.BigEndian.Uint16(b.Payload),
	}, nil
}

func DecodeForwardedNPDU(
	b *BVLC,
) (*ForwardedNPDU, error) {
	if b == nil {
		return nil, errors.New("nil BVLC")
	}

	if b.Function != defs.BVLCFunctionForwardedNPDU {
		return nil, fmt.Errorf(
			"BVLC function is %v, not Forwarded-NPDU",
			b.Function,
		)
	}

	if len(b.Payload) < 6 {
		return nil, fmt.Errorf(
			"invalid Forwarded-NPDU payload length: %d",
			len(b.Payload),
		)
	}

	var address IPAddress
	copy(address.IP[:], b.Payload[0:4])

	address.Port = binary.BigEndian.Uint16(
		b.Payload[4:6],
	)

	return &ForwardedNPDU{
		OriginatingAddress: address,
		NPDU: append(
			[]byte(nil),
			b.Payload[6:]...,
		),
	}, nil
}
