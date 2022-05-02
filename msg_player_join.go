package happy

import (
	"github.com/cnlisea/happy/pmgr/player"
)

func (h *_Happy) MsgPlayerJoinHandler(userKey interface{}, p *player.Player) {
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
				h.event.PlayerJoinFail(h, userKey, EventPlayerJoinFailKindViewOff, h.extend)
			}
			return
		}

		if !view {
			gameNum := h.pMgr.Len(func(p *player.Player) bool {
				return !p.View()
			})

			// 人数已满
			if maxNum := h.game.PlayerMaxNum(); maxNum > 0 && gameNum >= maxNum {
				if h.event != nil && h.event.PlayerJoinFail != nil {
					h.event.PlayerJoinFail(h, userKey, EventPlayerJoinFailKindFull, h.extend)
				}
				return
			}

			// IP相同限制
			if h.game.IpLimit() {
				pLocation := p.Location()
				if pLocation == nil || pLocation.Ip == "" {
					if h.event != nil && h.event.PlayerJoinFail != nil {
						h.event.PlayerJoinFail(h, userKey, EventPlayerJoinFailKindLocationOff, h.extend)
					}
					return
				}

				if gameNum > 0 {
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
							h.event.PlayerJoinFail(h, userKey, EventPlayerJoinFailKindLocationIpSame, h.extend)
						}
						return
					}
				}
			}

			distanceLimit := h.game.DistanceLimit()
			// 定位
			pLocation := p.Location()
			if distanceLimit > 0 && (pLocation == nil || pLocation.Longitude == "" || pLocation.Latitude == "") {
				if h.event != nil && h.event.PlayerJoinFail != nil {
					h.event.PlayerJoinFail(h, userKey, EventPlayerJoinFailKindLocationOff, h.extend)
				}
				return
			}

			if distanceLimit > 0 && (h.plugin == nil || h.plugin.LocationDistance == nil) {
				if h.event != nil && h.event.PlayerJoinFail != nil {
					h.event.PlayerJoinFail(h, userKey, EventPlayerJoinFailKindLocationTooClose, h.extend)
				}
				return
			}

			if gameNum > 0 {
				desLocation := make([]*player.Location, 0, gameNum)
				h.pMgr.Range(func(key interface{}, p *player.Player) bool {
					if !p.View() {
						desLocation = append(desLocation, p.Location())
					}
					return true
				})

				var distances []int
				if h.plugin != nil && h.plugin.LocationDistance != nil {
					distances, _ = h.plugin.LocationDistance(pLocation, desLocation...)
				}
				var tooClose bool
				if distanceLimit > 0 {
					tooClose = len(distances) == 0
					if !tooClose {
						for i := range distances {
							if distances[i] < distanceLimit {
								tooClose = true
								break
							}
						}
					}
				}

				if tooClose {
					if h.event != nil && h.event.PlayerJoinFail != nil {
						h.event.PlayerJoinFail(h, userKey, EventPlayerJoinFailKindLocationTooClose, h.extend)
					}
					return
				}
			}

			var site = uint32(1)
			if gameNum > 0 {
				var siteIndex int
				h.pMgr.Range(func(key interface{}, p *player.Player) bool {
					if !p.View() {
						siteIndex = siteIndex | (1 << (p.Site() - 1))
					}
					return true
				})
				for siteIndex&1 == 1 {
					siteIndex = siteIndex >> 1
					site++
				}
			}
			p.SetSite(site)
			// cancel quick vote
			if h.quickVote != nil {
				h.quickVote.Cancel()
			}
		}

		existPlayer = p
		if h.costMode == CostModeJoin && h.event != nil && h.event.Cost != nil {
			h.event.Cost(h, h.costMode, false, h.pMgr, h.extend)
		}
	}

	// clean player delay msg
	if exist {
		h.DelayMsgClean(userKey)
	}

	h.pMgr.Add(userKey, existPlayer)
	h.game.PlayerJoin(userKey, exist, view)
	if h.event != nil && h.event.PlayerJoinSuccess != nil {
		h.event.PlayerJoinSuccess(h, userKey, h.pMgr, exist, h.extend)
	}
	h.game.PlayerOp(userKey, exist, view)

	if !exist &&
		h.roundBeginPolicy == RoundBeginPolicyFullPlayer &&
		h.pMgr.Len(func(p *player.Player) bool {
			return !p.View()
		}) == h.game.PlayerMaxNum() {
		h.RoundBegin(false, false)
	}
}
