package delay

import (
	"container/list"
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

func (d *Delay) QueueRange(f func(unit *Unit) (bool, bool)) {
	if f == nil || d.QueueLen() == 0 {
		return
	}

	var (
		e        = d.queue.Front()
		nextE    *list.Element
		unit     *Unit
		con, del bool
	)
	for e != nil {
		nextE = e.Next()
		unit = e.Value.(*Unit)
		con, del = f(unit)
		if del {
			d.queue.Remove(e)
		}
		if !con {
			break
		}
		e = nextE
	}
}
