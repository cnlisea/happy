package happy

func (h *Happy) Finish(disband bool) {
	h.game.Finish(disband)
	if h.event != nil && h.event.Finish != nil {
		h.event.Finish(h.curRound, h.maxRound, h.pMgr, disband, h.extend)
	}

	if h.costMode == CostModeFinish {
		if h.event != nil && h.event.Cost != nil {
			h.event.Cost(h.costMode, h.pMgr, h.extend)
		}
	}

	if h.msgChan != nil {
		close(h.msgChan)
		h.msgChan = nil
	}
	panic(PanicDoneExit)
}
