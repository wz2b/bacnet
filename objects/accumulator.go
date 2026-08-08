package objects

import "github.com/wz2b/bacnet/bactypes"

type Accumulator struct {
	bactypes.DefaultObject
	PresentValue                 string `json:"presentValue"`
	StatusFlags                  string `json:"statusFlags"`
	EventState                   string `json:"eventState"`
	OutOfService                 string `json:"outOfService"`
	Scale                        string `json:"scale"`
	Units                        string `json:"units"`
	MaxPresValue                 string `json:"maxPresValue"`
	Description                  string `json:"description"`
	DeviceType                   string `json:"deviceType"`
	Reliability                  string `json:"reliability"`
	Prescale                     string `json:"prescale"`
	ValueChangeTime              string `json:"valueChangeTime"`
	ValueBeforeChange            string `json:"valueBeforeChange"`
	ValueSet                     string `json:"valueSet"`
	LoggingRecord                string `json:"loggingRecord"`
	LoggingObject                string `json:"loggingObject"`
	PulseRate                    string `json:"pulseRate"`
	HighLimit                    string `json:"highLimit"`
	LowLimit                     string `json:"lowLimit"`
	LimitMonitoringInterval      string `json:"limitMonitoringInterval"`
	NotificationClass            string `json:"notificationClass"`
	TimeDelay                    string `json:"timeDelay"`
	LimitEnable                  string `json:"limitEnable"`
	EventEnable                  string `json:"eventEnable"`
	AckedTransitions             string `json:"ackedTransitions"`
	NotifyType                   string `json:"notifyType"`
	EventTimeStamps              string `json:"eventTimeStamps"`
	EventMessageTexts            string `json:"eventMessageTexts"`
	ReliabilityEvaluationInhibit string `json:"reliabilityEvaluationInhibit"`
	ProfileName                  string `json:"profileName"`
}
