package happy

import "github.com/cnlisea/happy/pmgr/player"

func (h *Happy) MsgDisbandRejectHandler(userKey interface{}) {
	if h.curRound == 0 {
		return
	}

	if h.pMgr.Get(userKey) == nil {
		return
	}

	if h.vote == nil || h.vote.Num() == 0 || h.vote.Exist(userKey) {
		return
	}

	if h.event != nil && h.event.DisbandReject != nil {
		gameNum := h.pMgr.Len(func(p *player.Player) bool {
			return !p.View()
		})
		var op = make(map[interface{}]bool, gameNum)
		h.pMgr.Range(func(key interface{}, p *player.Player) bool {
			if !p.View() {
				op[key] = false
			}
			return true
		})
		h.vote.Range(func(key interface{}) bool {
			op[key] = true
			return true
		})
		h.event.DisbandReject(userKey, h.pMgr, op)
	}
	h.vote.Reset()
}
