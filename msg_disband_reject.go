package happy

func (h *_Happy) MsgDisbandRejectHandler(userKey interface{}) {
	if h.curRound == 0 {
		return
	}

	if h.pMgr.Get(userKey) == nil {
		return
	}

	if h.disbandVote == nil {
		return
	}

	h.disbandVote.Add(userKey, false)
}
