package happy

import (
	"context"
	"github.com/cnlisea/happy/pmgr/player"
	"github.com/cnlisea/happy/proxy"
	"time"
)

type GameBase struct {
	Ctx   context.Context
	PMgr  proxy.GamePMgr
	PMsg  proxy.PlayerMsg
	Delay proxy.GameDelay
}

func (g *GameBase) Init(ctx context.Context, delay proxy.GameDelay, pMgr proxy.GamePMgr, pMsg proxy.PlayerMsg) error {
	g.Ctx = ctx
	g.Delay = delay
	g.PMgr = pMgr
	g.PMsg = pMsg
	return nil
}

func (g *GameBase) PlayerMaxNum() int {
	return 0
}

func (g *GameBase) PlayerJoin(userKey interface{}, view bool) {}

func (g *GameBase) PlayerOp(userKey interface{}, view bool) {}

func (g *GameBase) PlayerExit(userKey interface{}, view bool) {}

func (g *GameBase) PlayerOfflineKickOut() time.Duration {
	return 0
}

func (g *GameBase) PlayerAuto(userKey interface{}) {}

func (g *GameBase) Msg(userKey interface{}, data interface{}) {}

func (g *GameBase) Begin(quick bool) {}

func (g *GameBase) End() {}

func (g *GameBase) GameEnd() {
	panic(PanicGameEnd)
}

func (g *GameBase) Auto() *proxy.GameAuto {
	return nil
}

func (g *GameBase) Quick(num int) bool {
	return false
}

func (g *GameBase) QuickTs() time.Duration {
	return 0
}

func (g *GameBase) View() bool {
	return false
}

func (g *GameBase) DisbandTs() time.Duration {
	return 2 * time.Minute
}

func (g *GameBase) IpLimit() bool {
	return false
}

func (g *GameBase) DistanceLimit() int {
	return 0
}

func (g *GameBase) Finish(disband bool) {}

func (g *GameBase) BroadCastDelayMsg(delay time.Duration, data ...interface{}) {
	if g.PMgr == nil || g.Delay == nil || len(data) == 0 {
		return
	}

	if delay <= 0 {
		g.BroadCastMsg(data...)
		return
	}

	keysLen := g.PMgr.Len(func(p *player.Player) bool {
		return !p.View()
	})
	if keysLen == 0 {
		return
	}

	keys := make([]interface{}, 0, keysLen)
	g.PMgr.Range(func(key interface{}, p *player.Player) bool {
		if !p.View() {
			keys = append(keys, key)
		}
		return true
	})

	g.Delay.DelayMsg(delay, keys, data...)
}

func (g *GameBase) BroadCastMsg(data ...interface{}) {
	if g.PMgr == nil || g.PMsg == nil || len(data) == 0 {
		return
	}

	keysLen := g.PMgr.Len(func(p *player.Player) bool {
		return !p.View()
	})
	if keysLen == 0 {
		return
	}

	keys := make([]interface{}, 0, keysLen)
	g.PMgr.Range(func(key interface{}, p *player.Player) bool {
		if !p.View() {
			keys = append(keys, key)
		}
		return true
	})

	g.PMsg.Send(g.Ctx, keys, data...)
}

func (g *GameBase) BroadCastDelayMsgView(delay time.Duration, data ...interface{}) {
	if g.PMgr == nil || g.Delay == nil || len(data) == 0 {
		return
	}

	if delay <= 0 {
		g.BroadcastMsgView(data...)
		return
	}

	keysLen := g.PMgr.Len(func(p *player.Player) bool {
		return p.View()
	})
	if keysLen == 0 {
		return
	}

	keys := make([]interface{}, 0, keysLen)
	g.PMgr.Range(func(key interface{}, p *player.Player) bool {
		if p.View() {
			keys = append(keys, key)
		}
		return true
	})

	g.Delay.DelayMsg(delay, keys, data...)
}

func (g *GameBase) BroadcastMsgView(data ...interface{}) {
	if g.PMgr == nil || g.PMsg == nil || len(data) == 0 {
		return
	}

	keysLen := g.PMgr.Len(func(p *player.Player) bool {
		return p.View()
	})
	if keysLen == 0 {
		return
	}

	keys := make([]interface{}, 0, keysLen)
	g.PMgr.Range(func(key interface{}, p *player.Player) bool {
		if p.View() {
			keys = append(keys, key)
		}
		return true
	})

	g.PMsg.Send(g.Ctx, keys, data...)
}
