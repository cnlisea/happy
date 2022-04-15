package happy

import (
	"github.com/cnlisea/happy/pmgr/player"
)

func (h *_Happy) MsgPlayerReadyHandler(userKey interface{}, site uint32) {
	if h.begin {
		return
	}

	p := h.pMgr.Get(userKey)
	if p == nil {
		return
	}

	if p.View() {
		return
	}

	if site > 0 && p.Site() != site {
		if h.curRound > 0 {
			return
		}

		if int(site) > h.game.PlayerMaxNum() {
			return
		}

		var exist bool
		h.pMgr.Range(func(key interface{}, p *player.Player) bool {
			if p.Site() == site {
				exist = true
			}
			return !exist
		})
		if exist {
			return
		}

		p.SetSite(site)
	}

	p.SetReady(!p.Ready())

	if maxNum := h.game.PlayerMaxNum(); maxNum <= 0 || h.pMgr.Len(func(p *player.Player) bool {
		return !p.View()
	}) == maxNum {
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
