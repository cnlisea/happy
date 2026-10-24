package happy

func (h *_Happy) MsgGameHandler(userKey interface{}, data interface{}, ext map[string]any) {
	p := h.pMgr.Get(userKey)
	if p == nil {
		return
	}

	h.game.Msg(h.begin, userKey, data, ext)
}
