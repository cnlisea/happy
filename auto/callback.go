package auto

type Callback struct {
	// 托管进入
	AutoPre   func(key interface{})
	Auto      func(key interface{})
	AutoAfter func(key interface{})

	// 托管操作
	OpPre   func(key interface{})
	Op      func(key interface{})
	opAfter func(key interface{})
}

func (a *Auto) Callback(cb *Callback) {
	a.cb = cb
}
