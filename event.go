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
	RoundBegin        func(h Happy, curRound, maxRound uint32, pMgr *pmgr.PMgr, extend map[string]interface{})
	RoundEnd          func(h Happy, curRound, maxRound uint32, pMgr *pmgr.PMgr, extend map[string]interface{})
	PlayerJoinSuccess func(h Happy, key interface{}, pMgr *pmgr.PMgr, alreadyExist bool, extend map[string]interface{})
	PlayerJoinFail    func(h Happy, key interface{}, kind EventPlayerJoinFailKind, extend map[string]interface{})
	PlayerExit        func(h Happy, key interface{}, pMgr *pmgr.PMgr, extend map[string]interface{})
	PlayerReady       func(h Happy, key interface{}, pMgr *pmgr.PMgr, extend map[string]interface{})
	PlayerLine        func(h Happy, key interface{}, pMgr *pmgr.PMgr, extend map[string]interface{})
	PlayerOp          func(h Happy, key interface{}, pMgr *pmgr.PMgr, extend map[string]interface{})
	PlayerAuto        func(h Happy, key interface{}, pMgr *pmgr.PMgr, extend map[string]interface{})
	PlayerSite        func(h Happy, key interface{}, mgr *pmgr.PMgr, extend map[string]interface{})
	PlayerScore       func(h Happy, key interface{}, mgr *pmgr.PMgr, extend map[string]interface{})
	Cost              func(h Happy, mode CostMode, back bool, pMgr *pmgr.PMgr, extend map[string]interface{})
	DisbandAgree      func(h Happy, ts time.Duration, deadlineTs int64, userKey interface{}, pMgr *pmgr.PMgr, op map[interface{}]bool, extend map[string]interface{})
	DisbandReject     func(h Happy, ts time.Duration, deadlineTs int64, userKey interface{}, pMgr *pmgr.PMgr, op map[interface{}]bool, extend map[string]interface{})
	DisbandPass       func(h Happy, deadlineTs int64, pMgr *pmgr.PMgr, op map[interface{}]bool, extend map[string]interface{})
	DisbandFail       func(h Happy, deadlineTs int64, pMgr *pmgr.PMgr, op map[interface{}]bool, extend map[string]interface{})
	QuickAgree        func(h Happy, ts time.Duration, deadlineTs int64, userKey interface{}, pMgr *pmgr.PMgr, op map[interface{}]bool, extend map[string]interface{})
	QuickReject       func(h Happy, ts time.Duration, deadlineTs int64, userKey interface{}, pMgr *pmgr.PMgr, op map[interface{}]bool, extend map[string]interface{})
	QuickPass         func(h Happy, deadlineTs int64, pMgr *pmgr.PMgr, op map[interface{}]bool, extend map[string]interface{})
	QuickFail         func(h Happy, deadlineTs int64, pMgr *pmgr.PMgr, op map[interface{}]bool, extend map[string]interface{})
	Finish            func(h Happy, curRound, maxRound uint32, pMgr *pmgr.PMgr, disband bool, extend map[string]interface{})
}

func (h *_Happy) Event(e *Event) {
	h.event = e
}
