package happy

import (
	"testing"
	"time"

	"github.com/cnlisea/happy/pmgr"
	"github.com/cnlisea/happy/pmgr/player"
	"github.com/cnlisea/happy/proxy"
)

func Test_Happy_EventRoundBegin(t *testing.T) {
	h := New(nil, 1, new(GameBase), nil)
	h.Heartbeat(3 * time.Second)
	h.Event(&Event{
		RoundBegin: func(curRound, maxRound uint32, pMgr *pmgr.PMgr, extend map[string]interface{}) {
			t.Log("round begin", curRound, maxRound, pMgr.Len())
		},
		Finish: func(curRound, maxRound uint32, pMgr *pmgr.PMgr, disband bool, extend map[string]interface{}) {
			t.Log("finish", time.Now().Unix())
		},
	})
	h.RoundBeginPolicy(RoundBeginPolicyAllPlayerReady)
	h.Init()

	go func() {
		h.Msg(&proxy.Msg{
			Kind:    proxy.MsgKindPlayerJoin,
			UserKey: 1,
			Data: &proxy.MsgKindPlayerJoinData{
				Player: player.New(),
			},
		})

		h.Msg(&proxy.Msg{
			Kind:    proxy.MsgKindPlayerReady,
			UserKey: 1,
			Data: &proxy.MsgKindPlayerReadyData{
				Site: 1,
			},
		})
	}()
	t.Log("run", time.Now().Unix())
	h.Run(false)
}

type EventRoundEndGame struct {
	GameBase
}

func (g *EventRoundEndGame) Begin(quick bool) {
	g.GameEnd()
}

func Test_Happy_EventRoundEnd(t *testing.T) {
	h := New(nil, 1, new(EventRoundEndGame), nil)
	h.Heartbeat(3 * time.Second)
	h.Event(&Event{
		RoundEnd: func(curRound, maxRound uint32, pMgr *pmgr.PMgr, extend map[string]interface{}) {
			t.Log("round end", curRound, maxRound, pMgr.Len())
		},
		Finish: func(curRound, maxRound uint32, pMgr *pmgr.PMgr, disband bool, extend map[string]interface{}) {
			t.Log("finish", time.Now().Unix())
		},
	})
	h.RoundBeginPolicy(RoundBeginPolicyAllPlayerReady)
	h.Init()

	go func() {
		h.Msg(&proxy.Msg{
			Kind:    proxy.MsgKindPlayerJoin,
			UserKey: 1,
			Data: &proxy.MsgKindPlayerJoinData{
				Player: player.New(),
			},
		})

		h.Msg(&proxy.Msg{
			Kind:    proxy.MsgKindPlayerReady,
			UserKey: 1,
			Data: &proxy.MsgKindPlayerReadyData{
				Site: 1,
			},
		})
	}()
	t.Log("run", time.Now().Unix())
	h.Run(false)
}

func Test_Happy_EventPlayerJoinSuccess(t *testing.T) {
	h := New(nil, 1, new(GameBase), nil)
	h.Heartbeat(3 * time.Second)
	h.Event(&Event{
		PlayerJoinSuccess: func(key interface{}, pMgr *pmgr.PMgr, alreadyExist bool, extend map[string]interface{}) {
			t.Log("user join success", key, pMgr.Len())
		},
		Finish: func(curRound, maxRound uint32, pMgr *pmgr.PMgr, disband bool, extend map[string]interface{}) {
			t.Log("finish", time.Now().Unix())
		},
	})
	h.RoundBeginPolicy(RoundBeginPolicyAllPlayerReady)
	h.Init()

	go func() {
		h.Msg(&proxy.Msg{
			Kind:    proxy.MsgKindPlayerJoin,
			UserKey: 1,
			Data: &proxy.MsgKindPlayerJoinData{
				Player: player.New(),
			},
		})
	}()
	t.Log("run", time.Now().Unix())
	h.Run(false)
}

type EventPlayerJoinFailGame struct {
	GameBase
}

func (g *EventPlayerJoinFailGame) PlayerMaxNum() int {
	return 1
}

