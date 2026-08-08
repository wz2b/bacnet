package objects

import "github.com/wz2b/bacnet/bactypes"

type LightingOutput struct {
	bactypes.DefaultObject
	PresentValue                   string `json:"presentValue"`
	TrackingValue                  string `json:"trackingValue"`
	LightingCommand                string `json:"lightingCommand"`
	InProgress                     string `json:"inProgress"`
	StatusFlags                    string `json:"statusFlags"`
	OutOfService                   string `json:"outOfService"`
	BlinkWarnEnable                string `json:"blinkWarnEnable"`
	EgressTime                     string `json:"egressTime"`
	EgressActive                   string `json:"egressActive"`
	DefaultFadeTime                string `json:"defaultFadeTime"`
	DefaultRampRate                string `json:"defaultRampRate"`
	DefaultStepIncrement           string `json:"defaultStepIncrement"`
	PriorityArray                  string `json:"priorityArray"`
	RelinquishDefault              string `json:"relinquishDefault"`
	LightingCommandDefaultPriority string `json:"lightingCommandDefaultPriority"`
	Description                    string `json:"description"`
	Reliability                    string `json:"reliability"`
	Transition                     string `json:"transition"`
	FeedbackValue                  string `json:"feedbackValue"`
	Power                          string `json:"power"`
	InstantaneousPower             string `json:"instantaneousPower"`
	MinActualValue                 string `json:"minActualValue"`
	MaxActualValue                 string `json:"maxActualValue"`
	CovIncrement                   string `json:"covIncrement"`
	ReliabilityEvaluationInhibit   string `json:"reliabilityEvaluationInhibit"`
	ProfileName                    string `json:"profileName"`
}
