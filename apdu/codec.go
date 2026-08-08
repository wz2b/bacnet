package apdu

import (
	"fmt"

	"github.com/wz2b/bacnet/codec"
	"github.com/wz2b/bacnet/defs"
)

func Encode(a APDU) ([]byte, error) {
	switch a.Type {
	case defs.PDUTypeUnconfirmedServiceRequest:
		result := make([]byte, 2+len(a.Data))

		result[0] = byte(a.Type)
		result[1] = a.ServiceChoice
		copy(result[2:], a.Data)

		return result, nil

	case defs.PDUTypeConfirmedServiceRequest:
		result := make([]byte, 4+len(a.Data))

		result[0] = byte(a.Type)

		if a.Segmented {
			// set segmented-message bit
		}

		result[1] = codec.EncodeMaxSegsMaxAPDU(
			int(a.MaxSegments),
			a.MaxAPDU,
		)

		result[2] = a.InvokeID
		result[3] = a.ServiceChoice
		copy(result[4:], a.Data)

		return result, nil

	case defs.PDUTypeComplexACK:
		result := make([]byte, 3+len(a.Data))

		result[0] = byte(a.Type)
		result[1] = a.InvokeID
		result[2] = a.ServiceChoice

		copy(result[3:], a.Data)

		return result, nil

	default:
		return nil, fmt.Errorf(
			"unsupported APDU type: 0x%02X",
			byte(a.Type),
		)
	}
}

func Decode(data []byte) (APDU, error) {
	var result APDU

	if len(data) == 0 {
		return result, fmt.Errorf("empty APDU")
	}

	result.Type = defs.PDUType(data[0] & 0xF0)

	switch result.Type {
	case defs.PDUTypeUnconfirmedServiceRequest:
		if len(data) < 2 {
			return result, fmt.Errorf("truncated unconfirmed request APDU")
		}

		result.ServiceChoice = data[1]
		result.Data = append([]byte(nil), data[2:]...)

		return result, nil

	case defs.PDUTypeConfirmedServiceRequest:
		if len(data) < 4 {
			return result, fmt.Errorf("truncated confirmed request APDU")
		}

		maxSegments, maxAPDU := codec.DecodeMaxSegsMaxAPDU(data[1])

		result.MaxSegments = uint8(maxSegments)
		result.MaxAPDU = maxAPDU

		result.InvokeID = data[2]
		result.ServiceChoice = data[3]
		result.Data = append([]byte(nil), data[4:]...)

		return result, nil

	case defs.PDUTypeComplexACK:
		if len(data) < 3 {
			return result, fmt.Errorf("truncated ComplexACK APDU")
		}

		result.Segmented = data[0]&0x08 != 0

		if result.Segmented {
			return result, fmt.Errorf(
				"segmented ComplexACK APDU not yet supported",
			)
		}

		result.InvokeID = data[1]
		result.ServiceChoice = data[2]
		result.Data = append([]byte(nil), data[3:]...)

		return result, nil

	default:
		return result, fmt.Errorf(
			"unsupported APDU type: 0x%02X",
			byte(result.Type),
		)
	}
}
