package happy

import (
	"github.com/cnlisea/happy/pmgr"
	"github.com/cnlisea/happy/proxy"
)

func (h *_Happy) MsgByUser(f func(userKey interface{}, data interface{}, delay proxy.Delay, curRound, maxRound uint32, pMgr *pmgr.PMgr, extend map[string]interface{})) {
	h.byUserHandler = f
}

func (h *_Happy) MsgByUserHandler(userKey interface{}, data interface{}) {
	if h.byUserHandler == nil {
		return
	}

	h.byUserHandler(userKey, data, h.delay, h.curRound, h.maxRound, h.pMgr, h.extend)
}
