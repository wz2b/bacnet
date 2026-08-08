package objects

import "github.com/wz2b/bacnet/bactypes"

type TrendLog struct {
	bactypes.DefaultObject
	Enable                       string `json:"enable"`
	StopWhenFull                 string `json:"stopWhenFull"`
	BufferSize                   string `json:"bufferSize"`
	LogBuffer                    string `json:"logBuffer"`
	RecordCount                  string `json:"recordCount"`
	TotalRecordCount             string `json:"totalRecordCount"`
	EventState                   string `json:"eventState"`
	LoggingType                  string `json:"loggingType"`
	StatusFlags                  string `json:"statusFlags"`
	Description                  string `json:"description"`
	StartTime                    string `json:"startTime"`
	StopTime                     string `json:"stopTime"`
	LogDeviceObjectProperty      string `json:"logDeviceObjectProperty"`
	LogInterval                  string `json:"logInterval"`
	CovResubscriptionInterval    string `json:"covResubscriptionInterval"`
	ClientCovIncrement           string `json:"clientCovIncrement"`
	NotificationThreshold        string `json:"notificationThreshold"`
	RecordsSinceNotification     string `json:"recordsSinceNotification"`
	LastNotifyRecord             string `json:"lastNotifyRecord"`
	NotificationClass            string `json:"notificationClass"`
	EventEnable                  string `json:"eventEnable"`
	AckedTransitions             string `json:"ackedTransitions"`
	NotifyType                   string `json:"notifyType"`
	EventTimeStamps              string `json:"eventTimeStamps"`
	EventMessageTexts            string `json:"eventMessageTexts"`
	AlignIntervals               string `json:"alignIntervals"`
	IntervalOffset               string `json:"intervalOffset"`
	Trigger                      string `json:"trigger"`
	Reliability                  string `json:"reliability"`
	ReliabilityEvaluationInhibit string `json:"reliabilityEvaluationInhibit"`
	ProfileName                  string `json:"profileName"`
}
