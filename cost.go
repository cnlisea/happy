package happy

type CostMode int

const (
	CostModeFree            CostMode = iota // 免费
	CostModeJoin                            // 进入时
	CostModeFirstRoundBegin                 // 首局开始
	CostModeFirstRoundEnd                   // 首局结束
	CostModeRoundBegin                      // 每局开始
	CostModeRoundEnd                        // 每局结束
	CostModeFinish                          // 最后结束
)

func (h *_Happy) Cost(mode CostMode) {
	h.costMode = mode
}
