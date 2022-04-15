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

	h.auto.Add(userKey, p.Op())
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
}
