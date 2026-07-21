package bacnet

import (
	"context"
	"net"
)

type IPTransport interface {
	SendUnicast(dst *net.UDPAddr, npdu []byte) error
	SendBroadcast(npdu []byte) error
	Receive(ctx context.Context) ([]byte, *net.UDPAddr, error)
}
