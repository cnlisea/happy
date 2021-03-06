package happy

func (h *_Happy) Finish(disband bool) {
	h.game.Finish(disband)

	if h.curRound > 0 && h.costMode == CostModeFinish {
		if h.event != nil && h.event.Cost != nil {
			h.event.Cost(h, h.costMode, false, h.pMgr, h.extend)
		}
	}

	if h.curRound == 0 && h.costMode == CostModeJoin && h.event != nil && h.event.Cost != nil {
		h.event.Cost(h, h.costMode, true, h.pMgr, h.extend)
	}

	if h.event != nil && h.event.Finish != nil {
		h.event.Finish(h, h.curRound, h.maxRound, h.pMgr, disband, h.extend)
	}

	if h.msgChan != nil {
		close(h.msgChan)
		h.msgChan = nil
	}
	panic(PanicDoneExit)
}
