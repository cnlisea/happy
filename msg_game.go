package happy

func (h *_Happy) MsgGameHandler(userKey interface{}, data interface{}) {
	p := h.pMgr.Get(userKey)
	if p == nil {
		return
	}

	h.game.Msg(userKey, data)
}
