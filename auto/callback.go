package auto

import "time"

type Callback struct {
	// 准备
	ReadyAuto func(key interface{}, ts time.Duration, deadlineTs int64)
	ReadyOp   func(key interface{}, ts time.Duration, deadlineTs int64)

	// 托管进入
	Auto func(ts int64, key interface{})

	// 托管操作
	Op func(ts int64, key interface{})
}

func (a *Auto) Callback(cb *Callback) {
	a.cb = cb
}
