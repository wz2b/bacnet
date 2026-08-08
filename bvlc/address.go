package bvlc

import "fmt"

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
