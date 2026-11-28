package player

import "time"

type _AutoOpReady struct {
	Ts       time.Duration
	Deadline int64
}

func (p *Player) AutoOpReady() (time.Duration, int64) {
	val := p.extend["_auto_op_ready"]
	if val == nil {
		return 0, 0
	}

	autoOpReady, ok := val.(*_AutoOpReady)
	if !ok || autoOpReady == nil {
		return 0, 0
	}

	return autoOpReady.Ts, autoOpReady.Deadline
}

func (p *Player) SetAutoOpReady(ts time.Duration, deadline int64, ext map[string]any) {
	oldTs, oldDeadlineTs := p.AutoOpReady()
	if oldTs == ts && oldDeadlineTs == deadline {
		return
	}
	p.extend["_auto_op_ready"] = &_AutoOpReady{
		Ts:       ts,
		Deadline: deadline,
	}

	if f := p.watch["auto_op_ready"]; f != nil {
		f(p, ext)
	}
}

func (p *Player) WatchAutoOpReady(f func(*Player, map[string]any)) {
	p.watch["auto_op_ready"] = f
}
