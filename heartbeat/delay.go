package heartbeat

import (
	"time"
)

func (h *Heartbeat) DelayAdd(delay time.Duration) {
	h.checkTs = h.delayProxy.Add(delay, h.Handler, h)
}

func (h *Heartbeat) DelayDel() {
	h.delayProxy.Del(func(ts int64, args interface{}) bool {
		hb, ok := args.(*Heartbeat)
		return ok && hb == h
	})
}
