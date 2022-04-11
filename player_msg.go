package happy

import (
	"github.com/cnlisea/happy/proxy"
)

func (h *Happy) PlayerMsg(msg proxy.PlayerMsg) {
	h.playerMsg = msg
}
