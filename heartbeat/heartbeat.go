package heartbeat

import (
	"time"

	"github.com/cnlisea/happy/proxy"
)

type Heartbeat struct {
	delayProxy proxy.Delay

	interval     time.Duration
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
	h.delayProxy.Add(h.interval, h.Handler, nil)

	return h
}

func (h *Heartbeat) Handler(ts int64, args interface{}) {
	if ts-h.lastActiveTs < int64(h.interval) {
		h.delayProxy.Add(h.interval, h.Handler, nil)
		return
	}

	for i := range h.fn {
		if h.fn[i] != nil {
			h.fn[i]()
		}
	}
}
