package happy

import (
	"github.com/cnlisea/happy/pmgr/player"
	"github.com/cnlisea/happy/vote"
)

func (h *Happy) MsgDisbandHandler(userKey interface{}) {
	if h.curRound == 0 {
		return
	}

	if h.pMgr.Get(userKey) == nil {
		return
	}

	if h.vote == nil {
		h.vote = vote.New(h.pMgr.Len(func(p *player.Player) bool {
			return !p.View()
		}), func() {
			if h.event != nil && h.event.DisbandFull != nil {
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
				h.event.DisbandFull(h.pMgr, op)
				h.vote.Reset()
			}
			h.Finish(true)
		})
	}

	// already exist
	if h.vote.Exist(userKey) {
		return
	}

	h.vote.Add(userKey)
	if h.event != nil && h.event.DisbandAgree != nil {
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
		h.event.DisbandAgree(0, userKey, h.pMgr, op)
	}
}
