package heartbeat

import "time"

func (h *Heartbeat) Active(ts int64) {
	if ts == 0 {
		ts = time.Now().Unix()
	}
	h.lastActiveTs = ts
}
