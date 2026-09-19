package happy

import "github.com/cnlisea/happy/pmgr/player"

func (h *_Happy) MsgPlayerKickOutHandler(userKey any, data any) {
	p := h.pMgr.Get(userKey)
	if p == nil {
		return
	}

	view := p.View()
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

	if h.event != nil && h.event.PlayerKickOut != nil {
		h.event.PlayerKickOut(h, userKey, data, h.pMgr, h.extend)
	}
	h.pMgr.Del(userKey)

	// cancel quick vote
	if !view && h.quickVote != nil {
		h.quickVote.Cancel()
	}
}
