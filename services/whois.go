package services

import (
	"fmt"

	"github.com/wz2b/bacnet/apdu"
	codec "github.com/wz2b/bacnet/codec"
	defs "github.com/wz2b/bacnet/defs"
)

type WhoIs struct {
	LowLimit  *uint32
	HighLimit *uint32
}

func (self *WhoIs) SetLimits(low_limit, high_limit uint32) *WhoIs {
	self.LowLimit = &low_limit
	self.HighLimit = &high_limit
	return self
}

func (w WhoIs) APDU() []byte {
	pdu := make([]byte, 2+10)

	encodeIdx := 2

	pdu[0] = byte(defs.PDUTypeUnconfirmedServiceRequest)
	pdu[1] = byte(defs.ServiceUnconfirmedWhoIs)

	if w.LowLimit != nil &&
		w.HighLimit != nil &&
		*w.LowLimit <= uint32(defs.MaxInstance) &&
		*w.HighLimit <= uint32(defs.MaxInstance) {

		encodeIdx += codec.EncodeContextTaggedUnsigned(
			pdu[encodeIdx:],
			0,
			*w.LowLimit,
		)

		encodeIdx += codec.EncodeContextTaggedUnsigned(
			pdu[encodeIdx:],
			1,
			*w.HighLimit,
		)
	}

	return pdu[:encodeIdx]
}

func DecodeWhoIs(a *apdu.APDU) (*WhoIs, error) {
	if a.Type != defs.PDUTypeUnconfirmedServiceRequest {
		return nil, fmt.Errorf("Who-Is requires an unconfirmed service request APDU")
	}

	if a.UnconfirmedServiceChoice != defs.ServiceUnconfirmedWhoIs {
		return nil, fmt.Errorf("APDU is not a Who-Is request")
	}

	w := &WhoIs{}

	if len(a.Data) == 0 {
		return w, nil
	}

	decodeIdx := 0

	n, tagNumber, length := codec.DecodeTagNumberAndValue(a.Data[decodeIdx:])
	decodeIdx += n

	if tagNumber != 0 || decodeIdx >= len(a.Data) {
		return nil, fmt.Errorf("invalid Who-Is low-limit tag")
	}

	n, value := codec.DecodeUnsigned(a.Data[decodeIdx:], length)
	decodeIdx += n

	if value > uint32(defs.MaxInstance) {
		return nil, fmt.Errorf("Who-Is low limit out of range: %d", value)
	}

	ll := uint32(value)
	w.LowLimit = &ll

	if decodeIdx >= len(a.Data) {
		return nil, fmt.Errorf("Who-Is high limit missing")
	}

	n, tagNumber, length = codec.DecodeTagNumberAndValue(a.Data[decodeIdx:])
	decodeIdx += n

	if tagNumber != 1 || decodeIdx >= len(a.Data) {
		return nil, fmt.Errorf("invalid Who-Is high-limit tag")
	}

	_, value = codec.DecodeUnsigned(a.Data[decodeIdx:], length)

	if value > defs.MaxInstance {
		return nil, fmt.Errorf("Who-Is high limit out of range: %d", value)
	}

	hl := value
	w.HighLimit = &hl

	return w, nil
}
