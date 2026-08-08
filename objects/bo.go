package objects

import "github.com/wz2b/bacnet/bactypes"

type BinaryOutput struct {
	bactypes.DefaultObject
	PresentValue                 string `json:"presentValue"`
	StatusFlags                  string `json:"statusFlags"`
	EventState                   string `json:"eventState"`
	OutOfService                 string `json:"outOfService"`
	Polarity                     string `json:"polarity"`
	PriorityArray                string `json:"priorityArray"`
	RelinquishDefault            string `json:"relinquishDefault"`
	Description                  string `json:"description"`
	DeviceType                   string `json:"deviceType"`
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
	TimeDelay                    string `json:"timeDelay"`
	NotificationClass            string `json:"notificationClass"`
	FeedbackValue                string `json:"feedbackValue"`
	EventEnable                  string `json:"eventEnable"`
	AckedTransitions             string `json:"ackedTransitions"`
	NotifyType                   string `json:"notifyType"`
	EventTimeStamps              string `json:"eventTimeStamps"`
	EventMessageTexts            string `json:"eventMessageTexts"`
	ReliabilityEvaluationInhibit string `json:"reliabilityEvaluationInhibit"`
	ProfileName                  string `json:"profileName"`
}
