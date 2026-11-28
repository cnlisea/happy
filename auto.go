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

	h.auto.Add(userKey, p.Auto())
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

func (h *_Happy) AutoPlayerTs(userKey interface{}) (bool, bool, int64) {
	if h.game.Auto() == nil {
		return false, false, 0
	}

	p := h.pMgr.Get(userKey)
	if p == nil {
		return false, false, 0
	}

	if p.View() {
		return false, false, 0
	}

	op, deadlineTs := h.auto.AutoTs(userKey)
	return deadlineTs != 0, op, deadlineTs
}
