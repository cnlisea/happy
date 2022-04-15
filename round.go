package happy

import (
	"github.com/cnlisea/happy/pmgr/player"
)

type RoundBeginPolicy int

const (
	RoundBeginPolicyAllPlayerReady RoundBeginPolicy = iota // 所有玩家准备
	RoundBeginPolicyFullPlayer                             // 满员
)

func (h *_Happy) RoundBeginPolicy(policy RoundBeginPolicy) {
	h.roundBeginPolicy = policy
}

func (h *_Happy) RoundBegin(resume bool, quick bool) {
	if h.begin {
		return
	}

	h.pMgr.Range(func(key interface{}, p *player.Player) bool {
		p.SetReady(false)
		return true
	})

	// clean player kick out
	if h.curRound == 0 {
		h.PlayerKickOutClean()
	}

	h.begin = true
	h.curRound++
	if h.event != nil && h.event.RoundBegin != nil {
		h.event.RoundBegin(h.curRound, h.maxRound, h.pMgr, h.extend)
	}
	if !resume && h.costMode == CostModeFirstRoundBegin || h.costMode == CostModeRoundBegin {
		if h.event != nil && h.event.Cost != nil {
			h.event.Cost(h.costMode, false, h.pMgr, h.extend)
		}
	}
	h.game.Begin(quick)
}

func (h *_Happy) RoundEnd() {
	if !h.begin {
		return
	}

	h.begin = false
	if h.event != nil && h.event.RoundEnd != nil {
		h.event.RoundEnd(h.curRound, h.maxRound, h.pMgr, h.extend)
	}

	if h.costMode == CostModeFirstRoundEnd || h.costMode == CostModeRoundEnd {
		if h.event != nil && h.event.Cost != nil {
			h.event.Cost(h.costMode, false, h.pMgr, h.extend)
		}
	}
	h.game.End()

	if h.curRound == h.maxRound {
		h.Finish(false)
	}
}
