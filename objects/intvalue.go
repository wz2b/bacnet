package objects

import "github.com/wz2b/bacnet/bactypes"

type IntegerValue struct {
	bactypes.DefaultObject
	PresentValue                 string `json:"presentValue"`
	StatusFlags                  string `json:"statusFlags"`
	Units                        string `json:"units"`
	Description                  string `json:"description"`
	EventState                   string `json:"eventState"`
	Reliability                  string `json:"reliability"`
	OutOfService                 string `json:"outOfService"`
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
	EventMessageTextsConfig      string `json:"eventMessageTextsConfig"`
	EventDetectionEnable         string `json:"eventDetectionEnable"`
	EventAlgorithmInhibitRef     string `json:"eventAlgorithmInhibitRef"`
	EventAlgorithmInhibit        string `json:"eventAlgorithmInhibit"`
	TimeDelayNormal              string `json:"timeDelayNormal"`
	ReliabilityEvaluationInhibit string `json:"reliabilityEvaluationInhibit"`
	MinPresValue                 string `json:"minPresValue"`
	MaxPresValue                 string `json:"maxPresValue"`
	Resolution                   string `json:"resolution"`
	ProfileName                  string `json:"profileName"`
	All                          string `json:"all"`
	Required                     string `json:"required"`
	Optional                     string `json:"optional"`
}
