package objects

import "github.com/wz2b/bacnet/bactypes"

type Averaging struct {
	bactypes.DefaultObject
	MinimumValue            string `json:"minimumValue"`
	AverageValue            string `json:"averageValue"`
	MaximumValue            string `json:"maximumValue"`
	StatusFlags             string `json:"statusFlags"`
	EventState              string `json:"eventState"`
	OutOfService            string `json:"outOfService"`
	Units                   string `json:"units"`
	ProfileName             string `json:"profileName"`
	MinimumValueTimestamp   string `json:"minimumValueTimestamp"`
	VarianceValue           string `json:"varianceValue"`
	MaximumValueTimestamp   string `json:"maximumValueTimestamp"`
	Description             string `json:"description"`
	AttemptedSamples        string `json:"attemptedSamples"`
	ValidSamples            string `json:"validSamples"`
	ObjectPropertyReference string `json:"objectPropertyReference"`
	WindowInterval          string `json:"windowInterval"`
	WindowSamples           string `json:"windowSamples"`
}
