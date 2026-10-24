package happy

func (h *_Happy) Extend(key string) interface{} {
	return h.extend[key]
}

func (h *_Happy) SetExtend(key string, val interface{}) {
	h.extend[key] = val
}

func (h *_Happy) CopyExtend(ext map[string]any) map[string]any {
	if ext == nil {
		return h.extend
	}

	for k, v := range h.extend {
		ext[k] = v
	}

	return ext
}
