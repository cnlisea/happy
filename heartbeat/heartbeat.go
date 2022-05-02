package heartbeat

import (
	"time"

	"github.com/cnlisea/happy/proxy"
)

type Heartbeat struct {
	delayProxy proxy.Delay

	interval     time.Duration
	checkTs      int64
	lastActiveTs int64
	fn           []func()
}

func New(delay proxy.Delay, interval time.Duration, fn ...func()) *Heartbeat {
	if interval == 0 {
		interval = time.Second
	}
	h := &Heartbeat{
		delayProxy: delay,
		interval:   interval,
		fn:         fn,
	}

	h.Active(0)
	h.DelayAdd(h.interval)

	return h
}

func (h *Heartbeat) Interval(interval time.Duration) {
	if interval == 0 {
		interval = time.Second
	}

	h.interval = interval
	h.DelayDel()
	h.DelayAdd(interval)
}

func (h *Heartbeat) Handler(ts int64, args interface{}) {
	if ts-h.lastActiveTs < int64(h.interval) {
		h.DelayAdd(time.Duration(h.lastActiveTs) + h.interval - time.Duration(ts))
		return
	}

	for i := range h.fn {
		if h.fn[i] != nil {
			h.fn[i]()
		}
	}
}