func Test_Happy_EventPlayerJoinFail_Full(t *testing.T) {
	h := New(nil, 1, new(EventPlayerJoinFailGame), nil)
	h.Heartbeat(3 * time.Second)
	h.Event(&Event{
		PlayerJoinFail: func(key interface{}, kind EventPlayerJoinFailKind, extend map[string]interface{}) {
			t.Log("join fail", key, kind)
		},
		Finish: func(curRound, maxRound uint32, pMgr *pmgr.PMgr, disband bool, extend map[string]interface{}) {
			t.Log("finish", time.Now().Unix())
		},
	})
	h.RoundBeginPolicy(RoundBeginPolicyAllPlayerReady)
	h.Init()

	go func() {
		h.Msg(&proxy.Msg{
			Kind:    proxy.MsgKindPlayerJoin,
			UserKey: 1,
			Data: &proxy.MsgKindPlayerJoinData{
				Player: player.New(),
			},
		})

		h.Msg(&proxy.Msg{
			Kind:    proxy.MsgKindPlayerJoin,
			UserKey: 2,
			Data: &proxy.MsgKindPlayerJoinData{
				Player: player.New(),
			},
		})
	}()
	t.Log("run", time.Now().Unix())
	h.Run(false)
}

func Test_Happy_EventPlayerJoinFail_ViewOff(t *testing.T) {
	h := New(nil, 1, new(EventPlayerJoinFailGame), nil)
	h.Heartbeat(3 * time.Second)
	h.Event(&Event{
		PlayerJoinFail: func(key interface{}, kind EventPlayerJoinFailKind, extend map[string]interface{}) {
			t.Log("join fail", key, kind)
		},
		Finish: func(curRound, maxRound uint32, pMgr *pmgr.PMgr, disband bool, extend map[string]interface{}) {
			t.Log("finish", time.Now().Unix())
		},
	})
	h.RoundBeginPolicy(RoundBeginPolicyAllPlayerReady)
	h.Init()

	go func() {
		h.Msg(&proxy.Msg{
			Kind:    proxy.MsgKindPlayerJoin,
			UserKey: 1,
			Data: &proxy.MsgKindPlayerJoinData{
				Player: player.New(),
			},
		})

		h.Msg(&proxy.Msg{
			Kind:    proxy.MsgKindPlayerReady,
			UserKey: 1,
			Data: &proxy.MsgKindPlayerReadyData{
				Site: 1,
			},
		})

		h.Msg(&proxy.Msg{
			Kind:    proxy.MsgKindPlayerJoin,
			UserKey: 2,
			Data: &proxy.MsgKindPlayerJoinData{
				Player: player.New(),
			},
		})
	}()
	t.Log("run", time.Now().Unix())
	h.Run(false)
}

type EventPlayerJoinFailGameLocationOff struct {
	GameBase
}

func (g *EventPlayerJoinFailGameLocationOff) PlayerMaxNum() int {
	return 1
}

func (g *EventPlayerJoinFailGameLocationOff) DistanceLimit() int {
	return 10
}

func (g *EventPlayerJoinFailGameLocationOff) IpLimit() bool {
	return true
}

func Test_Happy_EventPlayerJoinFail_LocationOff(t *testing.T) {
	h := New(nil, 1, new(EventPlayerJoinFailGameLocationOff), nil)
	h.Heartbeat(3 * time.Second)
	h.Event(&Event{
		PlayerJoinFail: func(key interface{}, kind EventPlayerJoinFailKind, extend map[string]interface{}) {
			t.Log("join fail", key, kind)
		},
		Finish: func(curRound, maxRound uint32, pMgr *pmgr.PMgr, disband bool, extend map[string]interface{}) {
			t.Log("finish", time.Now().Unix())
		},
	})
	h.RoundBeginPolicy(RoundBeginPolicyAllPlayerReady)
	h.Init()

	go func() {
		p := player.New()
		p.SetLocation(&player.Location{
			//Ip: "127.0.0.1",
		})
		h.Msg(&proxy.Msg{
			Kind:    proxy.MsgKindPlayerJoin,
			UserKey: 1,
			Data: &proxy.MsgKindPlayerJoinData{
				Player: p,
			},
		})
	}()
	t.Log("run", time.Now().Unix())
	h.Run(false)
}

type EventPlayerJoinFailGameLocationIpSame struct {
	GameBase
}

func (g *EventPlayerJoinFailGameLocationIpSame) PlayerMaxNum() int {
	return 2
}

func (g *EventPlayerJoinFailGameLocationIpSame) IpLimit() bool {
	return true
}

