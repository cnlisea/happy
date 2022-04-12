package happy

import "github.com/cnlisea/happy/pmgr/player"

func (h *Happy) MsgPlayerJoinHandler(userKey interface{}, p *player.Player) {
	if p == nil {
		return
	}

	var (
		existPlayer = h.pMgr.Get(userKey)
		exist       = existPlayer != nil
		view        = !exist && h.curRound > 0 || exist && existPlayer.View()
	)
	p.SetView(view)
	if !exist {
		// 禁止观战
		if view && !h.game.View() {
			if h.event != nil && h.event.PlayerJoinFail != nil {
				h.event.PlayerJoinFail(userKey, EventPlayerJoinFailKindViewOff, h.extend)
			}
			return
		}

		if !view {
			gameNum := h.pMgr.Len(func(p *player.Player) bool {
				return !p.View()
			})

			// 人数已满
			if maxNum := h.game.PlayerMaxNum(); maxNum > 0 && gameNum >= h.game.PlayerMaxNum() {
				if h.event != nil && h.event.PlayerJoinFail != nil {
					h.event.PlayerJoinFail(userKey, EventPlayerJoinFailKindFull, h.extend)
				}
				return
			}

			if gameNum > 0 {
				// IP相同限制
				if h.game.IpLimit() {
					pLocation := p.Location()
					if pLocation == nil || pLocation.Ip == "" {
						if h.event != nil && h.event.PlayerJoinFail != nil {
							h.event.PlayerJoinFail(userKey, EventPlayerJoinFailKindLocationOff, h.extend)
						}
						return
					}

					var same bool
					h.pMgr.Range(func(key interface{}, p *player.Player) bool {
						if p.View() {
							return true
						}

						if pLocation.Ip == p.Location().Ip {
							same = true
						}
						return !same
					})

					if same {
						if h.event != nil && h.event.PlayerJoinFail != nil {
							h.event.PlayerJoinFail(userKey, EventPlayerJoinFailKindLocationIpSame, h.extend)
						}
						return
					}
				}

				if distance := h.game.DistanceLimit(); distance > 0 {
					// 定位
					pLocation := p.Location()
					if pLocation == nil || pLocation.Longitude == "" || pLocation.Latitude == "" {
						if h.event != nil && h.event.PlayerJoinFail != nil {
							h.event.PlayerJoinFail(userKey, EventPlayerJoinFailKindLocationOff, h.extend)
						}
						return
					}

					h.pMgr.Range(func(key interface{}, p *player.Player) bool {
						if !p.View() {
						}
						return true
					})

					// TODO 距离相近
					if 0 > distance {
						if h.event != nil && h.event.PlayerJoinFail != nil {
							h.event.PlayerJoinFail(userKey, EventPlayerJoinFailKindLocationTooClose, h.extend)
						}
						return
					}
				}
			}
		}

		existPlayer = p
		if h.costMode == CostModeJoin && h.event != nil && h.event.Cost != nil {
			h.event.Cost(h.costMode, false, h.pMgr, h.extend)
		}
	}

	// clean player delay msg
	if exist {
		h.DelayMsgClean(userKey)
	}

	h.pMgr.Add(userKey, existPlayer)
	h.game.PlayerJoin(userKey, view)
	h.game.PlayerOp(userKey, view)
	if h.event != nil && h.event.PlayerJoinSuccess != nil {
		h.event.PlayerJoinSuccess(userKey, h.pMgr, exist, h.extend)
	}

	if !exist && h.roundBeginPolicy == RoundBeginPolicyFullPlayer {
		h.RoundBegin(false)
	}
}
