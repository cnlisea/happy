package vote

import (
	"time"

	"github.com/cnlisea/happy/proxy"
)

type Vote struct {
	agree   []interface{}
	reject  []interface{}
	full    bool
	fullEnd bool

	passFn   []func(ts int64)
	failFn   []func(ts int64)
	cancelFn []func(ts int64)
	addFn    []func(ts time.Duration, deadlineTs int64, key interface{}, agree bool)

	delay         proxy.Delay
	deadline      bool
	deadlineTs    time.Duration
	deadlineEndTs int64
	deadlinePass  bool
	deadlineFirst bool
	deadlineDelay bool
}

func New(min, max int) *Vote {
	return &Vote{
		agree:  make([]interface{}, 0, min),
		reject: make([]interface{}, 0, max-min+1),
	}
}

func (v *Vote) End() bool {
	return v.full
}

func (v *Vote) Cancel() {
	if v.full {
		return
	}

	v.full = true
	ts := time.Now().UnixNano()
	for i := range v.cancelFn {
		v.cancelFn[i](ts)
	}
}

func (v *Vote) FullEnd() {
	v.fullEnd = true
}

func (v *Vote) CallbackPass(f ...func(ts int64)) {
	v.passFn = f
}

func (v *Vote) CallbackFail(f ...func(ts int64)) {
	v.failFn = f
}

func (v *Vote) CallbackCancel(f ...func(ts int64)) {
	v.cancelFn = f
}

func (v *Vote) CallbackAdd(f ...func(ts time.Duration, deadlineTs int64, key interface{}, agree bool)) {
	v.addFn = f
}

func (v *Vote) Deadline(delay proxy.Delay, ts time.Duration, pass bool, first bool) {
	v.deadline = true
	v.delay = delay
	v.deadlineTs = ts
	v.deadlinePass = pass
	v.deadlineFirst = first
	if !v.deadlineFirst {
		v.deadlineDelayDel()
		v.deadlineDelayAdd()
	}
}

func (v *Vote) deadlineDelayAdd() {
	if v.delay == nil || v.deadlineDelay {
		return
	}

	v.deadlineEndTs = v.delay.Add(v.deadlineTs, func(ts int64, args interface{}) {
		vote := args.(*Vote)
		vote.Full(true, ts)
	}, v)
	v.deadlineDelay = true
}

func (v *Vote) deadlineDelayDel() {
	if v.delay == nil || !v.deadlineDelay {
		return
	}

	var (
		vote *Vote
		ok   bool
	)
	v.delay.Del(func(ts int64, args interface{}) bool {
		vote, ok = args.(*Vote)
		return ok && vote == v
	})
	v.deadlineEndTs = 0
	v.deadlineDelay = false
}

func (v *Vote) Add(key interface{}, agree bool) {
	if v.End() {
		return
	}

	if v.Exist(key) {
		return
	}

	if v.deadline && v.deadlineFirst && v.Num() == 0 {
		v.deadlineDelayAdd()
	}

	if agree {
		v.agree = append(v.agree, key)
	} else {
		v.reject = append(v.reject, key)
	}

	for i := range v.addFn {
		v.addFn[i](v.deadlineTs, v.deadlineEndTs, key, agree)
	}

	v.Full(false, time.Now().UnixNano())
}

func (v *Vote) Full(deadline bool, ts int64) {
	if v.full {
		return
	}

	var (
		agreeLen      = len(v.agree)
		fullAgreeLen  = cap(v.agree)
		rejectLen     = len(v.reject)
		fullRejectLen = cap(v.reject)
	)
	if !deadline {
		if v.fullEnd && agreeLen+rejectLen < fullAgreeLen+fullRejectLen-1 {
			return
		}

		if !v.fullEnd && agreeLen < fullAgreeLen && rejectLen < fullRejectLen {
			return
		}
	}

	v.full = true
	if ts <= 0 {
		ts = time.Now().UnixNano()
	}
	switch {
	case agreeLen == fullAgreeLen, deadline && v.deadlinePass:
		for i := range v.passFn {
			v.passFn[i](ts)
		}
	default:
		for i := range v.failFn {
			v.failFn[i](ts)
		}
	}
}

func (v *Vote) Num() int {
	return len(v.agree) + len(v.reject)
}

func (v *Vote) Exist(key interface{}, op ...bool) bool {
	var diffOp *bool
	if len(op) > 0 {
		diffOp = &op[0]
	}

	var exist bool
	v.Range(func(k interface{}, o bool) bool {
		if k == key && (diffOp == nil || *diffOp == o) {
			exist = true
		}
		return !exist
	})
	return exist
}

func (v *Vote) Range(f func(key interface{}, op bool) bool) {
	var i int
	for i = range v.agree {
		if !f(v.agree[i], true) {
			return
		}
	}

	for i = range v.reject {
		if !f(v.reject[i], false) {
			return
		}
	}
}

func (v *Vote) Reset() {
	v.agree = v.agree[:0]
	v.reject = v.reject[:0]
	v.full = false
	v.deadlineDelayDel()
	if !v.deadlineFirst {
		v.deadlineDelayAdd()
	}
}
