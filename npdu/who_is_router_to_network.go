package npdu

import (
	"fmt"

	"github.com/wz2b/bacnet/codec"
)

type WhoIsRouterToNetwork struct {
	// NetworkNumber is nil when no particular network was requested.
	// A nil value means "all reachable networks."
	NetworkNumber *uint16
}

func (w WhoIsRouterToNetwork) Encode() []byte {
	if w.NetworkNumber == nil {
		return nil
	}

	data := make([]byte, 2)

	codec.EncodeUnsigned16(
		data,
		*w.NetworkNumber,
	)

	return data
}

func DecodeWhoIsRouterToNetwork(
	data []byte,
) (*WhoIsRouterToNetwork, error) {
	result := &WhoIsRouterToNetwork{}

	switch len(data) {
	case 0:
		// An empty payload asks for all reachable networks.
		return result, nil

	case 2:
		n, networkNumber := codec.DecodeUnsigned16(data)
		if n != 2 {
			return nil, fmt.Errorf(
				"invalid Who-Is-Router-To-Network network number",
			)
		}

		result.NetworkNumber = &networkNumber

		return result, nil

	default:
		return nil, fmt.Errorf(
			"Who-Is-Router-To-Network payload must be 0 or 2 bytes, got %d",
			len(data),
		)
	}
}
