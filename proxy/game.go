package proxy

import (
	"github.com/cnlisea/happy/pmgr"
	"time"
)

type Game interface {
	Init(pMgr *pmgr.PMgr, pMsg PlayerMsg) error
	PlayerMaxNum() int
	PlayerJoin(userKey interface{}, view bool)
	PlayerOp(userKey interface{}, view bool)
	PlayerExit(userKey interface{}, view bool)
	PlayerOfflineKickOut() time.Duration
	Begin()
	End()
	View() bool
	DisbandTs() int64
	IpLimit() bool
	DistanceLimit() int
	Finish(disband bool)
}