func Test_Happy_EventPlayerJoinFail_LocationIpSame(t *testing.T) {
	h := New(nil, 1, new(EventPlayerJoinFailGameLocationIpSame), nil)
	h.Heartbeat(3 * time.Second)
	h.Event(&Event{
		PlayerJoinFail: func(key interface{}, kind EventPlayerJoinFailKind, extend map[string]interface{}) {
			t.Log("join fail", key, kind)
		},
		Finish: func(curRound, maxRound uint32, pMgr *pmgr.PMgr, disband bool, extend map[string]interface{}) {
			t.Log("finish", time.Now().Unix())
		},
	})
	h.RoundBeginPolicy(RoundBeginPolicyAllPlayerReady)
	h.Init()

	go func() {
		p := player.New()
		p.SetLocation(&player.Location{
			Ip: "127.0.0.1",
		})
		h.Msg(&proxy.Msg{
			Kind:    proxy.MsgKindPlayerJoin,
			UserKey: 1,
			Data: &proxy.MsgKindPlayerJoinData{
				Player: p,
			},
		})

		p = player.New()
		p.SetLocation(&player.Location{
			Ip: "127.0.0.1",
		})
		h.Msg(&proxy.Msg{
			Kind:    proxy.MsgKindPlayerJoin,
			UserKey: 2,
			Data: &proxy.MsgKindPlayerJoinData{
				Player: p,
			},
		})
	}()
	t.Log("run", time.Now().Unix())
	h.Run(false)
}

type EventPlayerJoinFailGameLocationTooClose struct {
	GameBase
}

func (g *EventPlayerJoinFailGameLocationTooClose) PlayerMaxNum() int {
	return 2
}

func (g *EventPlayerJoinFailGameLocationTooClose) DistanceLimit() int {
	return 2
}

func Test_Happy_EventPlayerJoinFail_LocationTooClose(t *testing.T) {
	h := New(nil, 1, new(EventPlayerJoinFailGameLocationTooClose), nil)
	h.Heartbeat(3 * time.Second)
	h.Event(&Event{
		PlayerJoinFail: func(key interface{}, kind EventPlayerJoinFailKind, extend map[string]interface{}) {
			t.Log("join fail", key, kind)
		},
		Finish: func(curRound, maxRound uint32, pMgr *pmgr.PMgr, disband bool, extend map[string]interface{}) {
			t.Log("finish", time.Now().Unix())
		},
	})
	h.Plugin(&Plugin{
		LocationDistance: func(origin *player.Location, des ...*player.Location) ([]int, error) {
			t.Log("location distance plugin", origin, des)
			ret := make([]int, len(des))
			for i := range ret {
				ret[i] = 1
			}
			return ret, nil
		},
	})
	h.RoundBeginPolicy(RoundBeginPolicyAllPlayerReady)
	h.Init()

	go func() {
		p := player.New()
		p.SetLocation(&player.Location{
			Longitude: "111",
			Latitude:  "1111",
		})
		h.Msg(&proxy.Msg{
			Kind:    proxy.MsgKindPlayerJoin,
			UserKey: 1,
			Data: &proxy.MsgKindPlayerJoinData{
				Player: p,
			},
		})

		p = player.New()
		p.SetLocation(&player.Location{
			Longitude: "222",
			Latitude:  "222",
		})
		h.Msg(&proxy.Msg{
			Kind:    proxy.MsgKindPlayerJoin,
			UserKey: 2,
			Data: &proxy.MsgKindPlayerJoinData{
				Player: p,
			},
		})
	}()
	t.Log("run", time.Now().Unix())
	h.Run(false)
}

func Test_Happy_EventPlayerExit(t *testing.T) {
	h := New(nil, 1, new(GameBase), nil)
	h.Heartbeat(3 * time.Second)
	h.Event(&Event{
		PlayerExit: func(key interface{}, pMgr *pmgr.PMgr, extend map[string]interface{}) {
			t.Log("player exit", key, pMgr.Len())
		},
		Finish: func(curRound, maxRound uint32, pMgr *pmgr.PMgr, disband bool, extend map[string]interface{}) {
			t.Log("finish", time.Now().Unix())
		},
	})
	h.Init()

	go func() {
		h.Msg(&proxy.Msg{
			Kind:    proxy.MsgKindPlayerJoin,
			UserKey: 1,
			Data: &proxy.MsgKindPlayerJoinData{
				Player: player.New(),
			},
		})

		h.Msg(&proxy.Msg{
			Kind:    proxy.MsgKindPlayerExit,
			UserKey: 1,
		})
	}()
	t.Log("run", time.Now().Unix())
	h.Run(false)
}

