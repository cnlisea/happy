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

func (d *Delay) Unit(delayTime time.Duration, f func(ts int64, args interface{}), arg interface{}) {
	if f == nil {
		return
	}

	d.QueueAdd(&Unit{
		DelayTime: delayTime,
		CallTs:    time.Now().Add(delayTime).UnixNano(),
		F:         f,
		Arg:       arg,
	})
}
