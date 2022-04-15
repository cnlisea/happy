package happy

type _DelayPlayerKickOut struct {
	UserKey interface{}
}

func (h *_Happy) PlayerKickOut(userKey interface{}) {
	if h.curRound > 0 {
		return
	}

	p := h.pMgr.Get(userKey)
	if p == nil {
		return
	}

	if p.View() {
		return
	}

	ts := h.game.PlayerOfflineKickOut()
	if ts <= 0 {
		return
	}

	h.PlayerKickOutDel(userKey)
	if !p.Line() {
		h.delay.Add(ts, func(ts int64, args interface{}) {
			h.MsgPlayerExitHandler(args.(*_DelayPlayerKickOut).UserKey)
		}, _DelayPlayerKickOut{
			UserKey: userKey,
		})
	}
}

func (h *_Happy) PlayerKickOutDel(userKey interface{}) {
	var (
		delay *_DelayPlayerKickOut
		ok    bool
	)
	h.delay.Del(func(ts int64, args interface{}) bool {
		delay, ok = args.(*_DelayPlayerKickOut)
		return ok && delay.UserKey == userKey
	})
}

func (h *_Happy) PlayerKickOutClean() {
	ts := h.game.PlayerOfflineKickOut()
	if ts <= 0 {
		return
	}

	var ok bool
	h.delay.Del(func(ts int64, args interface{}) bool {
		_, ok = args.(*_DelayPlayerKickOut)
		return ok
	})
}
