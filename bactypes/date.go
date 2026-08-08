package bactypes

const BACNET_WEEKDAY_MONDAY byte = 0x01
const BACNET_WEEKDAY_TUESDAY byte = 0x02
const BACNET_WEEKDAY_WEDNESDAY byte = 0x03
const BACNET_WEEKDAY_THURSDAY byte = 0x04
const BACNET_WEEKDAY_FRIDAY byte = 0x05
const BACNET_WEEKDAY_SATURDAY byte = 0x06
const BACNET_WEEKDAY_SUNDAY byte = 0x07

type Date struct {
	Year    uint16
	Month   byte
	Day     byte
	Weekday byte
}
