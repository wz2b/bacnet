package defs

type Segmentation byte

const (
	SegmentationBoth     Segmentation = 0x00
	SegmentationTransmit Segmentation = 0x01
	SegmentationReceive  Segmentation = 0x02
	SegmentationNone     Segmentation = 0x03
)

func (s Segmentation) Valid() bool {
	return s <= SegmentationNone
}
