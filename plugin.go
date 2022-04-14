package happy

type Plugin struct {
	PlayerExitDisband   func(userKey interface{}, ownerUserKey interface{}, extend map[string]interface{}) bool
	DisbandMinAgreeNum  func(num int) int
	DisbandDeadlinePass func(extend map[string]interface{}) bool
	QuickMinAgreeNum    func(num int) int
	QuickDeadlinePass   func(extend map[string]interface{}) bool
}
