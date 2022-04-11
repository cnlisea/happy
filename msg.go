package happy

import (
	"github.com/cnlisea/happy/pmgr/player"
	"github.com/cnlisea/happy/proxy"
)

func (h *Happy) MsgHandler(msg *proxy.Msg) {
	if msg == nil {
		return
	}

	switch msg.Kind {
	case proxy.MsgKindPlayerJoin:
		if p, ok := msg.Data.(*player.Player); ok {
			h.MsgPlayerJoinHandler(msg.UserKey, p)
		}
	case proxy.MsgKindPlayerExit:
		h.MsgPlayerExitHandler(msg.UserKey)
	case proxy.MsgKindDisband:
		// 申请解散
		h.MsgDisbandHandler(msg.UserKey)
	case proxy.MsgKindDisbandReject:
		// 拒绝解散
		h.MsgDisbandRejectHandler(msg.UserKey)
	case proxy.MsgKindDisbandIdle:
		// 解散房间
		if h.curRound > 0 {
			break
		}
		h.Finish(true)
	case proxy.MsgKindDisbandForce:
		// 强制解散房间
		h.Finish(true)
	case proxy.MsgKindGame:
	case proxy.MsgKindByUser:
	}
}
