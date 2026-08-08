package objects

import "github.com/wz2b/bacnet/bactypes"

type NotificationClass struct {
	bactypes.DefaultObject
	NotificationClass string `json:"notificationClass"`
	Priority          string `json:"priority"`
	AckRequired       string `json:"ackRequired"`
	RecipientList     string `json:"recipientList"`
	Description       string `json:"description"`
	ProfileName       string `json:"profileName"`
}
