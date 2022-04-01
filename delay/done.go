package delay

import "time"

func (d *Delay) Done() <-chan time.Time {
	return d.timer.C
}
