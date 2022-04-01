package delay

import (
	"container/list"
	"time"
)

type Delay struct {
	queue *list.List
	timer *time.Timer
}

func New() *Delay {
	t := time.NewTimer(0)
	t.Stop()
	return &Delay{
		queue: list.New(),
		timer: t,
	}
}
