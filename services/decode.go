package services

import (
	"fmt"

	"github.com/wz2b/bacnet/apdu"
	"github.com/wz2b/bacnet/defs"
)

func Decode(a *apdu.APDU) (any, error) {
	if a == nil {
		return nil, fmt.Errorf("nil APDU")
	}

	switch a.Type {
	case defs.PDUTypeUnconfirmedServiceRequest:
		switch defs.UnconfirmedServiceChoice(a.UnconfirmedServiceChoice) {
		case defs.ServiceUnconfirmedIAm:
			return DecodeIAm(a)

		case defs.ServiceUnconfirmedIHave:
			return DecodeIHave(a)

		case defs.ServiceUnconfirmedWhoHas:
			return DecodeWhoHas(a)

		case defs.ServiceUnconfirmedWhoIs:
			return DecodeWhoIs(a)

		case defs.ServiceUnconfirmedCOVNotification:
			return DecodeUnconfirmedCOVNotification(a)

		default:
			return &apdu.UnknownAPDUService{
				APDU: *a,
			}, nil
		}

	case defs.PDUTypeConfirmedServiceRequest:
		switch defs.ConfirmedServiceChoice(a.ConfirmedServiceChoice) {
		case defs.ServiceConfirmedAcknowledgeAlarm:
			return DecodeACKAlarm(a)

		case defs.ServiceConfirmedReadProperty:
			return DecodeReadPropertyRequest(a)

		case defs.ServiceConfirmedCOVNotification:
			return DecodeConfirmedCOVNotification(a)

		case defs.ServiceConfirmedSubscribeCOV:
			return DecodeSubscribeCOV(a)

		case defs.ServiceConfirmedSubscribeCOVProperty:
			return DecodeSubscribeCOVProperty(a)

		case defs.ServiceConfirmedAtomicReadFile:
			return DecodeAtomicReadFileRequest(a)

		default:
			return &apdu.UnknownAPDUService{
				APDU: *a,
			}, nil
		}

	case defs.PDUTypeComplexACK:
		switch defs.ConfirmedServiceChoice(a.ConfirmedServiceChoice) {
		case defs.ServiceConfirmedReadProperty:
			return DecodeReadPropertyACK(a)

		case defs.ServiceConfirmedAtomicReadFile:
			return DecodeAtomicReadFileACK(a)

		default:
			return &apdu.UnknownAPDUService{
				APDU: *a,
			}, nil
		}

	default:
		return &apdu.UnknownAPDUService{
			APDU: *a,
		}, nil
	}
}
