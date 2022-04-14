package pmgr

import (
	"container/list"
	"github.com/cnlisea/happy/pmgr/player"
)

type Player struct {
	*player.Player
	key interface{} // 标识
}

func (pm *PMgr) Len(filter ...func(*player.Player) bool) int {
	if filter == nil || len(filter) == 0 {
		return pm.players.Len()
	}

	var num int
	pm.Range(func(key interface{}, p *player.Player) bool {
		for i := range filter {
			if filter[i](p) {
				num++
				break
			}
		}
		return true
	})
	return num
}

func (pm *PMgr) Add(key interface{}, p *player.Player) {
	var (
		e  *list.Element
		pr *Player
	)
	for e = pm.players.Front(); e != nil; e = e.Next() {
		if pr = e.Value.(*Player); pr.key == key {
			pr.Player = p
			return
		}
	}

	var (
		newPlayer = &Player{
			Player: p,
			key:    key,
		}
		kind  WatchKind
		watch Watch
	)
	for kind, watch = range pm.watch {
		switch kind {
		case WatchKindLine:
			newPlayer.WatchLine(func(key interface{}, watch Watch) func(p *player.Player) {
				return func(p *player.Player) {
					watch(key, p)
				}
			}(key, watch))
		case WatchKindReady:
			newPlayer.WatchReady(func(key interface{}, watch Watch) func(p *player.Player) {
				return func(p *player.Player) {
					watch(key, p)
				}
			}(key, watch))
		case WatchKindOp:
			newPlayer.WatchOp(func(key interface{}, watch Watch) func(p *player.Player) {
				return func(p *player.Player) {
					watch(key, p)
				}
			}(key, watch))
		case WatchKindView:
			newPlayer.WatchView(func(key interface{}, watch Watch) func(p *player.Player) {
				return func(p *player.Player) {
					watch(key, p)
				}
			}(key, watch))
		case WatchKindAuto:
			newPlayer.WatchAuto(func(key interface{}, watch Watch) func(p *player.Player) {
				return func(p *player.Player) {
					watch(key, p)
				}
			}(key, watch))
		case WatchKindLocation:
			newPlayer.WatchLocation(func(key interface{}, watch Watch) func(p *player.Player) {
				return func(p *player.Player) {
					watch(key, p)
				}
			}(key, watch))
		case WatchKindScore:
			newPlayer.WatchScore(func(key interface{}, watch Watch) func(p *player.Player) {
				return func(p *player.Player) {
					watch(key, p)
				}
			}(key, watch))
		case WatchKindSite:
			newPlayer.WatchSite(func(key interface{}, watch Watch) func(p *player.Player) {
				return func(p *player.Player) {
					watch(key, p)
				}
			}(key, watch))
		case WatchKindExtend:
			newPlayer.WatchExtend(func(key interface{}, watch Watch) func(p *player.Player) {
				return func(p *player.Player) {
					watch(key, p)
				}
			}(key, watch))
		}
	}

	pm.players.PushBack(newPlayer)
}

func (pm *PMgr) Del(key interface{}) {
	var (
		e *list.Element
		p *Player
	)
	for e = pm.players.Front(); e != nil; e = e.Next() {
		if p = e.Value.(*Player); p.key == key {
			pm.players.Remove(e)
			break
		}
	}
}

func (pm *PMgr) Get(key interface{}) *player.Player {
	var (
		e *list.Element
		p *Player
	)
	for e = pm.players.Front(); e != nil; e = e.Next() {
		if p = e.Value.(*Player); p.key == key {
			return p.Player
		}
	}
	return nil
}

func (pm *PMgr) Range(f func(key interface{}, p *player.Player) bool) {
	var (
		e *list.Element
		p *Player
	)
	for e = pm.players.Front(); e != nil; e = e.Next() {
		if p = e.Value.(*Player); !f(p.key, p.Player) {
			break
		}
	}
}
