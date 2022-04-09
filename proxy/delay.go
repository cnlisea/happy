package proxy

import "time"

type Delay interface {
	Add(delayTime time.Duration, f func(ts int64, args interface{}), arg interface{})
	Range(f func(ts int64, args interface{}) bool)
	Del(f func(ts int64, args interface{}) bool)
}
