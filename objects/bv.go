package objects

import "github.com/wz2b/bacnet/bactypes"

type BinaryValue struct {
	bactypes.DefaultObject
	PresentValue                 string `json:"presentValue"`
	StatusFlags                  string `json:"statusFlags"`
	EventState                   string `json:"eventState"`
	OutOfService                 string `json:"outOfService"`
	Description                  string `json:"description"`
	Reliability                  string `json:"reliability"`
	InactiveText                 string `json:"inactiveText"`
	ActiveText                   string `json:"activeText"`
	ChangeOfStateTime            string `json:"changeOfStateTime"`
	ChangeOfStateCount           string `json:"changeOfStateCount"`
	TimeOfStateCountReset        string `json:"timeOfStateCountReset"`
	ElapsedActiveTime            string `json:"elapsedActiveTime"`
	TimeOfActiveTimeReset        string `json:"timeOfActiveTimeReset"`
	MinimumOffTime               string `json:"minimumOffTime"`
	MinimumOnTime                string `json:"minimumOnTime"`
	PriorityArray                string `json:"priorityArray"`
	RelinquishDefault            string `json:"relinquishDefault"`
	TimeDelay                    string `json:"timeDelay"`
	NotificationClass            string `json:"notificationClass"`
	AlarmValue                   string `json:"alarmValue"`
	EventEnable                  string `json:"eventEnable"`
	AckedTransitions             string `json:"ackedTransitions"`
	NotifyType                   string `json:"notifyType"`
	EventTimeStamps              string `json:"eventTimeStamps"`
	EventMessageTexts            string `json:"eventMessageTexts"`
	ReliabilityEvaluationInhibit string `json:"reliabilityEvaluationInhibit"`
	ProfileName                  string `json:"profileName"`
}
