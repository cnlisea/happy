package pmgr

import "github.com/cnlisea/happy/pmgr/player"

type WatchKind int

const (
	WatchKindLine     WatchKind = iota // 在线/离线
	WatchKindReady                     // 准备
	WatchKindOp                        // 操作
	WatchKindView                      // 观战
	WatchKindAuto                      // 托管
	WatchKindLocation                  // 定位
	WatchKindScore                     // 分数
	WatchKindExtend                    // 扩展
)

type Watch func(key interface{}, p *player.Player)

func (pm *PMgr) Watch(kind WatchKind, watch Watch) {
	pm.watch[kind] = watch
}
