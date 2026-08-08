package objects

import "github.com/wz2b/bacnet/bactypes"

type AnalogOutput struct {
	bactypes.DefaultObject
	PresentValue                 string `json:"presentValue"`
	StatusFlags                  string `json:"statusFlags"`
	EventState                   string `json:"eventState"`
	OutOfService                 string `json:"outOfService"`
	Units                        string `json:"units"`
	PriorityArray                string `json:"priorityArray"`
	RelinquishDefault            string `json:"relinquishDefault"`
	Description                  string `json:"description"`
	DeviceType                   string `json:"deviceType"`
	Reliability                  string `json:"reliability"`
	MinPresValue                 string `json:"minPresValue"`
	MaxPresValue                 string `json:"maxPresValue"`
	Resolution                   string `json:"resolution"`
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
