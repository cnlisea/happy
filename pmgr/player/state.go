package player

import "time"

func (p *Player) Offline(ts int64) {
	if ts <= 0 {
		ts = time.Now().Unix()
	}
	p.offlineTs = ts
}

func (p *Player) Online() {
	p.offlineTs = 0
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
}
