package happy

import (
	"runtime"

	"github.com/cnlisea/happy/proxy"
)

func (h *_Happy) Msg(msg *proxy.Msg) {
	if msg != nil && h.msgChan != nil {
		h.msgChan <- msg
	}
}

func (h *_Happy) MsgHandler(msg *proxy.Msg) {
	if msg == nil {
		return
	}

	defer func() {
		err := recover()
		switch err {
		case nil:
		case PanicGameEnd:
			h.RoundEnd()
		case PanicDoneExit:
			panic(err)
		default:
			var buf = make([]byte, 4096)
			n := runtime.Stack(buf, false)
			buf = buf[:n]
			// TODO log
		}
	}()

	if h.heartbeat != nil {
		h.heartbeat.Active(0)
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
