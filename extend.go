package happy

func (h *_Happy) Extend(key string) interface{} {
	return h.extend[key]
}

func (h *_Happy) SetExtend(key string, val interface{}) {
	h.extend[key] = val
}
