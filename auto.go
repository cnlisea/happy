package happy

func (h *_Happy) AutoPlayer(userKey interface{}) {
	if h.game.Auto() == nil {
		return
	}

	p := h.pMgr.Get(userKey)
	if p == nil {
		return
	}

	if p.View() {
		return
	}

	auto := p.Auto()
	h.auto.Add(userKey, auto)
	if !auto {
		p.SetAutoTs(h.auto.AutoTs(userKey, false))
	}
}

func (h *_Happy) AutoPlayerDel(userKey interface{}) {
	if h.game.Auto() == nil {
		return
	}

	p := h.pMgr.Get(userKey)
	if p == nil {
		return
	}

	if p.View() {
		return
	}

	h.auto.Del(userKey)
	p.SetAutoTs(0)
}
