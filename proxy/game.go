package proxy

import (
	"github.com/cnlisea/happy/pmgr"
	"time"
)

type GameAuto struct {
	AutoTs    time.Duration
	ReadyOpTs time.Duration
	OpTs      time.Duration
}

type Game interface {
	Init(pMgr *pmgr.PMgr, pMsg PlayerMsg) error
	PlayerMaxNum() int
	PlayerJoin(userKey interface{}, view bool)
	PlayerOp(userKey interface{}, view bool)
	PlayerExit(userKey interface{}, view bool)
	PlayerOfflineKickOut() time.Duration
	PlayerAuto(userKey interface{})
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
