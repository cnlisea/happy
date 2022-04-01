package delay

import (
	"time"
)

func (d *Delay) Handler() {
	var lastDelayTs int64
	d.QueueRange(func(u *Unit) (bool, bool) {
		if lastDelayTs != 0 && u.CallTs-lastDelayTs > 1 {
			d.timer.Reset(time.Duration(u.CallTs - lastDelayTs))
			return false, false
		}

		if u.F != nil {
			u.F(u.CallTs, u.Arg)
		}
		lastDelayTs = u.CallTs
		return true, true
	})
}
