package player

func (p *Player) Site() uint32 {
	return p.site
}

func (p *Player) SetSite(site uint32, ext map[string]any) {
	p.site = site
	if f := p.watch["site"]; f != nil {
		f(p, ext)
	}
}

func (p *Player) WatchSite(f func(*Player, map[string]any)) {
	p.watch["site"] = f
}
