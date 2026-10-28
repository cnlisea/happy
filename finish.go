package happy

func (h *_Happy) Finish(disband bool, owner bool, ext map[string]any) {
	h.game.Finish(disband, h.begin)

	if h.curRound > 0 && h.costMode == CostModeFinish {
		if h.event != nil && h.event.Cost != nil {
			h.event.Cost(h, h.costMode, false, h.pMgr, h.CopyExtend(ext))
		}
	}

	if h.curRound == 0 && h.costMode == CostModeJoin && h.event != nil && h.event.Cost != nil {
		h.event.Cost(h, h.costMode, true, h.pMgr, h.CopyExtend(ext))
	}

	if h.event != nil && h.event.Finish != nil {
		h.event.Finish(h, h.begin, h.curRound, h.game.RoundNum(), h.pMgr, disband, owner, h.CopyExtend(ext))
	}

	if h.msgChan != nil {
		close(h.msgChan)
		h.msgChan = nil
	}
	panic(PanicDoneExit)
}
