package pmgr

import "github.com/cnlisea/happy/pmgr/player"

type WatchKind int

const (
	WatchKindLine        WatchKind = iota // 在线/离线
	WatchKindReady                        // 准备
	WatchKindOp                           // 操作
	WatchKindView                         // 观战
	WatchKindAuto                         // 托管
	WatchKindAutoReady                    // 托管准备
	WatchKindAutoOp                       // 托管操作
	WatchKindAutoOpReady                  // 托管操作准备
	WatchKindLocation                     // 定位
	WatchKindScore                        // 分数
	WatchKindSite                         // 位置
	WatchKindExtend                       // 扩展
)

type Watch func(key interface{}, p *player.Player, ext map[string]any)

func (pm *PMgr) Watch(kind WatchKind, watch Watch) {
	pm.watch[kind] = watch
}
