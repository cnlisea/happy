package pmgr

import "container/list"

type PMgr struct {
	players *list.List          // 玩家列表
	watch   map[WatchKind]Watch // 监听
}

func New() *PMgr {
	return &PMgr{
		players: list.New(),
		watch:   make(map[WatchKind]Watch),
	}
}
