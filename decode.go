package bacnet

import (
	"fmt"

	"github.com/wz2b/bacnet/apdu"
	"github.com/wz2b/bacnet/bvlc"
	"github.com/wz2b/bacnet/defs"
	"github.com/wz2b/bacnet/npdu"
	"github.com/wz2b/bacnet/services"
)

func Decode(data []byte) (Message, error) {
	var result Message

	if len(data) == 0 {
		return result, fmt.Errorf("empty BACnet packet")
	}

	//
	// BVLC
	//

	b, err := bvlc.Decode(data)
	if err != nil {
		return result, fmt.Errorf("decode BVLC: %w", err)
	}

	result.BVLC = b

	//
	// BVLC messages that do not contain an NPDU will eventually
	// terminate here.
	//

	switch b.Function {
	// case defs.BVLCFunctionReadBroadcastDistributionTableACK:
	//     ...
	//
	// case defs.BVLCFunctionReadForeignDeviceTableACK:
	//     ...
	//
	// case defs.BVLCFunctionResult:
	//     ...
	}

	//
	// NPDU
	//

	npduData := b.Payload

	if b.Function == defs.BVLCFunctionForwardedNPDU {
		forwarded, err := bvlc.DecodeForwardedNPDU(b)
		if err != nil {
			return result, fmt.Errorf(
				"decode Forwarded-NPDU: %w",
				err,
			)
		}

		npduData = forwarded.NPDU
	}

	n, err := npdu.Decode(npduData)
	if err != nil {
		return result, fmt.Errorf("decode NPDU: %w", err)
	}

	result.NPDU = n

	//
	// Network-layer messages terminate here.
	//
	// Until the network-message dispatcher exists, preserve the
	// successfully decoded NPDU and its raw message-specific data.
	//

	if n.NetworkLayer {
		result.Value = &npdu.UnknownNetworkMessage{
			NPDU: *n,
			Data: append(
				[]byte(nil),
				n.Payload()...,
			),
		}

		return result, nil
	}

	//
	// APDU
	//

	a, err := apdu.Decode(n.Payload())
	if err != nil {
		return result, fmt.Errorf("decode APDU: %w", err)
	}

	result.APDU = &a

	//
	// Application service
	//

	value, err := services.Decode(&a)
	if err != nil {
		return result, fmt.Errorf("decode BACnet service: %w", err)
	}

	result.Value = value

	return result, nil
}
