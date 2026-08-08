package objects

import "github.com/wz2b/bacnet/bactypes"

type CharacterStringValue struct {
	bactypes.DefaultObject
	PresentValue                 string `json:"presentValue"`
	StatusFlags                  string `json:"statusFlags"`
	Description                  string `json:"description"`
	EventState                   string `json:"eventState"`
	Reliability                  string `json:"reliability"`
	OutOfService                 string `json:"outOfService"`
	PriorityArray                string `json:"priorityArray"`
	RelinquishDefault            string `json:"relinquishDefault"`
	TimeDelay                    string `json:"timeDelay"`
	NotificationClass            string `json:"notificationClass"`
	AlarmValues                  string `json:"alarmValues"`
	FaultValues                  string `json:"faultValues"`
	EventEnable                  string `json:"eventEnable"`
	AckedTransitions             string `json:"ackedTransitions"`
	NotifyType                   string `json:"notifyType"`
	EventTimeStamps              string `json:"eventTimeStamps"`
	EventMessageTexts            string `json:"eventMessageTexts"`
	ReliabilityEvaluationInhibit string `json:"reliabilityEvaluationInhibit"`
	ProfileName                  string `json:"profileName"`
}
