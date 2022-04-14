package happy

import (
	"github.com/cnlisea/happy/proxy"
)

func (h *Happy) MsgHandler(msg *proxy.Msg) {
	if msg == nil {
		return
	}

	switch msg.Kind {
	case proxy.MsgKindPlayerJoin:
		if data, ok := msg.Data.(*proxy.MsgKindPlayerJoinData); ok && data != nil {
			h.MsgPlayerJoinHandler(msg.UserKey, data.Player)
		}
	case proxy.MsgKindPlayerExit:
		h.MsgPlayerExitHandler(msg.UserKey)
	case proxy.MsgKindPlayerReady:
		if data, ok := msg.Data.(*proxy.MsgKindPlayerReadyData); ok && data != nil {
			h.MsgPlayerReadyHandler(msg.UserKey, data.Site)
		}
	case proxy.MsgKindDisband:
		// 申请解散
		h.MsgDisbandHandler(msg.UserKey)
	case proxy.MsgKindDisbandReject:
		// 拒绝解散
		h.MsgDisbandRejectHandler(msg.UserKey)
	case proxy.MsgKindQuick:
		// 申请少人开局
		h.MsgQuickHandler(msg.UserKey)
	case proxy.MsgKindQuickReject:
		// 拒绝少人开局
		h.MsgQuickRejectHandler(msg.UserKey)
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
		h.MsgGameHandler(msg.UserKey, msg.Data)
	case proxy.MsgKindByUser:
		h.MsgByUserHandler(msg.UserKey, msg.Data)
	}
}
