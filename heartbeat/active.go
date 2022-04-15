package heartbeat

import (
	"time"
)

func (h *Heartbeat) Active(ts int64) {
	if ts == 0 {
		ts = time.Now().UnixNano()
	}

	if h.lastActiveTs > 0 && ts+int64(h.interval)-h.checkTs > int64(time.Second) {
		h.DelayDel()
		h.DelayAdd(h.interval)
	}
	h.lastActiveTs = ts
}
