package happy

import (
	"context"
	"github.com/cnlisea/happy/proxy"
	"time"
)

type GameBase struct {
	Ctx   context.Context
	PMgr  proxy.GamePMgr
	PMsg  proxy.PlayerMsg
	Delay proxy.GameDelay
	Log   proxy.GameLog
}

func (g *GameBase) Init(ctx context.Context, delay proxy.GameDelay, pMgr proxy.GamePMgr, pMsg proxy.PlayerMsg, log proxy.GameLog) error {
	g.Ctx = ctx
	g.Delay = delay
	g.PMgr = pMgr
	g.PMsg = pMsg
	g.Log = log
	return nil
}

func (g *GameBase) PlayerMaxNum() int {
	return 0
}

func (g *GameBase) PlayerJoin(userKey interface{}, exist bool, view bool) {}

func (g *GameBase) PlayerOp(userKey interface{}, exist bool, view bool) {}

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
