package proxy

import (
	"context"
	"github.com/cnlisea/happy/log"
	"github.com/cnlisea/happy/pmgr/player"
	"time"
)

type GameAuto struct {
	AutoTs    time.Duration
	ReadyOpTs time.Duration
	OpTs      time.Duration
}

type Game interface {
	Init(ctx context.Context, delay GameDelay, pMgr GamePMgr, pMsg PlayerMsg, log GameLog) error
	PlayerMaxNum() int
	PlayerJoin(userKey interface{}, view bool)
	PlayerOp(userKey interface{}, view bool)
	PlayerExit(userKey interface{}, view bool)
	PlayerOfflineKickOut() time.Duration
	PlayerAuto(userKey interface{})
	Msg(userKey interface{}, data interface{})
	Begin(quick bool)
	End()
	Auto() *GameAuto
	Quick(num int) bool
	QuickTs() time.Duration
	View() bool
	DisbandTs() time.Duration
	IpLimit() bool
	DistanceLimit() int
	Finish(disband bool)
}

type GameDelay interface {
	DelayFunc(delayTs time.Duration, f func(args interface{}), args interface{})
	DelayMsg(delayTs time.Duration, userKey []interface{}, data ...interface{})
}

type GamePMgr interface {
	Get(key interface{}) *player.Player
	Len(filter ...func(*player.Player) bool) int
	Range(f func(key interface{}, p *player.Player) bool)
}

type GameLog interface {
	Debug(s string, fields ...log.Field)
	Info(s string, fields ...log.Field)
	Warn(s string, fields ...log.Field)
	Err(s string, fields ...log.Field)
}
