package bactypes

type ObjectType uint16

type Object interface {
	ObjectID() ObjectID
	ObjectName() string
	ObjectType() ObjectType
}

type DefaultObject struct {
	ObjectIdentifier uint32 `json:"objectIdentifier"`
	ObjectName       string `json:"objectName"`
	ObjectType       string `json:"objectType"`
}
