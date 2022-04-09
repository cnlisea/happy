package delay

func (d *Delay) Handler() {
	var lastDelayTs int64
	d.QueueRangeDel(func(u *Unit) (bool, bool) {
		if lastDelayTs != 0 && u.CallTs-lastDelayTs > 1 {
			return false, false
		}

		if u.F != nil {
			u.F(u.CallTs, u.Arg)
		}
		lastDelayTs = u.CallTs
		return true, true
	})
}
