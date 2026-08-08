package objects

import "github.com/wz2b/bacnet/bactypes"

type File struct {
	bactypes.DefaultObject
	FileType         string `json:"fileType"`
	FileSize         string `json:"fileSize"`
	ModificationDate string `json:"modificationDate"`
	Archive          string `json:"archive"`
	ReadOnly         string `json:"readOnly"`
	FileAccessMethod string `json:"fileAccessMethod"`
	Description      string `json:"description"`
	RecordCount      string `json:"recordCount"`
	ProfileName      string `json:"profileName"`
}
