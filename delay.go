package happy

import "time"

func (h *Happy) DelayFunc(delayTs time.Duration, f func(args interface{}), args interface{}) {
	h.delay.Add(delayTs, func(ts int64, args interface{}) {
		f(args)
	}, args)
}
