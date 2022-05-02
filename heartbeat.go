package happy

import (
	"errors"
	"github.com/cnlisea/happy/heartbeat"
	"time"
)

func (h *_Happy) Heartbeat(interval time.Duration) error {
	if interval <= 0 {
		return errors.New("interval invalid")
	}

	switch h.heartbeat {
	case nil:
		h.heartbeat = heartbeat.New(h.delay, interval, func() {
			// 超时解散
			h.Finish(false)
		})
	default:
		h.heartbeat.Interval(interval)
	}
	return nil
}
