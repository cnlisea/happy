package player

func (p *Player) Site() uint32 {
	return p.site
}

func (p *Player) SetSite(site uint32) {
	p.site = site
	if f := p.watch["site"]; f != nil {
		f(p)
	}
}

func (p *Player) WatchSite(f func(*Player)) {
	p.watch["site"] = f
}
