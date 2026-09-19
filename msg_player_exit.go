package happy

import "github.com/cnlisea/happy/pmgr/player"

func (h *_Happy) MsgPlayerExitHandler(userKey interface{}) {
	p := h.pMgr.Get(userKey)
	if p == nil {
		return
	}

	view := p.View()
	if !view && h.curRound > 0 {
		return
	}

	h.game.PlayerExit(userKey, view)
	if !view {
		if h.plugin != nil && h.plugin.PlayerExitDisband != nil && h.plugin.PlayerExitDisband(userKey, h.ownerUserKey, h.extend) {
			h.Finish(true)
			return
		}

		if h.pMgr.Len(func(p *player.Player) bool {
			return !p.View()
		}) == 0 {
			h.Finish(true)
			return
		}
	}
	if h.event != nil && h.event.PlayerExit != nil {
		h.event.PlayerExit(h, userKey, h.pMgr, h.extend)
	}
	h.pMgr.Del(userKey)

	// cancel quick vote
	if h.quickVote != nil {
		h.quickVote.Cancel()
	}
}
