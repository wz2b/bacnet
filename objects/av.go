package objects

import "github.com/wz2b/bacnet/bactypes"

type AnalogValue struct {
	bactypes.DefaultObject
	PresentValue                 string `json:"presentValue"`
	StatusFlags                  string `json:"statusFlags"`
	EventState                   string `json:"eventState"`
	OutOfService                 string `json:"outOfService"`
	Units                        string `json:"units"`
	Description                  string `json:"description"`
	Reliability                  string `json:"reliability"`
	PriorityArray                string `json:"priorityArray"`
	RelinquishDefault            string `json:"relinquishDefault"`
	CovIncrement                 string `json:"covIncrement"`
	TimeDelay                    string `json:"timeDelay"`
	NotificationClass            string `json:"notificationClass"`
	HighLimit                    string `json:"highLimit"`
	LowLimit                     string `json:"lowLimit"`
	Deadband                     string `json:"deadband"`
	LimitEnable                  string `json:"limitEnable"`
	EventEnable                  string `json:"eventEnable"`
	AckedTransitions             string `json:"ackedTransitions"`
	NotifyType                   string `json:"notifyType"`
	EventTimeStamps              string `json:"eventTimeStamps"`
	EventMessageTexts            string `json:"eventMessageTexts"`
	ReliabilityEvaluationInhibit string `json:"reliabilityEvaluationInhibit"`
	ProfileName                  string `json:"profileName"`
}
