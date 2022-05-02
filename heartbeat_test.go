package happy

import (
	"testing"
	"time"

	"github.com/cnlisea/happy/pmgr"
	"github.com/cnlisea/happy/proxy"
)

func Test_Happy_Heartbeat(t *testing.T) {
	h := New(nil, 1, new(GameBase), nil)
	h.Heartbeat(3 * time.Second)
	h.Event(&Event{
		Finish: func(h Happy, curRound, maxRound uint32, pMgr *pmgr.PMgr, disband bool, extend map[string]interface{}) {
			t.Log("finish", time.Now().Unix())
		},
	})
	h.Init()

	t.Log("run", time.Now().Unix())
	h.Run(false)
}

func Test_Happy_HeartbeatMsg(t *testing.T) {
	h := New(nil, 1, new(GameBase), nil)
	h.Heartbeat(3 * time.Second)
	h.Event(&Event{
		Finish: func(h Happy, curRound, maxRound uint32, pMgr *pmgr.PMgr, disband bool, extend map[string]interface{}) {
			t.Log(time.Now().Unix())
		},
	})
	h.Init()
	t.Log(time.Now().Unix())

	go func() {
		for i := 0; i < 2; i++ {
			time.Sleep(2900 * time.Millisecond)
			h.Msg(&proxy.Msg{
				Kind: proxy.MsgKindGame,
			})
			t.Log("send msg ts:", time.Now().Unix())
		}
	}()
	h.Run(false)
}
