package player

import "time"

type _AutoReady struct {
	Ts       time.Duration
	Deadline int64
}

func (p *Player) AutoReady() (time.Duration, int64) {
	val := p.extend["_auto_ready"]
	if val == nil {
		return 0, 0
	}

	autoReady, ok := val.(*_AutoReady)
	if !ok || autoReady == nil {
		return 0, 0
	}

	return autoReady.Ts, autoReady.Deadline
}

func (p *Player) SetAutoReady(ts time.Duration, deadline int64, ext map[string]any) {
	oldTs, oldDeadline := p.AutoReady()
	if oldTs == ts && oldDeadline == deadline {
		return
	}
	p.extend["_auto_ready"] = &_AutoReady{
		Ts:       ts,
		Deadline: deadline,
	}

	if f := p.watch["auto_ready"]; f != nil {
		f(p, ext)
	}
}

func (p *Player) WatchAutoReady(f func(*Player, map[string]any)) {
	p.watch["auto_ready"] = f
}
