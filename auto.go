package happy

func (h *Happy) AutoPlayer(userKey interface{}) {
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

	h.auto.Add(userKey, p.Op())
}

func (h *Happy) AutoPlayerDel(userKey interface{}) {
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
}
