package happy

import (
	"errors"
	"github.com/cnlisea/happy/heartbeat"
	"time"
)

func (h *Happy) Heartbeat(interval time.Duration) error {
	if interval <= 0 {
		return errors.New("interval invalid")
	}

	if h.heartbeat != nil {
		return errors.New("heartbeat already existed")
	}

	h.heartbeat = heartbeat.New(h.delay, interval, func() {
		// 超时解散
		h.Finish(false)
	})
	return nil
}
