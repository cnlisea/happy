package happy

import (
	"time"
)

func (h *Happy) DelayFunc(delayTs time.Duration, f func(args interface{}), args interface{}) {
	h.delay.Add(delayTs, func(ts int64, args interface{}) {
		f(args)
	}, args)
}

type _DelayMsg struct {
	UserKey []interface{}
	Data    []interface{}
}

func (h *Happy) DelayMsg(delayTs time.Duration, userKey []interface{}, data ...interface{}) {
	if h.playerMsg == nil || userKey == nil || data == nil || len(userKey) == 0 || len(data) == 0 {
		return
	}

	h.delay.Add(delayTs, func(ts int64, args interface{}) {
		var (
			delayMsg = args.(*_DelayMsg)
			err      error
		)
		if err = h.playerMsg.Send(h.ctx, delayMsg.UserKey, delayMsg.Data...); err != nil {
			// TODO log
		}
	}, &_DelayMsg{
		UserKey: userKey,
		Data:    data,
	})
}

func (h *Happy) DelayMsgClean(userKey interface{}) {
	if userKey == nil {
		return
	}

	var (
		msg        *_DelayMsg
		ok         bool
		userKeyLen int
		num        int
		i          int
	)
	h.delay.Del(func(ts int64, args interface{}) bool {
		if msg, ok = args.(*_DelayMsg); !ok {
			return false
		}

		userKeyLen = len(msg.UserKey)
		num = userKeyLen
		for i = 0; i < userKeyLen; i++ {
			if msg.UserKey[i] == userKey {
				msg.UserKey[i] = 0
				num--
				break
			}
		}

		if num > 0 && num != userKeyLen {
			copy(msg.UserKey[i:], msg.UserKey[i+1:])
			msg.UserKey = msg.UserKey[:num]
		}

		return num == 0
	})
}
