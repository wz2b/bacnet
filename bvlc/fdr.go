package bvlc

import (
	"encoding/binary"
	"fmt"

	"github.com/wz2b/bacnet"
	"github.com/wz2b/bacnet/defs"
)

/*
 * Read-Foreign-Device-Table
 */

func (r ReadForeignDeviceTable) BVLC() BVLC {
	return BVLC{
		BVLLType: BVLCTypeBACnetIP,
		Function: defs.BVLCFunctionReadForeignDeviceTable,
	}
}

func DecodeReadForeignDeviceTable(
	b *BVLC,
) (*ReadForeignDeviceTable, error) {
	if b == nil {
		return nil, fmt.Errorf(
			"decode Read-Foreign-Device-Table: nil BVLC",
		)
	}

	if b.Function != defs.BVLCFunctionReadForeignDeviceTable {
		return nil, fmt.Errorf(
			"decode Read-Foreign-Device-Table: unexpected BVLC function %v",
			b.Function,
		)
	}

	if len(b.Payload) != 0 {
		return nil, fmt.Errorf(
			"decode Read-Foreign-Device-Table: expected empty payload, got %d bytes",
			len(b.Payload),
		)
	}

	return &ReadForeignDeviceTable{}, nil
}

/*
 * Read-Foreign-Device-Table-Ack
 */

func (r ReadForeignDeviceTableAck) BVLC() BVLC {
	const entryLength = 10

	payload := make(
		[]byte,
		len(r.Entries)*entryLength,
	)

	offset := 0

	for _, entry := range r.Entries {
		copy(
			payload[offset:offset+4],
			entry.Address.IP[:],
		)
		offset += 4

		binary.BigEndian.PutUint16(
			payload[offset:offset+2],
			entry.Address.Port,
		)
		offset += 2

		binary.BigEndian.PutUint16(
			payload[offset:offset+2],
			entry.TTL,
		)
		offset += 2

		binary.BigEndian.PutUint16(
			payload[offset:offset+2],
			entry.RemainingTime,
		)
		offset += 2
	}

	return BVLC{
		BVLLType: BVLCTypeBACnetIP,
		Function: defs.BVLCFunctionReadForeignDeviceTableACK,
		Payload:  payload,
	}
}

func DecodeReadForeignDeviceTableAck(
	b *BVLC,
) (*ReadForeignDeviceTableAck, error) {
	if b == nil {
		return nil, fmt.Errorf(
			"decode Read-Foreign-Device-Table-Ack: nil BVLC",
		)
	}

	if b.Function != defs.BVLCFunctionReadForeignDeviceTableACK {
		return nil, fmt.Errorf(
			"decode Read-Foreign-Device-Table-Ack: unexpected BVLC function %v",
			b.Function,
		)
	}

	const entryLength = 10

	if len(b.Payload)%entryLength != 0 {
		return nil, fmt.Errorf(
			"decode Read-Foreign-Device-Table-Ack: payload length %d is not a multiple of %d",
			len(b.Payload),
			entryLength,
		)
	}

	result := &ReadForeignDeviceTableAck{
		Entries: make(
			[]ForeignDeviceTableEntry,
			0,
			len(b.Payload)/entryLength,
		),
	}

	for offset := 0; offset < len(b.Payload); offset += entryLength {
		var address bacnet.IPAddress

		copy(
			address.IP[:],
			b.Payload[offset:offset+4],
		)

		address.Port = binary.BigEndian.Uint16(
			b.Payload[offset+4 : offset+6],
		)

		entry := ForeignDeviceTableEntry{
			Address: address,
			TTL: binary.BigEndian.Uint16(
				b.Payload[offset+6 : offset+8],
			),
			RemainingTime: binary.BigEndian.Uint16(
				b.Payload[offset+8 : offset+10],
			),
		}

		result.Entries = append(
			result.Entries,
			entry,
		)
	}

	return result, nil
}
