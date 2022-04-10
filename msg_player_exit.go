package happy

import "github.com/cnlisea/happy/pmgr/player"

func (h *Happy) MsgPlayerExitHandler(userKey interface{}) {
	if h.curRound > 0 {
		return
	}

	p := h.pMgr.Get(userKey)
	if p == nil {
		return
	}

	h.game.PlayerExit(userKey, p.View())
	if !p.View() {
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
		h.event.PlayerExit(userKey, h.pMgr, h.extend)
	}
	h.pMgr.Del(userKey)
}
