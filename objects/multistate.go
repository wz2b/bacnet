package objects

import "github.com/wz2b/bacnet/bactypes"

type MultistateInput struct {
	bactypes.DefaultObject
	PresentValue                 string `json:"presentValue"`
	StatusFlags                  string `json:"statusFlags"`
	EventState                   string `json:"eventState"`
	OutOfService                 string `json:"outOfService"`
	NumberOfStates               string `json:"numberOfStates"`
	Description                  string `json:"description"`
	DeviceType                   string `json:"deviceType"`
	Reliability                  string `json:"reliability"`
	StateText                    string `json:"stateText"`
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

type MultistateOutput struct {
	bactypes.DefaultObject
	PresentValue                 string `json:"presentValue"`
	StatusFlags                  string `json:"statusFlags"`
	EventState                   string `json:"eventState"`
	OutOfService                 string `json:"outOfService"`
	NumberOfStates               string `json:"numberOfStates"`
	PriorityArray                string `json:"priorityArray"`
	RelinquishDefault            string `json:"relinquishDefault"`
	Description                  string `json:"description"`
	DeviceType                   string `json:"deviceType"`
	Reliability                  string `json:"reliability"`
	StateText                    string `json:"stateText"`
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

type MultistateValue struct {
	bactypes.DefaultObject
	PresentValue                 string `json:"presentValue"`
	StatusFlags                  string `json:"statusFlags"`
	EventState                   string `json:"eventState"`
	OutOfService                 string `json:"outOfService"`
	NumberOfStates               string `json:"numberOfStates"`
	Description                  string `json:"description"`
	Reliability                  string `json:"reliability"`
	StateText                    string `json:"stateText"`
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
