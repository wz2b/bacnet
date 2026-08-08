package bacnet

import (
	"github.com/wz2b/bacnet/apdu"
	"github.com/wz2b/bacnet/bactypes"
	"github.com/wz2b/bacnet/bvlc"
	"github.com/wz2b/bacnet/npdu"
)

type Message struct {
	BVLC  *bvlc.BVLC
	NPDU  *npdu.NPDU
	APDU  *apdu.APDU
	Value any

	Source *bactypes.Address
}
