package bacnet

import (
	"encoding/binary"
	"fmt"
)

const BVLCTypeBACnetIP byte = 0x81

type BVLCFunction byte

const (
	BVLCResult                        BVLCFunction = 0x00
	WriteBroadcastDistributionTable   BVLCFunction = 0x01
	ReadBroadcastDistributionTable    BVLCFunction = 0x02
	ReadBroadcastDistributionTableAck BVLCFunction = 0x03
	ForwardedNPDU                     BVLCFunction = 0x04
	RegisterForeignDevice             BVLCFunction = 0x05
	ReadForeignDeviceTable            BVLCFunction = 0x06
	ReadForeignDeviceTableAck         BVLCFunction = 0x07
	DeleteForeignDeviceTableEntry     BVLCFunction = 0x08
	DistributeBroadcastToNetwork      BVLCFunction = 0x09
	OriginalUnicastNPDU               BVLCFunction = 0x0A
	OriginalBroadcastNPDU             BVLCFunction = 0x0B
)

func EncodeBVLC(function BVLCFunction, payload []byte) ([]byte, error) {
	length := 4 + len(payload)
	if length > 0xffff {
		return nil, fmt.Errorf("BVLC packet too large: %d bytes", length)
	}

	packet := make([]byte, length)

	packet[0] = BVLCTypeBACnetIP
	packet[1] = byte(function)
	binary.BigEndian.PutUint16(packet[2:4], uint16(length))
	copy(packet[4:], payload)

	return packet, nil
}

func EncodeRegisterForeignDevice(ttlSeconds uint16) ([]byte, error) {
	payload := make([]byte, 2)
	binary.BigEndian.PutUint16(payload, ttlSeconds)

	return EncodeBVLC(RegisterForeignDevice, payload)
}
