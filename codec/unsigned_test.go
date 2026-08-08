package codec

import (
	"bytes"
	"testing"
)

func TestEncodeApplicationTaggedUnsigned(t *testing.T) {
	tests := []struct {
		value uint32
		want  []byte
	}{
		{0, []byte{0x21, 0x00}},
		{255, []byte{0x21, 0xff}},
		{256, []byte{0x22, 0x01, 0x00}},
		{65535, []byte{0x22, 0xff, 0xff}},
		{65536, []byte{0x23, 0x01, 0x00, 0x00}},
	}

	for _, tt := range tests {
		got := EncodeApplicationTaggedUnsigned(tt.value)
		if !bytes.Equal(got, tt.want) {
			t.Errorf("%d: got % x, want % x", tt.value, got, tt.want)
		}
	}
}
