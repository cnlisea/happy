package delay

import (
	"container/list"
	"time"
)

type Delay struct {
	queue *list.List
	timer *time.Timer
}

func New() *Delay {
	t := time.NewTimer(0)
	t.Stop()
	return &Delay{
		queue: list.New(),
		timer: t,
	}
}

func (d *Delay) Add(delayTime time.Duration, f func(ts int64, args interface{}), arg interface{}) int64 {
	if f == nil {
		return 0
	}

	u := &Unit{
		DelayTime: delayTime,
		CallTs:    time.Now().Add(delayTime).UnixNano(),
		F:         f,
		Arg:       arg,
	}
	d.QueueAdd(u)
	return u.CallTs
}

func (d *Delay) Range(f func(ts int64, args interface{}) bool) {
	d.QueueRange(func(u *Unit) bool {
		return f(u.CallTs, u.Arg)
	})
}

func (d *Delay) Del(f func(ts int64, args interface{}) bool) {
	d.QueueRangeDel(func(u *Unit) (bool, bool) {
		return true, f(u.CallTs, u.Arg)
	})
}
