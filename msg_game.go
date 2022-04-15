package happy

import (
	"runtime"
)

func (h *_Happy) MsgGameHandler(userKey interface{}, data interface{}) {
	p := h.pMgr.Get(userKey)
	if p == nil {
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

	h.game.Msg(userKey, data)
}