func Test_Happy_EventPlayerReady(t *testing.T) {
	h := New(nil, 1, new(GameBase), nil)
	h.Heartbeat(3 * time.Second)
	h.Event(&Event{
		PlayerReady: func(key interface{}, pMgr *pmgr.PMgr, extend map[string]interface{}) {
			t.Log("player ready", key, pMgr.Get(key).Ready(), pMgr.Len())
		},
		Finish: func(curRound, maxRound uint32, pMgr *pmgr.PMgr, disband bool, extend map[string]interface{}) {
			t.Log("finish", time.Now().Unix())
		},
	})
	h.Init()

	go func() {
		h.Msg(&proxy.Msg{
			Kind:    proxy.MsgKindPlayerJoin,
			UserKey: 1,
			Data: &proxy.MsgKindPlayerJoinData{
				Player: player.New(),
			},
		})

		h.Msg(&proxy.Msg{
			Kind:    proxy.MsgKindPlayerReady,
			UserKey: 1,
			Data: &proxy.MsgKindPlayerReadyData{
				Site: 1,
			},
		})
	}()
	t.Log("run", time.Now().Unix())
	h.Run(false)
}

func Test_Happy_EventPlayerLine(t *testing.T) {
	h := New(nil, 1, new(GameBase), nil)
	h.Heartbeat(3 * time.Second)
	h.Event(&Event{
		PlayerLine: func(key interface{}, pMgr *pmgr.PMgr, extend map[string]interface{}) {
			t.Log("player line", key, pMgr.Get(key).OfflineTs(), pMgr.Len())
		},
		Finish: func(curRound, maxRound uint32, pMgr *pmgr.PMgr, disband bool, extend map[string]interface{}) {
			t.Log("finish", time.Now().Unix())
		},
	})
	h.Init()

	go func() {
		p := player.New()
		h.Msg(&proxy.Msg{
			Kind:    proxy.MsgKindPlayerJoin,
			UserKey: 1,
			Data: &proxy.MsgKindPlayerJoinData{
				Player: p,
			},
		})

		time.Sleep(1 * time.Second)
		p.Offline(0)
	}()
	t.Log("run", time.Now().Unix())
	h.Run(false)
}

func Test_Happy_EventPlayerOp(t *testing.T) {
	h := New(nil, 1, new(GameBase), nil)
	h.Heartbeat(3 * time.Second)
	h.Event(&Event{
		PlayerOp: func(key interface{}, pMgr *pmgr.PMgr, extend map[string]interface{}) {
			t.Log("player op", key, pMgr.Get(key).Op(), pMgr.Len())
		},
		Finish: func(curRound, maxRound uint32, pMgr *pmgr.PMgr, disband bool, extend map[string]interface{}) {
			t.Log("finish", time.Now().Unix())
		},
	})
	h.Init()

	go func() {
		p := player.New()
		h.Msg(&proxy.Msg{
			Kind:    proxy.MsgKindPlayerJoin,
			UserKey: 1,
			Data: &proxy.MsgKindPlayerJoinData{
				Player: p,
			},
		})

		time.Sleep(1 * time.Second)
		p.SetOp(true)
		time.Sleep(1 * time.Second)
		p.SetOp(false)
	}()
	t.Log("run", time.Now().Unix())
	h.Run(false)
}

func Test_Happy_EventPlayerAuto(t *testing.T) {
	h := New(nil, 1, new(GameBase), nil)
	h.Heartbeat(3 * time.Second)
	h.Event(&Event{
		PlayerAuto: func(key interface{}, pMgr *pmgr.PMgr, extend map[string]interface{}) {
			t.Log("player auto", key, pMgr.Get(key).Auto(), pMgr.Len())
		},
		Finish: func(curRound, maxRound uint32, pMgr *pmgr.PMgr, disband bool, extend map[string]interface{}) {
			t.Log("finish", time.Now().Unix())
		},
	})
	h.Init()

	go func() {
		p := player.New()
		h.Msg(&proxy.Msg{
			Kind:    proxy.MsgKindPlayerJoin,
			UserKey: 1,
			Data: &proxy.MsgKindPlayerJoinData{
				Player: p,
			},
		})

		time.Sleep(1 * time.Second)
		p.SetAuto(true)
		time.Sleep(1 * time.Second)
		p.SetAuto(false)
	}()
	t.Log("run", time.Now().Unix())
	h.Run(false)
}

type EventPlayerSite struct {
	GameBase
}

func (g *EventPlayerSite) PlayerMaxNum() int {
	return 2
}

