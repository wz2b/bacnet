package objects

import "github.com/wz2b/bacnet/bactypes"

type Command struct {
	bactypes.DefaultObject
	PresentValue        string `json:"presentValue"`
	InProcess           string `json:"inProcess"`
	AllWritesSuccessful string `json:"allWritesSuccessful"`
	Action              string `json:"action"`
	Description         string `json:"description"`
	ActionText          string `json:"actionText"`
	ProfileName         string `json:"profileName"`
}
