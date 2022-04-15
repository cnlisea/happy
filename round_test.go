package happy

import (
	"github.com/cnlisea/happy/pmgr"
	"github.com/cnlisea/happy/pmgr/player"
	"github.com/cnlisea/happy/proxy"
	"testing"
	"time"
)

type RoundBeginPolicyFullPlayerGame struct {
	GameBase
}

func (g *RoundBeginPolicyFullPlayerGame) PlayerMaxNum() int {
	return 2
}

func Test_Happy_RoundBeginPolicyFullPlayer(t *testing.T) {
	h := New(nil, 1, new(RoundBeginPolicyFullPlayerGame), nil)
	h.Heartbeat(3 * time.Second)
	h.Event(&Event{
		RoundBegin: func(curRound, maxRound uint32, pMgr *pmgr.PMgr, extend map[string]interface{}) {
			t.Log("round begin", curRound, maxRound, pMgr.Len())
		},
		PlayerJoinSuccess: func(key interface{}, pMgr *pmgr.PMgr, alreadyExist bool, extend map[string]interface{}) {
			t.Log("user join", key, pMgr.Get(key))
		},
		Finish: func(curRound, maxRound uint32, pMgr *pmgr.PMgr, disband bool, extend map[string]interface{}) {
			t.Log("finish", time.Now().Unix())
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
	}()
	t.Log("run", time.Now().Unix())
	h.Run(false)
}

func Test_Happy_RoundBeginPolicyAllPlayerReady(t *testing.T) {
	h := New(nil, 1, new(RoundBeginPolicyFullPlayerGame), nil)
	h.Heartbeat(3 * time.Second)
	h.Event(&Event{
		RoundBegin: func(curRound, maxRound uint32, pMgr *pmgr.PMgr, extend map[string]interface{}) {
			t.Log("round begin", curRound, maxRound, pMgr.Len())
		},
		PlayerJoinSuccess: func(key interface{}, pMgr *pmgr.PMgr, alreadyExist bool, extend map[string]interface{}) {
			t.Log("user join", key, pMgr.Get(key))
		},
		PlayerReady: func(key interface{}, pMgr *pmgr.PMgr, extend map[string]interface{}) {
			t.Log("user ready", key, pMgr.Get(key).Ready())
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
