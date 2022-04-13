package happy

import (
	"github.com/cnlisea/happy/pmgr"
	"time"
)

type EventPlayerJoinFailKind int

const (
	EventPlayerJoinFailKindFull             EventPlayerJoinFailKind = iota // 人数已满
	EventPlayerJoinFailKindViewOff                                         // 禁止观战
	EventPlayerJoinFailKindLocationOff                                     // 定位未开启
	EventPlayerJoinFailKindLocationIpSame                                  // 定位IP相同
	EventPlayerJoinFailKindLocationTooClose                                // 定位距离过近
)

type Event struct {
	RoundBegin        func(curRound, maxRound uint32, pMgr *pmgr.PMgr, extend map[string]interface{})
	RoundEnd          func(curRound, maxRound uint32, pMgr *pmgr.PMgr, extend map[string]interface{})
	PlayerJoinSuccess func(key interface{}, pMgr *pmgr.PMgr, alreadyExist bool, extend map[string]interface{})
	PlayerJoinFail    func(key interface{}, kind EventPlayerJoinFailKind, extend map[string]interface{})
	PlayerExit        func(key interface{}, pMgr *pmgr.PMgr, extend map[string]interface{})
	PlayerReady       func(key interface{}, pMgr *pmgr.PMgr, extend map[string]interface{})
	PlayerLine        func(key interface{}, pMgr *pmgr.PMgr, extend map[string]interface{})
	PlayerOp          func(key interface{}, pMgr *pmgr.PMgr, extend map[string]interface{})
	PlayerAuto        func(key interface{}, pMgr *pmgr.PMgr, extend map[string]interface{})
	PlayerScore       func(key interface{}, mgr *pmgr.PMgr, extend map[string]interface{})
	Cost              func(mode CostMode, back bool, pMgr *pmgr.PMgr, extend map[string]interface{})
	DisbandAgree      func(ts time.Duration, userKey interface{}, pMgr *pmgr.PMgr, op map[interface{}]bool)
	DisbandReject     func(userKey interface{}, pMgr *pmgr.PMgr, op map[interface{}]bool)
	DisbandPass       func(pMgr *pmgr.PMgr, op map[interface{}]bool)
	DisbandFail       func(pMgr *pmgr.PMgr, op map[interface{}]bool)
	Finish            func(curRound, maxRound uint32, pMgr *pmgr.PMgr, disband bool, extend map[string]interface{})
}
