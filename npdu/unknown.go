package npdu

import "fmt"

type UnknownNetworkMessage struct {
	NPDU NPDU
	Data []byte
}

func (u UnknownNetworkMessage) String() string {
	return fmt.Sprintf(
		"Unknown network message: type=0x%02X data=% X",
		u.NPDU.MessageType,
		u.Data,
	)
}
