package proxy

import "github.com/cnlisea/happy/pmgr/player"

type MsgKind int

const (
	MsgKindPlayerJoin    MsgKind = iota // 玩家加入
	MsgKindPlayerExit                   // 玩家退出
	MsgKindPlayerReady                  // 玩家准备
	MsgKindDisband                      // 申请解散
	MsgKindDisbandReject                // 拒绝解散
	MsgKindQuick                        // 申请少人开局
	MsgKindQuickReject                  // 拒绝少人开局
	MsgKindDisbandIdle                  // 解散房间
	MsgKindDisbandForce                 // 强制解散房间
	MsgKindGame                         // 游戏操作
	MsgKindByUser                       // 自定义
)

type MsgKindPlayerJoinData struct {
	Player *player.Player
}

type MsgKindPlayerReadyData struct {
	Site uint32
}

type Msg struct {
	Kind    MsgKind
	UserKey interface{}
	Data    interface{}
}
