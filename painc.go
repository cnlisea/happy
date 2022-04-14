package happy

import "errors"

var (
	PanicDoneExit = errors.New("done exit")
	PanicGameEnd  = errors.New("game end")
)
