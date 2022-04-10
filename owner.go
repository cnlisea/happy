package happy

func (h *Happy) Owner(userKey interface{}) {
	h.ownerUserKey = userKey
}
