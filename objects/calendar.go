package objects

import "github.com/wz2b/bacnet/bactypes"

type Calendar struct {
	bactypes.DefaultObject
	PresentValue string `json:"presentValue"`
	DateList     string `json:"dateList"`
	Description  string `json:"description"`
	ProfileName  string `json:"profileName"`
}
