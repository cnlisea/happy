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
