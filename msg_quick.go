package happy

import (
	"github.com/cnlisea/happy/pmgr/player"
	"github.com/cnlisea/happy/vote"
	"time"
)

func (h *_Happy) MsgQuickHandler(userKey interface{}) {
	if h.quickVote == nil || h.quickVote.End() {
		if h.curRound > 0 {
			return
		}

		gamePlayerLen := h.pMgr.Len(func(p *player.Player) bool {
			return !p.View()
		})
		if gamePlayerLen == h.game.PlayerMaxNum() {
			return
		}

		if !h.game.Quick(gamePlayerLen) {
			return
		}
	}

	if h.pMgr.Get(userKey) == nil {
		return
	}

	if h.quickVote == nil {
		var (
			gamePlayerLen = h.pMgr.Len(func(p *player.Player) bool {
				return !p.View()
			})
			minNum = gamePlayerLen
		)
		if h.plugin != nil && h.plugin.QuickMinAgreeNum != nil {
			if minNum = h.plugin.QuickMinAgreeNum(gamePlayerLen); minNum <= 0 {
				minNum = gamePlayerLen
			}
		}
		h.quickVote = vote.New(minNum, gamePlayerLen)
		var pass bool
		if h.plugin != nil && h.plugin.QuickDeadlinePass != nil {
			pass = h.plugin.QuickDeadlinePass(h.extend)
		}
		h.quickVote.Deadline(h.delay, h.game.QuickTs(), pass, true)
		h.quickVote.CallbackPass(func(ts int64) {
			if h.event != nil && h.event.QuickPass != nil {
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
				h.quickVote.Range(func(key interface{}, o bool) bool {
					op[key] = o
					return true
				})
				h.event.QuickPass(ts, h.pMgr, op, h.extend)
			}
			h.RoundBegin(false, true)
		})
		h.quickVote.CallbackFail(func(ts int64) {
			if h.event != nil && h.event.QuickFail != nil {
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
				h.quickVote.Range(func(key interface{}, o bool) bool {
					op[key] = o
					return true
				})
				h.event.QuickFail(ts, h.pMgr, op, h.extend)
			}
		})
		h.quickVote.CallbackAdd(func(ts time.Duration, deadlineTs int64, key interface{}, agree bool) {
			if h.event == nil || (agree && h.event.QuickAgree == nil) || (!agree && h.event.QuickReject == nil) {
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
			h.quickVote.Range(func(key interface{}, o bool) bool {
				op[key] = o
				return true
			})
			switch agree {
			case true:
				h.event.QuickAgree(ts, deadlineTs, userKey, h.pMgr, op, h.extend)
			default:
				h.event.QuickReject(ts, deadlineTs, userKey, h.pMgr, op, h.extend)
			}
		})
	}

	if h.quickVote.End() {
		h.quickVote.Reset()
	}
	h.quickVote.Add(userKey, true)
}
