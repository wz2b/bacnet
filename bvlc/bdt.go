package bvlc

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/wz2b/bacnet/defs"
)

type WriteBroadcastDistributionTable struct {
	Entries []BroadcastDistributionTableEntry
}

type ReadBroadcastDistributionTable struct{}

type ReadBroadcastDistributionTableAck struct {
	Entries []BroadcastDistributionTableEntry
}

func (p *ReadBroadcastDistributionTable) Encode() ([]byte, error) {
	if p == nil {
		return nil, errors.New("nil ReadBroadcastDistributionTable")
	}

	b, err := p.BVLC()
	if err != nil {
		return nil, err
	}

	return b.Encode()
}

func (p *ReadBroadcastDistributionTable) BVLC() (*BVLC, error) {
	if p == nil {
		return nil, errors.New("nil ReadBroadcastDistributionTable")
	}

	return &BVLC{
		BVLLType: BVLCTypeBACnetIP,
		Function: defs.BVLCFunctionReadBroadcastDistributionTable,
	}, nil
}

func DecodeReadBroadcastDistributionTableAck(
	b *BVLC,
) (*ReadBroadcastDistributionTableAck, error) {
	if b == nil {
		return nil, errors.New("nil BVLC")
	}

	if b.Function != defs.BVLCFunctionReadBroadcastDistributionTableACK {
		return nil, fmt.Errorf(
			"BVLC function is %v, not Read-Broadcast-Distribution-Table-Ack",
			b.Function,
		)
	}

	// Each BDT entry is:
	//
	// 4 bytes IPv4 address
	// 2 bytes UDP port
	// 4 bytes broadcast distribution mask
	//
	// = 10 bytes total
	if len(b.Payload)%10 != 0 {
		return nil, fmt.Errorf(
			"invalid Read-Broadcast-Distribution-Table-Ack payload length: %d",
			len(b.Payload),
		)
	}

	result := &ReadBroadcastDistributionTableAck{
		Entries: make(
			[]BroadcastDistributionTableEntry,
			0,
			len(b.Payload)/10,
		),
	}

	for offset := 0; offset < len(b.Payload); offset += 10 {
		var address IPAddress

		copy(
			address.IP[:],
			b.Payload[offset:offset+4],
		)

		address.Port = binary.BigEndian.Uint16(
			b.Payload[offset+4 : offset+6],
		)

		entry := BroadcastDistributionTableEntry{
			Address: address,
			Mask: binary.BigEndian.Uint32(
				b.Payload[offset+6 : offset+10],
			),
		}

		result.Entries = append(result.Entries, entry)
	}

	return result, nil
}
