package player

import "time"

func (p *Player) Auto() bool {
	return p.state&8 == 8
}

func (p *Player) SetAuto(b bool, ext map[string]any) {
	if p.Auto() == b {
		return
	}

	if b {
		p.state |= 8
	} else {
		p.state ^= 8
	}

	if f := p.watch["auto"]; f != nil {
		f(p, ext)
	}
}

func (p *Player) WatchAuto(f func(*Player, map[string]any)) {
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

func (p *Player) SetAutoTs(ts time.Duration, ext map[string]any) {
	if p.AutoTs() == ts {
		return
	}
	p.extend["_auto_ts"] = ts

	if f := p.watch["auto_ts"]; f != nil {
		f(p, ext)
	}
}

func (p *Player) WatchAutoTs(f func(*Player, map[string]any)) {
	p.watch["auto_ts"] = f
}