func Test_Happy_EventPlayerSite(t *testing.T) {
	h := New(nil, 1, new(EventPlayerSite), nil)
	h.Heartbeat(3 * time.Second)
	h.Event(&Event{
		PlayerSite: func(key interface{}, pMgr *pmgr.PMgr, extend map[string]interface{}) {
			t.Log("player site", key, pMgr.Get(key).Site(), pMgr.Len())
		},
		Finish: func(curRound, maxRound uint32, pMgr *pmgr.PMgr, disband bool, extend map[string]interface{}) {
			t.Log("finish", time.Now().Unix())
		},
	})
	h.Init()

	go func() {
		p := player.New()
		h.Msg(&proxy.Msg{
			Kind:    proxy.MsgKindPlayerJoin,
			UserKey: 1,
			Data: &proxy.MsgKindPlayerJoinData{
				Player: p,
			},
		})
		h.Msg(&proxy.Msg{
			Kind:    proxy.MsgKindPlayerReady,
			UserKey: 1,
			Data: &proxy.MsgKindPlayerReadyData{
				Site: 2,
			},
		})
	}()
	t.Log("run", time.Now().Unix())
	h.Run(false)
}

func Test_Happy_EventPlayerScore(t *testing.T) {
	h := New(nil, 1, new(GameBase), nil)
	h.Heartbeat(3 * time.Second)
	h.Event(&Event{
		PlayerScore: func(key interface{}, pMgr *pmgr.PMgr, extend map[string]interface{}) {
			t.Log("player score", key, pMgr.Get(key).Score(0), pMgr.Len())
		},
		Finish: func(curRound, maxRound uint32, pMgr *pmgr.PMgr, disband bool, extend map[string]interface{}) {
			t.Log("finish", time.Now().Unix())
		},
	})
	h.Init()

	go func() {
		p := player.New()
		h.Msg(&proxy.Msg{
			Kind:    proxy.MsgKindPlayerJoin,
			UserKey: 1,
			Data: &proxy.MsgKindPlayerJoinData{
				Player: p,
			},
		})
		time.Sleep(time.Second)
		p.AddScore(10)
	}()
	t.Log("run", time.Now().Unix())
	h.Run(false)
}

type EventCostGame struct {
	GameBase
}

func (g *EventCostGame) PlayerMaxNum() int {
	return 1
}

func Test_Happy_EventCost_Join(t *testing.T) {
	h := New(nil, 1, new(EventCostGame), nil)
	h.Heartbeat(3 * time.Second)
	h.Event(&Event{
		Cost: func(mode CostMode, back bool, pMgr *pmgr.PMgr, extend map[string]interface{}) {
			t.Log("cost", mode, back, pMgr.Len())
		},
		Finish: func(curRound, maxRound uint32, pMgr *pmgr.PMgr, disband bool, extend map[string]interface{}) {
			t.Log("finish", time.Now().Unix())
		},
	})
	h.Cost(CostModeJoin)
	h.Init()

	go func() {
		p := player.New()
		h.Msg(&proxy.Msg{
			Kind:    proxy.MsgKindPlayerJoin,
			UserKey: 1,
			Data: &proxy.MsgKindPlayerJoinData{
				Player: p,
			},
		})
	}()
	t.Log("run", time.Now().Unix())
	h.Run(false)
}

func Test_Happy_EventCost_FirstRoundBegin(t *testing.T) {
	h := New(nil, 1, new(EventCostGame), nil)
	h.Heartbeat(3 * time.Second)
	h.Event(&Event{
		Cost: func(mode CostMode, back bool, pMgr *pmgr.PMgr, extend map[string]interface{}) {
			t.Log("cost", mode, back, pMgr.Len())
		},
		Finish: func(curRound, maxRound uint32, pMgr *pmgr.PMgr, disband bool, extend map[string]interface{}) {
			t.Log("finish", time.Now().Unix())
		},
	})
	h.Cost(CostModeFirstRoundBegin)
	h.RoundBeginPolicy(RoundBeginPolicyFullPlayer)
	h.Init()

	go func() {
		p := player.New()
		h.Msg(&proxy.Msg{
			Kind:    proxy.MsgKindPlayerJoin,
			UserKey: 1,
			Data: &proxy.MsgKindPlayerJoinData{
				Player: p,
			},
		})
	}()
	t.Log("run", time.Now().Unix())
	h.Run(false)
}

type EventCostFirstRoundEndGame struct {
	GameBase
}

func (g *EventCostFirstRoundEndGame) PlayerMaxNum() int {
	return 1
}

func (g *EventCostFirstRoundEndGame) Begin(quick bool) {
	g.GameEnd()
}

