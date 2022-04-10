package happy

type Plugin struct {
	PlayerExitDisband func(userKey interface{}, ownerUserKey interface{}, extend map[string]interface{}) bool
}
