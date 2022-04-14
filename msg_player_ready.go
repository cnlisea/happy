package happy

import "github.com/cnlisea/happy/pmgr/player"

func (h *Happy) MsgPlayerReadyHandler(userKey interface{}) {
	if h.curRound > 0 {
		return
	}

	p := h.pMgr.Get(userKey)
	if p == nil {
		return
	}

	if p.View() {
		return
	}

	p.SetReady(!p.Ready())
	if h.roundBeginPolicy == RoundBeginPolicyAllPlayerReady {
		var allReady = true
		h.pMgr.Range(func(key interface{}, p *player.Player) bool {
			if !p.View() && !p.Ready() {
				allReady = false
			}
			return allReady
		})
		if allReady {
			h.RoundBegin(false, false)
		}
	}
}
