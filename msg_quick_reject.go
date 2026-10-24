package happy

func (h *_Happy) MsgQuickRejectHandler(userKey interface{}, ext map[string]any) {
	if h.curRound > 0 {
		return
	}

	if h.pMgr.Get(userKey) == nil {
		return
	}

	if h.quickVote == nil || h.quickVote.End() {
		return
	}

	h.quickVote.Add(userKey, false, ext)
}
