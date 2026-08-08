package bactypes

import "github.com/wz2b/bacnet/defs"

type ObjectID struct {
	Type     defs.ObjectType
	Instance uint32
}
