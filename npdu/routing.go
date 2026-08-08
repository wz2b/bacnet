package npdu

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/wz2b/bacnet"
	"github.com/wz2b/bacnet/defs"
)

type RouterPort struct {
	Dnet     uint16      /**< The DNET number that identifies this port. */
	Id       byte        /**< Either 0 or some ill-defined, meaningless value. */
	Info     []byte      /**< Info like 'modem dialing string' */
	Info_len byte        /**< Length of info[]. */
	Next     *RouterPort /**< Point to next in linked list */
}

type RoutingTableRequest struct {
	// Empty means "return your current routing table".
	//
	// Later, if we ever support actually initializing a routing table,
	// this would contain the port mappings to install.
	Ports []RoutingTableEntry
}

type RoutingTableEntry struct {
	DNET     uint16
	PortID   byte
	PortInfo []byte
}

type RoutingTable struct {
	Entries []RoutingTableEntry
}

func (pdu *NPDU) IsRoutingTableAck() bool {
	return pdu.NetworkLayer &&
		pdu.MessageType == byte(defs.NetworkMessageInitializeRoutingTableACK)
}

func (pdu *NPDU) ToRoutingTableAck() (*RoutingTable, error) {
	if pdu == nil {
		return nil, errors.New("nil NPDU")
	}

	if !pdu.IsRoutingTableAck() {
		return nil, fmt.Errorf(
			"NPDU is not an Initialize-Routing-Table-Ack: message type 0x%02X",
			pdu.MessageType,
		)
	}

	if pdu.Length < 0 || pdu.Length > len(pdu.PDU) {
		return nil, fmt.Errorf(
			"invalid NPDU payload offset %d for PDU length %d",
			pdu.Length,
			len(pdu.PDU),
		)
	}

	payload := pdu.PDU[pdu.Length:]
	if len(payload) < 1 {
		return nil, errors.New("routing table ACK missing number-of-ports")
	}

	portCount := int(payload[0])
	offset := 1

	table := &RoutingTable{
		Entries: make([]RoutingTableEntry, 0, portCount),
	}

	for i := 0; i < portCount; i++ {
		// DNET(2) + PortID(1) + PortInfoLength(1)
		if len(payload)-offset < 4 {
			return nil, fmt.Errorf(
				"routing table entry %d truncated: need at least 4 bytes, have %d",
				i,
				len(payload)-offset,
			)
		}

		dnet := binary.BigEndian.Uint16(payload[offset : offset+2])
		offset += 2

		portID := payload[offset]
		offset++

		portInfoLength := int(payload[offset])
		offset++

		if len(payload)-offset < portInfoLength {
			return nil, fmt.Errorf(
				"routing table entry %d truncated port info: need %d bytes, have %d",
				i,
				portInfoLength,
				len(payload)-offset,
			)
		}

		portInfo := make([]byte, portInfoLength)
		copy(portInfo, payload[offset:offset+portInfoLength])
		offset += portInfoLength

		table.Entries = append(table.Entries, RoutingTableEntry{
			DNET:     dnet,
			PortID:   portID,
			PortInfo: portInfo,
		})
	}

	if offset != len(payload) {
		return nil, fmt.Errorf(
			"routing table ACK has %d trailing bytes",
			len(payload)-offset,
		)
	}

	return table, nil
}

func NewRoutingTableInitRequest(count int) (*RoutingTableRequest, error) {
	if count < 0 || count > 255 {
		return nil, fmt.Errorf(
			"routing table port count %d out of range",
			count,
		)
	}

	return &RoutingTableRequest{
		Ports: make([]RoutingTableEntry, count),
	}, nil
}

func (r *RoutingTableRequest) ToNPDU(dest *bacnet.Address) (*NPDU, error) {
	return r.ToNPDUDest(dest)
}

func (r *RoutingTableRequest) ToNPDUDest(dest *bacnet.Address) (*NPDU, error) {
	if r == nil {
		return nil, fmt.Errorf("nil routing table request")
	}

	if len(r.Ports) > 255 {
		return nil, fmt.Errorf(
			"routing table port count %d exceeds maximum 255",
			len(r.Ports),
		)
	}

	//
	// Build the Initialize-Routing-Table message payload.
	//
	// The first byte is Number of Ports.  A value of zero is
	// the routing-table query form.
	//
	payloadLength := 1

	for _, entry := range r.Ports {
		if len(entry.PortInfo) > 255 {
			return nil, fmt.Errorf(
				"port info for DNET %d is too long: %d bytes",
				entry.DNET,
				len(entry.PortInfo),
			)
		}

		// DNET(2) + PortID(1) + PortInfoLength(1) + PortInfo
		payloadLength += 4 + len(entry.PortInfo)
	}

	payload := make([]byte, payloadLength)

	offset := 0

	payload[offset] = byte(len(r.Ports))
	offset++

	for _, entry := range r.Ports {
		binary.BigEndian.PutUint16(
			payload[offset:offset+2],
			entry.DNET,
		)
		offset += 2

		payload[offset] = entry.PortID
		offset++

		payload[offset] = byte(len(entry.PortInfo))
		offset++

		copy(
			payload[offset:offset+len(entry.PortInfo)],
			entry.PortInfo,
		)
		offset += len(entry.PortInfo)
	}

	//
	// NPDU.Encode() writes the NPDU header directly into PDU,
	// starting at PDU[0].  Therefore PDU must already contain
	// enough writable space for both the header and our payload.
	//
	// 64 bytes is more than enough room for the NPDU header;
	// payloadLength accounts for the network-message payload.
	//
	pdu := make([]byte, 64+payloadLength)

	npdu := NewNetworkLayerNPDU(
		pdu,
		dest,
		nil,
		false,
		defs.MessagePriorityNormal,
	)

	npdu.MessageType = byte(defs.NetworkMessageInitializeRoutingTable)
	npdu.HopCount = 255

	//
	// Encode the NPDU header.  Encode() sets npdu.Length to the
	// first byte following the NPDU header/message type.
	//
	npdu.Encode()

	required := npdu.Length + len(payload)
	if required > len(npdu.PDU) {
		return nil, fmt.Errorf(
			"routing table NPDU requires %d bytes, buffer has %d",
			required,
			len(npdu.PDU),
		)
	}

	//
	// Append the Initialize-Routing-Table-specific payload.
	//
	copy(
		npdu.PDU[npdu.Length:required],
		payload,
	)

	//
	// PDU should contain only bytes that are actually transmitted.
	//
	npdu.PDU = npdu.PDU[:required]

	return npdu, nil
}

func (r *RoutingTableRequest) Encode() (*NPDU, error) {
	npdu, err := r.ToNPDU(nil)
	if err != nil {
		return nil, err
	}

	return npdu.Encode(), nil
}
