package npdu

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/wz2b/bacnet/bactypes"
	"github.com/wz2b/bacnet/codec"
	"github.com/wz2b/bacnet/defs"
)

type NPDU struct {
	ProtocolVersion byte
	ExpectingReply  bool
	NetworkLayer    bool
	Priority        byte

	MessageType byte
	VendorID    uint16
	HopCount    byte

	Dest   *bactypes.Address
	Source *bactypes.Address

	// PDU contains the complete encoded NPDU.
	PDU []byte

	// Length is the offset within PDU where the NPDU payload begins.
	//
	// For an application NPDU, PDU[Length:] is the APDU.
	// For a network-layer message, PDU[Length:] is the
	// message-specific network-layer payload.
	Length int
}

func NewNPDU(
	payload []byte,
	dest, src *bactypes.Address,
	expectReply bool,
	priority byte,
) *NPDU {
	return &NPDU{
		PDU:             append([]byte(nil), payload...),
		Dest:            dest,
		Source:          src,
		ExpectingReply:  expectReply,
		Priority:        priority,
		ProtocolVersion: 0x01,
		HopCount:        defs.HopCountDefault,
	}
}

func NewNetworkLayerNPDU(
	payload []byte,
	dest, src *bactypes.Address,
	expectReply bool,
	priority byte,
) *NPDU {
	n := NewNPDU(
		payload,
		dest,
		src,
		expectReply,
		priority,
	)

	n.NetworkLayer = true

	return n
}

func (n *NPDU) Encode() ([]byte, error) {
	if n == nil {
		return nil, errors.New("nil NPDU")
	}

	if n.Priority > 3 {
		return nil, fmt.Errorf(
			"invalid NPDU priority: %d",
			n.Priority,
		)
	}

	control := n.Priority & 0x03

	if n.NetworkLayer {
		control |= codec.BIT7
	}

	if n.Dest != nil && n.Dest.Net != 0 {
		control |= codec.BIT5
	}

	if n.Source != nil &&
		n.Source.Net != 0 &&
		n.Source.Len != 0 {
		control |= codec.BIT3
	}

	if n.ExpectingReply {
		control |= codec.BIT2
	}

	header := make([]byte, 0, 32)

	header = append(
		header,
		n.ProtocolVersion,
		control,
	)

	if n.Dest != nil && n.Dest.Net != 0 {
		if int(n.Dest.Len) > len(n.Dest.Adr) {
			return nil, fmt.Errorf(
				"destination address length %d exceeds available address bytes %d",
				n.Dest.Len,
				len(n.Dest.Adr),
			)
		}

		var network [2]byte
		binary.BigEndian.PutUint16(
			network[:],
			n.Dest.Net,
		)

		header = append(
			header,
			network[:]...,
		)

		header = append(
			header,
			n.Dest.Len,
		)

		header = append(
			header,
			n.Dest.Adr[:n.Dest.Len]...,
		)
	}

	if n.Source != nil &&
		n.Source.Net != 0 &&
		n.Source.Len != 0 {
		if int(n.Source.Len) > len(n.Source.Adr) {
			return nil, fmt.Errorf(
				"source address length %d exceeds available address bytes %d",
				n.Source.Len,
				len(n.Source.Adr),
			)
		}

		var network [2]byte
		binary.BigEndian.PutUint16(
			network[:],
			n.Source.Net,
		)

		header = append(
			header,
			network[:]...,
		)

		header = append(
			header,
			n.Source.Len,
		)

		header = append(
			header,
			n.Source.Adr[:n.Source.Len]...,
		)
	}

	if n.Dest != nil && n.Dest.Net != 0 {
		header = append(
			header,
			n.HopCount,
		)
	}

	if n.NetworkLayer {
		header = append(
			header,
			n.MessageType,
		)

		if n.MessageType >= 0x80 {
			var vendor [2]byte
			binary.BigEndian.PutUint16(
				vendor[:],
				n.VendorID,
			)

			header = append(
				header,
				vendor[:]...,
			)
		}
	}

	result := make(
		[]byte,
		0,
		len(header)+len(n.PDU),
	)

	result = append(
		result,
		header...,
	)

	n.Length = len(header)

	result = append(
		result,
		n.PDU...,
	)

	return result, nil
}

