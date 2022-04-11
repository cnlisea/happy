package happy

import "context"

type PlayerMsg interface {
	Send(ctx context.Context, userKey []interface{}, data ...interface{}) error
	ReConn(ctx context.Context, userKey []interface{}) error
}

func (h *Happy) PlayerMsg(msg PlayerMsg) {
	h.playerMsg = msg
}
