package happy

import (
	"github.com/cnlisea/happy/pmgr"
	"github.com/cnlisea/happy/proxy"
	"testing"
	"time"
)

func Test_Happy_MsgByUser(t *testing.T) {
	h := New(nil, 1, new(GameBase), nil)
	h.Heartbeat(3 * time.Second)
	h.Event(&Event{
		Finish: func(h Happy, curRound, maxRound uint32, pMgr *pmgr.PMgr, disband bool, extend map[string]interface{}) {
			t.Log("finish", time.Now().Unix())
		},
	})
	h.MsgByUser(func(userKey interface{}, data interface{}, delay proxy.Delay, curRound, maxRound uint32, pMgr *pmgr.PMgr, extend map[string]interface{}) {
		t.Log("userKey", userKey, "data", data)
	})
	h.Init()

	go func() {
		h.Msg(&proxy.Msg{
			Kind:    proxy.MsgKindByUser,
			UserKey: 1,
			Data:    "by user",
		})
	}()
	t.Log("run", time.Now().Unix())
	h.Run(false)
}