func Test_Happy_EventCost_FirstRoundEnd(t *testing.T) {
	h := New(nil, 1, new(EventCostFirstRoundEndGame), nil)
	h.Heartbeat(3 * time.Second)
	h.Event(&Event{
		Cost: func(mode CostMode, back bool, pMgr *pmgr.PMgr, extend map[string]interface{}) {
			t.Log("cost", mode, back, pMgr.Len())
		},
		Finish: func(curRound, maxRound uint32, pMgr *pmgr.PMgr, disband bool, extend map[string]interface{}) {
			t.Log("finish", time.Now().Unix())
		},
	})
	h.Cost(CostModeFirstRoundEnd)
	h.RoundBeginPolicy(RoundBeginPolicyFullPlayer)
	h.Init()

	go func() {
		p := player.New()
		h.Msg(&proxy.Msg{
			Kind:    proxy.MsgKindPlayerJoin,
			UserKey: 1,
			Data: &proxy.MsgKindPlayerJoinData{
				Player: p,
			},
		})
	}()
	t.Log("run", time.Now().Unix())
	h.Run(false)
}

type EventCostRoundBeginGame struct {
	GameBase
}

func (g *EventCostRoundBeginGame) PlayerMaxNum() int {
	return 1
}

func (g *EventCostRoundBeginGame) Begin(quick bool) {
	g.GameEnd()
}

func Test_Happy_EventCost_RoundBegin(t *testing.T) {
	h := New(nil, 2, new(EventCostRoundBeginGame), nil)
	h.Heartbeat(3 * time.Second)
	h.Event(&Event{
		RoundBegin: func(curRound, maxRound uint32, pMgr *pmgr.PMgr, extend map[string]interface{}) {
			t.Log("round begin", curRound, maxRound)
		},
		RoundEnd: func(curRound, maxRound uint32, pMgr *pmgr.PMgr, extend map[string]interface{}) {
			t.Log("round end", curRound, maxRound)
			h.Msg(&proxy.Msg{
				Kind:    proxy.MsgKindPlayerReady,
				UserKey: 1,
				Data: &proxy.MsgKindPlayerReadyData{
					Site: 1,
				},
			})
		},
		Cost: func(mode CostMode, back bool, pMgr *pmgr.PMgr, extend map[string]interface{}) {
			t.Log("cost", mode, back, pMgr.Len())
		},
		Finish: func(curRound, maxRound uint32, pMgr *pmgr.PMgr, disband bool, extend map[string]interface{}) {
			t.Log("finish", time.Now().Unix())
		},
	})
	h.Cost(CostModeRoundBegin)
	h.RoundBeginPolicy(RoundBeginPolicyFullPlayer)
	h.Init()

	go func() {
		p := player.New()
		h.Msg(&proxy.Msg{
			Kind:    proxy.MsgKindPlayerJoin,
			UserKey: 1,
			Data: &proxy.MsgKindPlayerJoinData{
				Player: p,
			},
		})
	}()
	t.Log("run", time.Now().Unix())
	h.Run(false)
}

type EventCostRoundEndGame struct {
	GameBase
}

func (g *EventCostRoundEndGame) PlayerMaxNum() int {
	return 1
}

func (g *EventCostRoundEndGame) Begin(quick bool) {
	g.GameEnd()
}

func Test_Happy_EventCost_RoundEnd(t *testing.T) {
	h := New(nil, 2, new(EventCostRoundEndGame), nil)
	h.Heartbeat(3 * time.Second)
	h.Event(&Event{
		RoundBegin: func(curRound, maxRound uint32, pMgr *pmgr.PMgr, extend map[string]interface{}) {
			t.Log("round begin", curRound, maxRound)
		},
		RoundEnd: func(curRound, maxRound uint32, pMgr *pmgr.PMgr, extend map[string]interface{}) {
			t.Log("round end", curRound, maxRound)
			h.Msg(&proxy.Msg{
				Kind:    proxy.MsgKindPlayerReady,
				UserKey: 1,
				Data: &proxy.MsgKindPlayerReadyData{
					Site: 1,
				},
			})
		},
		Cost: func(mode CostMode, back bool, pMgr *pmgr.PMgr, extend map[string]interface{}) {
			t.Log("cost", mode, back, pMgr.Len())
		},
		Finish: func(curRound, maxRound uint32, pMgr *pmgr.PMgr, disband bool, extend map[string]interface{}) {
			t.Log("finish", time.Now().Unix())
		},
	})
	h.Cost(CostModeRoundEnd)
	h.RoundBeginPolicy(RoundBeginPolicyFullPlayer)
	h.Init()

	go func() {
		p := player.New()
		h.Msg(&proxy.Msg{
			Kind:    proxy.MsgKindPlayerJoin,
			UserKey: 1,
			Data: &proxy.MsgKindPlayerJoinData{
				Player: p,
			},
		})
	}()
	t.Log("run", time.Now().Unix())
	h.Run(false)
}

