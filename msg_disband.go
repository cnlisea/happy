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

	if h.disbandVote == nil {
		gamePlayerLen := h.pMgr.Len(func(p *player.Player) bool {
			return !p.View()
		})
		h.disbandVote = vote.New(gamePlayerLen, gamePlayerLen)
		h.disbandVote.Deadline(h.delay, h.game.DisbandTs(), false, true)
		h.disbandVote.CallbackPass(func() {
			if h.event != nil && h.event.DisbandPass != nil {
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
				h.disbandVote.Range(func(key interface{}, o bool) bool {
					op[key] = o
					return true
				})
				h.event.DisbandPass(h.pMgr, op)
			}
			h.Finish(true)
		})
		h.disbandVote.CallbackFail(func() {
			if h.event != nil && h.event.DisbandFail != nil {
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
				h.disbandVote.Range(func(key interface{}, o bool) bool {
					op[key] = o
					return true
				})
				h.event.DisbandFail(h.pMgr, op)
			}
		})
		h.disbandVote.CallbackAdd(func(key interface{}, agree bool) {
			if h.event == nil || (agree && h.event.DisbandAgree == nil) || (!agree && h.event.DisbandReject == nil) {
				return
			}
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
			h.disbandVote.Range(func(key interface{}, o bool) bool {
				op[key] = o
				return true
			})
			switch agree {
			case true:
				h.event.DisbandAgree(h.game.DisbandTs(), userKey, h.pMgr, op)
			default:
				h.event.DisbandReject(userKey, h.pMgr, op)
			}
		})
	}

	if h.disbandVote.End() {
		h.disbandVote.Reset()
	}

	h.disbandVote.Add(userKey, true)
}
