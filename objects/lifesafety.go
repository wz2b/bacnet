package objects

import "github.com/wz2b/bacnet/bactypes"

type LifeSafetyPoint struct {
	bactypes.DefaultObject
	PresentValue                 string `json:"presentValue"`
	TrackingValue                string `json:"trackingValue"`
	StatusFlags                  string `json:"statusFlags"`
	EventState                   string `json:"eventState"`
	OutOfService                 string `json:"outOfService"`
	Reliability                  string `json:"reliability"`
	Mode                         string `json:"mode"`
	AcceptedModes                string `json:"acceptedModes"`
	Silenced                     string `json:"silenced"`
	OperationExpected            string `json:"operationExpected"`
	Description                  string `json:"description"`
	DeviceType                   string `json:"deviceType"`
	NotificationClass            string `json:"notificationClass"`
	LifeSafetyAlarmValues        string `json:"lifeSafetyAlarmValues"`
	AlarmValues                  string `json:"alarmValues"`
	FaultValues                  string `json:"faultValues"`
	EventEnable                  string `json:"eventEnable"`
	AckedTransitions             string `json:"ackedTransitions"`
	NotifyType                   string `json:"notifyType"`
	EventTimeStamps              string `json:"eventTimeStamps"`
	EventMessageTexts            string `json:"eventMessageTexts"`
	MaintenanceRequired          string `json:"maintenanceRequired"`
	Setting                      string `json:"setting"`
	DirectReading                string `json:"directReading"`
	Units                        string `json:"units"`
	MemberOf                     string `json:"memberOf"`
	ReliabilityEvaluationInhibit string `json:"reliabilityEvaluationInhibit"`
	ProfileName                  string `json:"profileName"`
}
