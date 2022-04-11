package player

func (p *Player) Extend(key string) interface{} {
	return p.extend[key]
}

func (p *Player) SetExtend(key string, val interface{}) {
	p.extend[key] = val
	if f := p.watch["extend"]; f != nil {
		f(p)
	}
}

func (p *Player) WatchExtend(f func(*Player)) {
	p.watch["extend"] = f
}
