package player

import "container/list"

type Player struct {
	offlineTs int64      // 离线时间戳
	location  *Location  // 地址信息
	state     uint32     // 玩家状态
	score     *list.List // 分数

	extend map[string]interface{} // 扩展信息

	watch map[string]func(*Player)
}

func New() *Player {
	return &Player{
		extend: make(map[string]interface{}),
		watch:  make(map[string]func(*Player)),
	}
}
