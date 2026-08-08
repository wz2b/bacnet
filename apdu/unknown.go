package apdu

import (
	"fmt"

	"github.com/wz2b/bacnet/defs"
)

type UnknownAPDUService struct {
	APDU APDU
}

func (u UnknownAPDUService) String() string {
	switch u.APDU.Type {
	case defs.PDUTypeUnconfirmedServiceRequest:
		return fmt.Sprintf(
			"Unknown unconfirmed APDU service: service=%d invokeID=%d segmented=%t data=% X",
			u.APDU.UnconfirmedServiceChoice,
			u.APDU.InvokeID,
			u.APDU.Segmented,
			u.APDU.Data,
		)

	case defs.PDUTypeConfirmedServiceRequest,
		defs.PDUTypeComplexACK:
		return fmt.Sprintf(
			"Unknown confirmed APDU service: service=%d invokeID=%d segmented=%t data=% X",
			u.APDU.ConfirmedServiceChoice,
			u.APDU.InvokeID,
			u.APDU.Segmented,
			u.APDU.Data,
		)

	default:
		return fmt.Sprintf(
			"Unknown APDU service: type=0x%02X data=% X",
			byte(u.APDU.Type),
			u.APDU.Data,
		)
	}
}
