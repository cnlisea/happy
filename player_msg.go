package happy

import (
	"github.com/cnlisea/happy/proxy"
)

func (h *_Happy) PlayerMsg(msg proxy.PlayerMsg) {
	h.playerMsg = msg
}