type EventCostFinishGame struct {
	GameBase
}

func (g *EventCostFinishGame) PlayerMaxNum() int {
	return 1
}

func (g *EventCostFinishGame) Begin(quick bool) {
	g.GameEnd()
}

func Test_Happy_EventCost_Finish(t *testing.T) {
	h := New(nil, 2, new(EventCostFinishGame), nil)
	h.Heartbeat(3 * time.Second)
	h.Event(&Event{
		RoundBegin: func(curRound, maxRound uint32, pMgr *pmgr.PMgr, extend map[string]interface{}) {
			t.Log("round begin", curRound, maxRound)
		},
		RoundEnd: func(curRound, maxRound uint32, pMgr *pmgr.PMgr, extend map[string]interface{}) {
			t.Log("round end", curRound, maxRound)
			h.Msg(&proxy.Msg{
				Kind:    proxy.MsgKindPlayerReady,
				UserKey: 1,
				Data: &proxy.MsgKindPlayerReadyData{
					Site: 1,
				},
			})
		},
		Cost: func(mode CostMode, back bool, pMgr *pmgr.PMgr, extend map[string]interface{}) {
			t.Log("cost", mode, back, pMgr.Len())
		},
		Finish: func(curRound, maxRound uint32, pMgr *pmgr.PMgr, disband bool, extend map[string]interface{}) {
			t.Log("finish", time.Now().Unix())
		},
	})
	h.Cost(CostModeFinish)
	h.RoundBeginPolicy(RoundBeginPolicyFullPlayer)
	h.Init()

	go func() {
		p := player.New()
		h.Msg(&proxy.Msg{
			Kind:    proxy.MsgKindPlayerJoin,
			UserKey: 1,
			Data: &proxy.MsgKindPlayerJoinData{
				Player: p,
			},
		})
	}()
	t.Log("run", time.Now().Unix())
	h.Run(false)
}

type EventDisbandGame struct {
	GameBase
}

func (g *EventDisbandGame) PlayerMaxNum() int {
	return 2
}

func (g *EventDisbandGame) DisbandTs() time.Duration {
	return 2 * time.Second
}

func Test_Happy_EventDisband(t *testing.T) {
	h := New(nil, 1, new(EventDisbandGame), nil)
	h.Heartbeat(3 * time.Second)
	h.Event(&Event{
		PlayerJoinSuccess: func(key interface{}, pMgr *pmgr.PMgr, alreadyExist bool, extend map[string]interface{}) {
			t.Log("player join success", key, pMgr.Len())
		},
		PlayerJoinFail: func(key interface{}, kind EventPlayerJoinFailKind, extend map[string]interface{}) {
			t.Log("player join fail", key, kind)
		},
		RoundBegin: func(curRound, maxRound uint32, pMgr *pmgr.PMgr, extend map[string]interface{}) {
			t.Log("round begin", curRound, maxRound)
		},
		RoundEnd: func(curRound, maxRound uint32, pMgr *pmgr.PMgr, extend map[string]interface{}) {
			t.Log("round end", curRound, maxRound)
		},
		Finish: func(curRound, maxRound uint32, pMgr *pmgr.PMgr, disband bool, extend map[string]interface{}) {
			t.Log("finish", time.Now().Unix())
		},
		DisbandAgree: func(ts time.Duration, deadlineTs int64, userKey interface{}, pMgr *pmgr.PMgr, op map[interface{}]bool, extend map[string]interface{}) {
			t.Log("disband agree", ts, deadlineTs, userKey, pMgr.Len(), op)
		},
		DisbandReject: func(ts time.Duration, deadlineTs int64, userKey interface{}, pMgr *pmgr.PMgr, op map[interface{}]bool, extend map[string]interface{}) {
			t.Log("disband reject", ts, deadlineTs, userKey, pMgr.Len(), op)
		},
		DisbandPass: func(deadlineTs int64, pMgr *pmgr.PMgr, op map[interface{}]bool, extend map[string]interface{}) {
			t.Log("disband pass", deadlineTs, pMgr.Len(), op)
		},
		DisbandFail: func(deadlineTs int64, pMgr *pmgr.PMgr, op map[interface{}]bool, extend map[string]interface{}) {
			t.Log("disband fail", deadlineTs, pMgr.Len(), op)
		},
	})
	h.RoundBeginPolicy(RoundBeginPolicyFullPlayer)
	h.Init()

	go func() {
		h.Msg(&proxy.Msg{
			Kind:    proxy.MsgKindPlayerJoin,
			UserKey: 1,
			Data: &proxy.MsgKindPlayerJoinData{
				Player: player.New(),
			},
		})

		h.Msg(&proxy.Msg{
			Kind:    proxy.MsgKindPlayerJoin,
			UserKey: 2,
			Data: &proxy.MsgKindPlayerJoinData{
				Player: player.New(),
			},
		})

		h.Msg(&proxy.Msg{
			Kind:    proxy.MsgKindDisband,
			UserKey: 1,
		})

		time.Sleep(time.Second)
		h.Msg(&proxy.Msg{
			Kind:    proxy.MsgKindDisbandReject,
			UserKey: 2,
		})
	}()
	t.Log("run", time.Now().Unix())
	h.Run(false)
}

