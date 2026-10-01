package happy

func (h *_Happy) Finish(disband bool, owner bool) {
	h.game.Finish(disband, h.begin)

	if h.curRound > 0 && h.costMode == CostModeFinish {
		if h.event != nil && h.event.Cost != nil {
			h.event.Cost(h, h.costMode, false, h.pMgr, h.extend)
		}
	}

	if h.curRound == 0 && h.costMode == CostModeJoin && h.event != nil && h.event.Cost != nil {
		h.event.Cost(h, h.costMode, true, h.pMgr, h.extend)
	}

	if h.event != nil && h.event.Finish != nil {
		h.event.Finish(h, h.begin, h.curRound, h.maxRound, h.pMgr, disband, owner, h.extend)
	}

	if h.msgChan != nil {
		close(h.msgChan)
		h.msgChan = nil
	}
	panic(PanicDoneExit)
}
