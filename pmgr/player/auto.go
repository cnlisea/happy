package player

import "time"

func (p *Player) Auto() bool {
	return p.state&8 == 8
}

func (p *Player) SetAuto(b bool) {
	if p.Auto() == b {
		return
	}

	if b {
		p.state |= 8
	} else {
		p.state ^= 8
	}

	if f := p.watch["auto"]; f != nil {
		f(p)
	}
}

func (p *Player) WatchAuto(f func(*Player)) {
	p.watch["auto"] = f
}

func (p *Player) AutoTs() time.Duration {
	val := p.extend["_auto_ts"]
	if val == nil {
		return 0
	}

	ts, ok := val.(time.Duration)
	if !ok {
		return 0
	}

	return ts
}

func (p *Player) SetAutoTs(ts time.Duration) {
	if p.AutoTs() == ts {
		return
	}
	p.extend["_auto_ts"] = ts

	if f := p.watch["auto_ts"]; f != nil {
		f(p)
	}
}

func (p *Player) WatchAutoTs(f func(*Player)) {
	p.watch["auto_ts"] = f
}