type EventQuickGame struct {
	GameBase
}

func (g *EventQuickGame) PlayerMaxNum() int {
	return 3
}

func (g *EventQuickGame) Quick(num int) bool {
	return true
}

func (g *EventQuickGame) QuickTs() time.Duration {
	return 1 * time.Second
}

func Test_Happy_EventQuick(t *testing.T) {
	h := New(nil, 1, new(EventQuickGame), nil)
	h.Heartbeat(3 * time.Second)
	h.Event(&Event{
		PlayerJoinSuccess: func(key interface{}, pMgr *pmgr.PMgr, alreadyExist bool, extend map[string]interface{}) {
			t.Log("player join success", key, pMgr.Len())
		},
		PlayerJoinFail: func(key interface{}, kind EventPlayerJoinFailKind, extend map[string]interface{}) {
			t.Log("player join fail", key, kind)
		},
		RoundBegin: func(curRound, maxRound uint32, pMgr *pmgr.PMgr, extend map[string]interface{}) {
			t.Log("round begin", curRound, maxRound)
		},
		RoundEnd: func(curRound, maxRound uint32, pMgr *pmgr.PMgr, extend map[string]interface{}) {
			t.Log("round end", curRound, maxRound)
		},
		QuickAgree: func(ts time.Duration, deadlineTs int64, userKey interface{}, pMgr *pmgr.PMgr, op map[interface{}]bool, extend map[string]interface{}) {
			t.Log("quick agree", ts, deadlineTs, userKey, pMgr.Len(), op)
		},
		QuickReject: func(ts time.Duration, deadlineTs int64, userKey interface{}, pMgr *pmgr.PMgr, op map[interface{}]bool, extend map[string]interface{}) {
			t.Log("quick reject", ts, deadlineTs, userKey, pMgr.Len(), op)
		},
		QuickPass: func(deadlineTs int64, pMgr *pmgr.PMgr, op map[interface{}]bool, extend map[string]interface{}) {
			t.Log("quick pass", deadlineTs, pMgr.Len(), op)
		},
		QuickFail: func(deadlineTs int64, pMgr *pmgr.PMgr, op map[interface{}]bool, extend map[string]interface{}) {
			t.Log("quick fail", deadlineTs, pMgr.Len(), op)
		},
		Finish: func(curRound, maxRound uint32, pMgr *pmgr.PMgr, disband bool, extend map[string]interface{}) {
			t.Log("finish", time.Now().Unix())
		},
	})
	h.Init()

	go func() {
		h.Msg(&proxy.Msg{
			Kind:    proxy.MsgKindPlayerJoin,
			UserKey: 1,
			Data: &proxy.MsgKindPlayerJoinData{
				Player: player.New(),
			},
		})

		h.Msg(&proxy.Msg{
			Kind:    proxy.MsgKindPlayerJoin,
			UserKey: 2,
			Data: &proxy.MsgKindPlayerJoinData{
				Player: player.New(),
			},
		})

		h.Msg(&proxy.Msg{
			Kind:    proxy.MsgKindQuick,
			UserKey: 1,
		})

		time.Sleep(500 * time.Millisecond)
		h.Msg(&proxy.Msg{
			Kind:    proxy.MsgKindQuickReject,
			UserKey: 2,
		})
	}()
	t.Log("run", time.Now().Unix())
	h.Run(false)
}
