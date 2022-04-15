package happy

import "github.com/cnlisea/happy/pmgr/player"

type Plugin struct {
	PlayerExitDisband   func(userKey interface{}, ownerUserKey interface{}, extend map[string]interface{}) bool
	DisbandMinAgreeNum  func(num int) int
	DisbandDeadlinePass func(extend map[string]interface{}) bool
	QuickMinAgreeNum    func(num int) int
	QuickDeadlinePass   func(extend map[string]interface{}) bool
	LocationDistance    func(origin *player.Location, des ...*player.Location) ([]int, error)
}

func (h *_Happy) Plugin(p *Plugin) {
	h.plugin = p
}
