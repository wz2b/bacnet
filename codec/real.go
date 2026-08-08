package codec

import (
	"encoding/binary"
	"math"

	"github.com/wz2b/bacnet/defs"
)

func EncodeApplicationTaggedReal(value float32) []byte {
	result := make([]byte, 5)

	result[0] =
		(defs.ApplicationTagReal << 4) |
			4

	binary.BigEndian.PutUint32(
		result[1:5],
		math.Float32bits(value),
	)

	return result
}
