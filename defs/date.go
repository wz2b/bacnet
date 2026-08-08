package defs

type DayOfWeekEnum byte

// Date constants.
const (
	DaysOfWeekMonday    DayOfWeekEnum = 0x00
	DaysOfWeekTuesday   DayOfWeekEnum = 0x01
	DaysOfWeekWednesday DayOfWeekEnum = 0x02
	DaysOfWeekThursday  DayOfWeekEnum = 0x03
	DaysOfWeekFriday    DayOfWeekEnum = 0x04
	DaysOfWeekSaturday  DayOfWeekEnum = 0x05
	DaysOfWeekSunday    DayOfWeekEnum = 0x06
	MaxDaysOfWeek       DayOfWeekEnum = 0x07
)
