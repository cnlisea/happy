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

func (p *Player) SetLocation(location *Location) {
	p.location = location
	if f := p.watch["location"]; f != nil {
		f(p)
	}
}

func (p *Player) WatchLocation(f func(*Player)) {
	p.watch["location"] = f
}
