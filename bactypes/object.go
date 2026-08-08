package bactypes

import "github.com/wz2b/bacnet/defs"

type Object interface {
	ObjectID() ObjectID
	ObjectName() string
	ObjectType() defs.ObjectType
}

type DefaultObject struct {
	ObjectIdentifier uint32 `json:"objectIdentifier"`
	ObjectName       string `json:"objectName"`
	ObjectType       string `json:"objectType"`
}
