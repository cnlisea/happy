package happy

func (h *Happy) Extend(key string) interface{} {
	return h.extend[key]
}

func (h *Happy) SetExtend(key string, val interface{}) {
	h.extend[key] = val
}
