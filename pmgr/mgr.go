package pmgr

import "container/list"

type PMgr struct {
	players *list.List // 玩家列表
}

func New() *PMgr {
	return &PMgr{
		players: list.New(),
	}
}
