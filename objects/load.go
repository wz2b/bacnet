package objects

import "github.com/wz2b/bacnet/bactypes"

type LoadControl struct {
	bactypes.DefaultObject
	PresentValue                 string `json:"presentValue"`
	StatusFlags                  string `json:"statusFlags"`
	EventState                   string `json:"eventState"`
	RequestedShedLevel           string `json:"requestedShedLevel"`
	StartTime                    string `json:"startTime"`
	ShedDuration                 string `json:"shedDuration"`
	DutyWindow                   string `json:"dutyWindow"`
	Enable                       string `json:"enable"`
	ExpectedShedLevel            string `json:"expectedShedLevel"`
	ActualShedLevel              string `json:"actualShedLevel"`
	ShedLevels                   string `json:"shedLevels"`
	ShedLevelDescriptions        string `json:"shedLevelDescriptions"`
	Description                  string `json:"description"`
	StateDescription             string `json:"stateDescription"`
	Reliability                  string `json:"reliability"`
	FullDutyBaseline             string `json:"fullDutyBaseline"`
	NotificationClass            string `json:"notificationClass"`
	TimeDelay                    string `json:"timeDelay"`
	EventEnable                  string `json:"eventEnable"`
	AckedTransitions             string `json:"ackedTransitions"`
	NotifyType                   string `json:"notifyType"`
	EventTimeStamps              string `json:"eventTimeStamps"`
	EventMessageTexts            string `json:"eventMessageTexts"`
	ReliabilityEvaluationInhibit string `json:"reliabilityEvaluationInhibit"`
	ProfileName                  string `json:"profileName"`
}
