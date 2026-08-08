package bactypes

import "github.com/wz2b/bacnet"

type TimeValue struct {
	Time  Time
	Value bacnet.BACNET_APPLICATION_DATA_VALUE
}
