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
	a.delayProxy.Add(ts, func(ts int64, args interface{}) {
		auto := args.(*_AutoDelay)
		switch auto.Op {
		case true:
			// 托管操作
			if a.cb != nil {
				if a.cb.OpPre != nil {
					a.cb.OpPre(auto.key)
				}
				if a.cb.Op != nil {
					a.cb.Op(auto.key)
				}
				if a.cb.opAfter != nil {
					a.cb.opAfter(auto.key)
				}
			}
		default:
			// 进入托管
			if a.cb != nil {
				if a.cb.AutoPre != nil {
					a.cb.AutoPre(auto.key)
				}
				if a.cb.Auto != nil {
					a.cb.Auto(auto.key)
				}
				if a.cb.AutoAfter != nil {
					a.cb.AutoAfter(auto.key)
				}
			}
			a.Add(key, true)
		}
	}, &_AutoDelay{
		key: key,
		Op:  op,
	})
}

func (a *Auto) AutoTs(key interface{}, op bool) time.Duration {
	var (
		ret  time.Duration
		auto *_AutoDelay
		ok   bool
	)
	a.delayProxy.Range(func(ts int64, args interface{}) bool {
		if auto, ok = args.(*_AutoDelay); ok && auto != nil && auto.key == key && auto.Op == op {
			ret = time.Duration(ts)
		}
		return ret == 0
	})
	return ret
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