func Decode(data []byte) (*NPDU, error) {
	if len(data) < 2 {
		return nil, fmt.Errorf(
			"NPDU too short: %d bytes",
			len(data),
		)
	}

	result := &NPDU{
		PDU: append([]byte(nil), data...),
	}

	result.ProtocolVersion = data[0]

	if result.ProtocolVersion != 0x01 {
		return nil, fmt.Errorf(
			"unsupported NPDU protocol version: 0x%02X",
			result.ProtocolVersion,
		)
	}

	control := data[1]

	result.NetworkLayer =
		control&codec.BIT7 != 0

	result.ExpectingReply =
		control&codec.BIT2 != 0

	result.Priority =
		control & 0x03

	offset := 2

	//
	// Destination
	//

	if control&codec.BIT5 != 0 {
		if len(data)-offset < 3 {
			return nil, errors.New(
				"truncated NPDU destination address",
			)
		}

		network := binary.BigEndian.Uint16(
			data[offset : offset+2],
		)
		offset += 2

		addressLength := int(data[offset])
		offset++

		if addressLength > bactypes.MAX_MAC_LEN {
			return nil, fmt.Errorf(
				"destination address length %d exceeds maximum %d",
				addressLength,
				bactypes.MAX_MAC_LEN,
			)
		}

		if len(data)-offset < addressLength {
			return nil, errors.New(
				"truncated NPDU destination address",
			)
		}

		dest := bactypes.NewAddress()
		dest.Net = network
		dest.Len = byte(addressLength)

		copy(
			dest.Adr,
			data[offset:offset+addressLength],
		)

		offset += addressLength

		result.Dest = dest
	}

	//
	// Source
	//

	if control&codec.BIT3 != 0 {
		if len(data)-offset < 3 {
			return nil, errors.New(
				"truncated NPDU source address",
			)
		}

		network := binary.BigEndian.Uint16(
			data[offset : offset+2],
		)
		offset += 2

		addressLength := int(data[offset])
		offset++

		if addressLength > bactypes.MAX_MAC_LEN {
			return nil, fmt.Errorf(
				"source address length %d exceeds maximum %d",
				addressLength,
				bactypes.MAX_MAC_LEN,
			)
		}

		if len(data)-offset < addressLength {
			return nil, errors.New(
				"truncated NPDU source address",
			)
		}

		source := bactypes.NewAddress()
		source.Net = network
		source.Len = byte(addressLength)

		copy(
			source.Adr,
			data[offset:offset+addressLength],
		)

		offset += addressLength

		result.Source = source
	}

	//
	// Hop count is present whenever destination information is present.
	//

	if control&codec.BIT5 != 0 {
		if offset >= len(data) {
			return nil, errors.New(
				"truncated NPDU hop count",
			)
		}

		result.HopCount = data[offset]
		offset++
	}

	//
	// Network-layer message information
	//

	if result.NetworkLayer {
		if offset >= len(data) {
			return nil, errors.New(
				"truncated NPDU network message type",
			)
		}

		result.MessageType = data[offset]
		offset++

		if result.MessageType >= 0x80 {
			if len(data)-offset < 2 {
				return nil, errors.New(
					"truncated NPDU vendor ID",
				)
			}

			result.VendorID = binary.BigEndian.Uint16(
				data[offset : offset+2],
			)
			offset += 2
		}
	}

	result.Length = offset

	return result, nil
}

func (n *NPDU) Payload() []byte {
	if n == nil {
		return nil
	}

	if n.Length < 0 ||
		n.Length > len(n.PDU) {
		return nil
	}

	return n.PDU[n.Length:]
}

func (n *NPDU) APDU() ([]byte, bool) {
	if n == nil || n.NetworkLayer {
		return nil, false
	}

	return n.Payload(), true
}

func (n *NPDU) NetworkMessageData() ([]byte, bool) {
	if n == nil || !n.NetworkLayer {
		return nil, false
	}

	return n.Payload(), true
}
