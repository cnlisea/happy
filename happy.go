package happy

import (
	"context"
	"github.com/cnlisea/happy/auto"
	"github.com/cnlisea/happy/delay"
	"github.com/cnlisea/happy/heartbeat"
	"github.com/cnlisea/happy/pmgr"
	"github.com/cnlisea/happy/pmgr/player"
	"github.com/cnlisea/happy/proxy"
	"github.com/cnlisea/happy/vote"
	"time"
)

type Happy struct {
	ctx       context.Context
	delay     *delay.Delay
	heartbeat *heartbeat.Heartbeat
	pMgr      *pmgr.PMgr
	extend    map[string]interface{}

	msgChan chan *proxy.Msg

	event  *Event
	plugin *Plugin

	playerMsg proxy.PlayerMsg

	costMode CostMode

	roundBeginPolicy RoundBeginPolicy

	auto *auto.Auto

	disbandVote *vote.Vote
	vote        *vote.Vote

	ownerUserKey interface{}

	game               proxy.Game
	begin              bool   // 开始状态
	curRound, maxRound uint32 // 局数
}

func New(ctx context.Context, game proxy.Game) *Happy {
	if ctx == nil {
		ctx = context.Background()
	}
	return &Happy{
		ctx:   ctx,
		delay: delay.New(),
		pMgr:  pmgr.New(),
		game:  game,
	}
}

func (h *Happy) Init() error {
	h.auto = auto.New(h.delay, func(op bool) time.Duration {
		a := h.game.Auto()
		if a == nil {
			return 0
		}

		ts := a.AutoTs
		if op {
			ts = a.ReadyOpTs
			if h.begin {
				ts = a.OpTs
			}
		}
		return ts
	})
	h.auto.Callback(&auto.Callback{
		Auto: func(key interface{}) {
			p := h.pMgr.Get(key)
			if p == nil {
				return
			}

			if p.Auto() {
				return
			}

			p.SetAuto(true)
			if h.event != nil && h.event.PlayerAuto != nil {
				h.event.PlayerAuto(key, h.pMgr, h.extend)
			}
		},
		Op: func(key interface{}) {
			p := h.pMgr.Get(key)
			if p == nil {
				return
			}

			if !p.Auto() || !p.Op() {
				return
			}

			p.SetOp(false)
			if h.event != nil && h.event.PlayerOp != nil {
				h.event.PlayerOp(key, h.pMgr, h.extend)
			}
			h.game.PlayerAuto(key)
		},
	})

	h.pMgr.Watch(pmgr.WatchKindLine, func(key interface{}, p *player.Player) {
		if h.event != nil && h.event.PlayerLine != nil {
			h.event.PlayerLine(key, h.pMgr, h.extend)
		}

		// 玩家踢出
		h.PlayerKickOut(key)
	})
	h.pMgr.Watch(pmgr.WatchKindReady, func(key interface{}, p *player.Player) {
		if h.event != nil && h.event.PlayerReady != nil {
			h.event.PlayerReady(key, h.pMgr, h.extend)
		}
	})
	h.pMgr.Watch(pmgr.WatchKindOp, func(key interface{}, p *player.Player) {
		if h.event != nil && h.event.PlayerOp != nil {
			h.event.PlayerOp(key, h.pMgr, h.extend)
		}
		if p.Op() {
			h.AutoPlayer(key)
		}
	})
	h.pMgr.Watch(pmgr.WatchKindAuto, func(key interface{}, p *player.Player) {
		if h.event != nil && h.event.PlayerAuto != nil {
			h.event.PlayerAuto(key, h.pMgr, h.extend)
		}
		if p.Op() {
			switch p.Auto() {
			case true:
				h.AutoPlayer(key)
			default:
				h.AutoPlayerDel(key)
			}
		}
	})
	h.pMgr.Watch(pmgr.WatchKindScore, func(key interface{}, p *player.Player) {
		if h.event != nil && h.event.PlayerScore != nil {
			h.event.PlayerScore(key, h.pMgr, h.extend)
		}
	})

	return nil
}

func (h *Happy) Run(resume bool) {
	defer func() {
		err := recover()
		switch err {
		case nil:
		case PanicDoneExit:
			// TODO log
		default:
		}
	}()

	if resume {
		time.Sleep(200 * time.Millisecond)
		if h.playerMsg != nil {
			userKeys := make([]interface{}, 0, h.pMgr.Len())
			h.pMgr.Range(func(key interface{}, p *player.Player) bool {
				userKeys = append(userKeys, key)
				return true
			})
			h.playerMsg.ReConn(h.ctx, userKeys)
		}
		if h.begin {
			h.begin = false
			h.curRound--
			h.RoundBegin(true)
		}
	}
	h.Loop(0)
}

func (h *Happy) Loop(timeout time.Duration) {
	var (
		timeoutTimer = time.NewTimer(timeout)
		msg          *proxy.Msg
	)
	if timeout <= 0 {
		timeoutTimer.Stop()
	}
Loop:
	for {
		if h.msgChan == nil {
			break Loop
		}

		select {
		case msg = <-h.msgChan:
			if msg == nil {
				continue
			}
			h.MsgHandler(msg)
		case <-timeoutTimer.C:
			if timeout > 0 {
				break Loop
			}
		case <-h.delay.Done():
			h.delay.Handler()
		case <-h.ctx.Done():
			h.Finish(false)
		}
	}
}
