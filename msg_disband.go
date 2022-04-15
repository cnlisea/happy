package happy

import (
	"github.com/cnlisea/happy/pmgr/player"
	"github.com/cnlisea/happy/vote"
	"time"
)

func (h *_Happy) MsgDisbandHandler(userKey interface{}) {
	if h.curRound == 0 {
		return
	}

	if h.pMgr.Get(userKey) == nil {
		return
	}

	if h.disbandVote == nil {
		var (
			gamePlayerLen = h.pMgr.Len(func(p *player.Player) bool {
				return !p.View()
			})
			minNum = gamePlayerLen
		)
		if h.plugin != nil && h.plugin.DisbandMinAgreeNum != nil {
			if minNum = h.plugin.DisbandMinAgreeNum(gamePlayerLen); minNum <= 0 {
				minNum = gamePlayerLen
			}
		}
		h.disbandVote = vote.New(minNum, gamePlayerLen)
		var pass bool
		if h.plugin != nil && h.plugin.DisbandDeadlinePass != nil {
			pass = h.plugin.DisbandDeadlinePass(h.extend)
		}
		h.disbandVote.Deadline(h.delay, h.game.DisbandTs(), pass, true)
		h.disbandVote.CallbackPass(func(ts int64) {
			if h.event != nil && h.event.DisbandPass != nil {
				var op = make(map[interface{}]bool, h.disbandVote.Num())
				h.disbandVote.Range(func(key interface{}, o bool) bool {
					op[key] = o
					return true
				})
				h.event.DisbandPass(ts, h.pMgr, op, h.extend)
			}
			h.Finish(true)
		})
		h.disbandVote.CallbackFail(func(ts int64) {
			if h.event != nil && h.event.DisbandFail != nil {
				var op = make(map[interface{}]bool, h.disbandVote.Num())
				h.disbandVote.Range(func(key interface{}, o bool) bool {
					op[key] = o
					return true
				})
				h.event.DisbandFail(ts, h.pMgr, op, h.extend)
			}
		})
		h.disbandVote.CallbackAdd(func(ts time.Duration, deadlineTs int64, key interface{}, agree bool) {
			if h.event == nil || (agree && h.event.DisbandAgree == nil) || (!agree && h.event.DisbandReject == nil) {
				return
			}
			var op = make(map[interface{}]bool, h.disbandVote.Num())
			h.disbandVote.Range(func(key interface{}, o bool) bool {
				op[key] = o
				return true
			})
			switch agree {
			case true:
				h.event.DisbandAgree(ts, deadlineTs, userKey, h.pMgr, op, h.extend)
			default:
				h.event.DisbandReject(ts, deadlineTs, userKey, h.pMgr, op, h.extend)
			}
		})
	}

	if h.disbandVote.End() {
		h.disbandVote.Reset()
	}

	h.disbandVote.Add(userKey, true)
}
