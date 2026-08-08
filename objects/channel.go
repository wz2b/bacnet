package objects

import "github.com/wz2b/bacnet/bactypes"

type Channel struct {
	bactypes.DefaultObject
	PresentValue                   string `json:"presentValue"`
	LastPriority                   string `json:"lastPriority"`
	WriteStatus                    string `json:"writeStatus"`
	StatusFlags                    string `json:"statusFlags"`
	OutOfService                   string `json:"outOfService"`
	ListOfObjectPropertyReferences string `json:"listOfObjectPropertyReferences"`
	ChannelNumber                  string `json:"channelNumber"`
	ControlGroups                  string `json:"controlGroups"`
	Description                    string `json:"description"`
	Reliability                    string `json:"reliability"`
	ExecutionDelay                 string `json:"executionDelay"`
	AllowGroupDelayInhibit         string `json:"allowGroupDelayInhibit"`
	EventDetectionEnable           string `json:"eventDetectionEnable"`
	NotificationClass              string `json:"notificationClass"`
	EventEnable                    string `json:"eventEnable"`
	EventState                     string `json:"eventState"`
	AckedTransitions               string `json:"ackedTransitions"`
	NotifyType                     string `json:"notifyType"`
	EventTimeStamps                string `json:"eventTimeStamps"`
	EventMessageTexts              string `json:"eventMessageTexts"`
	EventMessageTextsConfig        string `json:"eventMessageTextsConfig"`
	ReliabilityEvaluationInhibit   string `json:"reliabilityEvaluationInhibit"`
	ProfileName                    string `json:"profileName"`
}
