package bacnet

import (
	"fmt"
	"strconv"
	"strings"
)

func NewAddress() *Address {
	self := &Address{}
	self.Mac = make([]byte, MAX_MAC_LEN)
	self.Adr = make([]byte, MAX_MAC_LEN)
	return self
}

const MAX_MAC_LEN = 0x07

type Address struct {
	/* mac_len = 0 is a broadcast address */
	Mac_len byte
	/* note: MAC for IP addresses uses 4 bytes for addr, 2 bytes for port */
	/* use de/encode_unsigned32/16 for re/storing the IP address */
	Mac []byte
	/* DNET,DLEN,DADR or SNET,SLEN,SADR */
	/* the following are used if the device is behind a router */
	/* net = 0 indicates local */
	Net uint16 /* BACnet network number */
	/* LEN = 0 denotes broadcast MAC ADR and ADR field is absent */
	/* LEN > 0 specifies length of ADR field */
	Len byte   /* length of MAC address */
	Adr []byte /* hwaddr (MAC) address */
}

/* define a MAC address for manipulation */
type BacnetMacAddress struct {
	Len byte /* length of MAC address */
	Adr []byte
}

/* Define a BACNET IP Address */
type IPAddress struct {
	IP   [4]byte
	Port uint16
}

func (a IPAddress) String() string {
	return fmt.Sprintf(
		"%d.%d.%d.%d:%d",
		a.IP[0],
		a.IP[1],
		a.IP[2],
		a.IP[3],
		a.Port,
	)
}

func ParseAddress(s string) (*Address, error) {
	networkText, macText, ok := strings.Cut(s, ":")
	if !ok {
		return nil, fmt.Errorf(
			"invalid BACnet address %q: expected <network>:<mac>",
			s,
		)
	}

	network, err := strconv.ParseUint(networkText, 10, 16)
	if err != nil {
		return nil, fmt.Errorf(
			"invalid BACnet network %q: %w",
			networkText,
			err,
		)
	}

	mac, err := strconv.ParseUint(macText, 10, 8)
	if err != nil {
		return nil, fmt.Errorf(
			"invalid BACnet MAC %q: %w",
			macText,
			err,
		)
	}

	address := NewAddress()

	address.Net = uint16(network)
	address.Len = 1
	address.Adr[0] = byte(mac)

	return address, nil
}
