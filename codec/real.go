package codec

import (
	"encoding/binary"
	"math"

	"github.com/wz2b/bacnet/defs"
)

func DecodeApplicationTaggedReal(data []byte) (float64, bool) {
	if len(data) != 5 {
		return 0, false
	}

	expectedTag :=
		byte(defs.ApplicationTagReal<<4) | 4

	if data[0] != expectedTag {
		return 0, false
	}

	bits := binary.BigEndian.Uint32(data[1:5])
	return float64(math.Float32frombits(bits)), true
}

func EncodeApplicationTaggedReal(value float64) []byte {
	data := make([]byte, 5)

	data[0] =
		byte(defs.ApplicationTagReal<<4) | 4

	bits := math.Float32bits(float32(value))
	binary.BigEndian.PutUint32(data[1:], bits)

	return data
}
