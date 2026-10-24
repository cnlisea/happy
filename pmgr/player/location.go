package player

type Location struct {
	Longitude string // 经度
	Latitude  string // 纬度
	Address   string // 地址
	Ip        string // IP
}

func (p *Player) Location() *Location {
	return p.location
}

func (p *Player) SetLocation(location *Location, ext map[string]any) {
	p.location = location
	if f := p.watch["location"]; f != nil {
		f(p, ext)
	}
}

func (p *Player) WatchLocation(f func(*Player, map[string]any)) {
	p.watch["location"] = f
}
