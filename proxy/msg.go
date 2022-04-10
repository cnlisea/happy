package proxy

type MsgKind int

const (
	MsgKindPlayerJoin    MsgKind = iota // 玩家加入
	MsgKindPlayerExit                   // 玩家退出
	MsgKindDisband                      // 申请解散
	MsgKindDisbandReject                // 拒绝解散
	MsgKindDisbandIdle                  // 解散房间
	MsgKindDisbandForce                 // 强制解散房间
	MsgKindGame                         // 游戏操作
	MsgKindByUser                       // 自定义
)

type Msg struct {
	Kind    MsgKind
	UserKey interface{}
	Data    interface{}
}
