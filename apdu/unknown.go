package apdu

import "fmt"

type UnknownAPDUService struct {
	APDU APDU
}

func (u UnknownAPDUService) String() string {
	return fmt.Sprintf(
		"Unknown APDU service: type=0x%02X service=%d invokeID=%d segmented=%t data=% X",
		byte(u.APDU.Type),
		u.APDU.ServiceChoice,
		u.APDU.InvokeID,
		u.APDU.Segmented,
		u.APDU.Data,
	)
}
