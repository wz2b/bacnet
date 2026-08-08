package defs

type NotificationType byte

// Notify constants.
const (
	NotifyAlarm           NotificationType = 0x00
	NotifyEvent           NotificationType = 0x01
	NotifyACKNotification NotificationType = 0x02
)
