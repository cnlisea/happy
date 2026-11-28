package auto

import (
	"github.com/cnlisea/happy/proxy"
	"time"
)

type Auto struct {
	delayProxy proxy.Delay

	delayFn func(op bool) time.Duration
	cb      *Callback
}

func New(delayProxy proxy.Delay, delayFn func(op bool) time.Duration) *Auto {
	return &Auto{
		delayProxy: delayProxy,
		delayFn:    delayFn,
	}
}

type _AutoDelay struct {
	key interface{}
	Op  bool
}

func (a *Auto) Add(key interface{}, op bool) {
	ts := a.delayFn(op)
	if ts <= 0 {
		return
	}

	a.Del(key)
	deadline := a.delayProxy.Add(ts, func(ts int64, args interface{}) {
		auto := args.(*_AutoDelay)
		switch auto.Op {
		case true:
			// 托管操作
			if a.cb != nil {
				if a.cb.Op != nil {
					a.cb.Op(ts, auto.key)
				}
			}
		default:
			// 进入托管
			if a.cb != nil {
				if a.cb.Auto != nil {
					a.cb.Auto(ts, auto.key)
				}
			}
			var exist bool
			a.Range(func(auto *_AutoDelay) bool {
				if auto.key == key && auto.Op {
					exist = true
				}
				return !exist
			})
			if !exist {
				a.Add(key, true)
			}
		}
	}, &_AutoDelay{
		key: key,
		Op:  op,
	})
	if op {
		if a.cb != nil && a.cb.ReadyOp != nil {
			a.cb.ReadyOp(key, ts, deadline)
		}
	} else {
		if a.cb != nil && a.cb.ReadyAuto != nil {
			a.cb.ReadyAuto(key, ts, deadline)
		}
	}
}

func (a *Auto) AutoTs(key interface{}) (bool, int64) {
	var (
		retOp bool
		retTs int64
		auto  *_AutoDelay
		ok    bool
	)
	a.delayProxy.Range(func(ts int64, args interface{}) bool {
		if auto, ok = args.(*_AutoDelay); ok && auto != nil && auto.key == key {
			retOp = auto.Op
			retTs = ts
		}
		return retTs == 0
	})
	return retOp, retTs
}

func (a *Auto) Range(f func(auto *_AutoDelay) bool) {
	var (
		auto *_AutoDelay
		ok   bool
	)
	a.delayProxy.Range(func(ts int64, args interface{}) bool {
		auto, ok = args.(*_AutoDelay)
		return !ok || f(auto)
	})
}

func (a *Auto) Del(key interface{}) {
	var (
		auto *_AutoDelay
		ok   bool
	)
	a.delayProxy.Del(func(ts int64, arg interface{}) bool {
		auto, ok = arg.(*_AutoDelay)
		if !ok {
			return false
		}

		return auto == nil || auto.key == key
	})
}

func (a *Auto) Clean() {
	var ok bool
	a.delayProxy.Del(func(ts int64, arg interface{}) bool {
		_, ok = arg.(*_AutoDelay)
		return ok
	})
}
