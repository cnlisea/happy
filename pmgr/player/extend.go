package player

func (p *Player) Extend(key string) interface{} {
	return p.extend[key]
}

func (p *Player) SetExtend(key string, val interface{}) {
	p.extend[key] = val
}
