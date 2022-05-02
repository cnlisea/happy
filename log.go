package happy

import (
	"github.com/cnlisea/happy/log"
)

func (h *_Happy) LogSetting(path string, level log.Level) error {
	instance, err := log.NewLogger(log.EncodingJson, path, level, 0)
	if err != nil {
		return err
	}

	if h.log != nil {
		h.log.Close()
	}
	h.log = instance
	return nil
}

func (h *_Happy) Log() *log.Logger {
	return h.log
}
