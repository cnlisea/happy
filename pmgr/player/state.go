package player

import "time"

func (p *Player) Line() bool {
	return p.offlineTs == 0
}

func (p *Player) OfflineTs() int64 {
	return p.offlineTs
}

func (p *Player) Offline(ts int64) {
	if ts <= 0 {
		ts = time.Now().Unix()
	}
	p.offlineTs = ts
	if f := p.watch["line"]; f != nil {
		f(p)
	}
}

func (p *Player) Online() {
	p.offlineTs = 0
	if f := p.watch["line"]; f != nil {
		f(p)
	}
}

func (p *Player) WatchLine(f func(*Player)) {
	p.watch["line"] = f
}

func (p *Player) Ready() bool {
	return p.state&1 == 1
}

func (p *Player) SetReady(b bool) {
	if p.Ready() == b {
		return
	}

	if b {
		p.state |= 1
	} else {
		p.state ^= 1
	}

	if f := p.watch["ready"]; f != nil {
		f(p)
	}
}

func (p *Player) WatchReady(f func(*Player)) {
	p.watch["ready"] = f
}

func (p *Player) Op() bool {
	return p.state&2 == 2
}

func (p *Player) SetOp(b bool) {
	if p.Op() == b {
		return
	}

	if b {
		p.state |= 2
	} else {
		p.state ^= 2
	}

	if f := p.watch["op"]; f != nil {
		f(p)
	}
}

func (p *Player) WatchOp(f func(*Player)) {
	p.watch["op"] = f
}

func (p *Player) View() bool {
	return p.state&4 == 4
}

func (p *Player) SetView(b bool) {
	if p.View() == b {
		return
	}

	if b {
		p.state |= 4
	} else {
		p.state ^= 4
	}
	if f := p.watch["view"]; f != nil {
		f(p)
	}
}

func (p *Player) WatchView(f func(*Player)) {
	p.watch["view"] = f
}

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
