package player

func (p *Player) Extend(key string) interface{} {
	return p.extend[key]
}

func (p *Player) SetExtend(key string, val interface{}, ext map[string]any) {
	p.extend[key] = val
	if f := p.watch["extend"]; f != nil {
		f(p, ext)
	}
}

func (p *Player) WatchExtend(f func(*Player, map[string]any)) {
	p.watch["extend"] = f
}
