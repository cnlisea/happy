package delay

import (
	"time"
)

type Unit struct {
	DelayTime time.Duration
	CallTs    int64
	F         func(ts int64, arg interface{})
	Arg       interface{}
}
