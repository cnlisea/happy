package player

func (p *Player) AutoOp() int64 {
	val := p.extend["_auto_op"]
	if val == nil {
		return 0
	}

	ts, ok := val.(int64)
	if !ok {
		return 0
	}

	return ts
}

func (p *Player) SetAutoOp(ts int64, ext map[string]any) {
	if p.AutoOp() == ts {
		return
	}
	p.extend["_auto_op"] = ts

	if f := p.watch["auto_op"]; f != nil {
		f(p, ext)
	}
}

func (p *Player) WatchAutoOp(f func(*Player, map[string]any)) {
	p.watch["auto_op"] = f
}
