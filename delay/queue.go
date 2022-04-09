package delay

import (
	"container/list"
	"time"
)

func (d *Delay) QueueLen() int {
	if d.queue == nil {
		return 0
	}
	return d.queue.Len()
}

func (d *Delay) QueueAdd(unit *Unit) {
	if unit == nil {
		return
	}

	if d.queue == nil {
		d.queue = list.New()
	}

	if d.queue.Len() == 0 {
		d.queue.PushBack(unit)
		d.timer.Reset(unit.DelayTime)
		return
	}

	var (
		front = d.queue.Front()
		e     = front
		u     *Unit
	)
	for ; e != nil; e = e.Next() {
		u = e.Value.(*Unit)
		if u.CallTs > unit.CallTs {
			break
		}
	}

	switch e {
	case nil:
		d.queue.PushBack(unit)
	case front:
		d.queue.PushFront(unit)
		d.timer.Reset(unit.DelayTime)
	default:
		d.queue.InsertBefore(unit, e)
	}
}

func (d *Delay) QueueRangeDel(f func(unit *Unit) (bool, bool)) {
	if f == nil || d.QueueLen() == 0 {
		return
	}

	var (
		front    = d.queue.Front()
		e        = front
		nextE    *list.Element
		unit     *Unit
		con, del bool
		frontDel bool
	)
	for e != nil {
		nextE = e.Next()
		unit = e.Value.(*Unit)
		con, del = f(unit)
		if del {
			if e == front {
				frontDel = true
			}
			d.queue.Remove(e)
		}
		if !con {
			break
		}
		e = nextE
	}

	if frontDel {
		d.timer.Stop()
		d.QueueRange(func(u *Unit) bool {
			d.timer.Reset(time.Duration(u.CallTs - time.Now().UnixNano()))
			return false
		})
	}
}

func (d *Delay) QueueRange(f func(unit *Unit) bool) {
	if f == nil || d.QueueLen() == 0 {
		return
	}

	var (
		e    *list.Element
		unit *Unit
	)
	for e = d.queue.Front(); e != nil; e = e.Next() {
		if unit = e.Value.(*Unit); !f(unit) {
			break
		}
	}
}
