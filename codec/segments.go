package codec

/* from clause 20.1.2.4 max-segments-accepted */
/* and clause 20.1.2.5 max-APDU-length-accepted */
/* returns the encoded octet */
func EncodeMaxSegsMaxAPDU(max_segs int, max_apdu uint32) byte {
	var octet byte = 0

	if max_segs < 2 {
		octet = 0
	} else if max_segs < 4 {
		octet = 0x10
	} else if max_segs < 8 {
		octet = 0x20
	} else if max_segs < 16 {
		octet = 0x30
	} else if max_segs < 32 {
		octet = 0x40
	} else if max_segs < 64 {
		octet = 0x50
	} else if max_segs == 64 {
		octet = 0x60
	} else {
		octet = 0x70
	}

	if max_apdu <= 50 {
		octet |= 0x00
	} else if max_apdu <= 128 {
		octet |= 0x01
	} else if max_apdu <= 206 {
		octet |= 0x02
	} else if max_apdu <= 480 {
		octet |= 0x03
	} else if max_apdu <= 1024 {
		octet |= 0x04
	} else if max_apdu <= 1476 {
		octet |= 0x05
	}

	return octet
}

func DecodeMaxSegsMaxAPDU(value byte) (int, uint32) {
	maxSegsCode := (value >> 4) & 0x07
	maxAPDUCode := value & 0x0F

	var maxSegs int

	switch maxSegsCode {
	case 0:
		maxSegs = 0
	case 1:
		maxSegs = 2
	case 2:
		maxSegs = 4
	case 3:
		maxSegs = 8
	case 4:
		maxSegs = 16
	case 5:
		maxSegs = 32
	case 6:
		maxSegs = 64
	case 7:
		maxSegs = 65
	}

	var maxAPDU uint32

	switch maxAPDUCode {
	case 0:
		maxAPDU = 50
	case 1:
		maxAPDU = 128
	case 2:
		maxAPDU = 206
	case 3:
		maxAPDU = 480
	case 4:
		maxAPDU = 1024
	case 5:
		maxAPDU = 1476
	default:
		maxAPDU = 0
	}

	return maxSegs, maxAPDU
}
