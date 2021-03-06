package happy

import (
	"context"
	"fmt"
	"runtime"
	"time"

	"github.com/cnlisea/happy/auto"
	"github.com/cnlisea/happy/delay"
	"github.com/cnlisea/happy/heartbeat"
	"github.com/cnlisea/happy/log"
	"github.com/cnlisea/happy/pmgr"
	"github.com/cnlisea/happy/pmgr/player"
	"github.com/cnlisea/happy/proxy"
	"github.com/cnlisea/happy/vote"
)

type Happy interface {
	Context() context.Context
	Init() error
	Resume(begin bool, curRound uint32)
	Run(resume bool)
	Cost(mode CostMode)
	Event(e *Event)
	Heartbeat(interval time.Duration) error
	Msg(msg *proxy.Msg)
	MsgByUser(f func(userKey interface{}, data interface{}, delay proxy.Delay, curRound, maxRound uint32, pMgr *pmgr.PMgr, extend map[string]interface{}))
	Owner(userKey interface{})
	PlayerMsg(msg proxy.PlayerMsg)
	Plugin(p *Plugin)
	RoundBeginPolicy(policy RoundBeginPolicy)
	LogSetting(path string, level log.Level) error
	Log() *log.Logger
}

type _Happy struct {
	ctx                context.Context
	delay              *delay.Delay
	heartbeat          *heartbeat.Heartbeat
	pMgr               *pmgr.PMgr
	event              *Event
	plugin             *Plugin
	costMode           CostMode
	auto               *auto.Auto
	disbandVote        *vote.Vote
	quickVote          *vote.Vote
	msgChan            chan *proxy.Msg
	byUserHandler      func(userKey interface{}, data interface{}, delay proxy.Delay, curRound, maxRound uint32, pMgr *pmgr.PMgr, extend map[string]interface{})
	playerMsg          proxy.PlayerMsg
	roundBeginPolicy   RoundBeginPolicy
	ownerUserKey       interface{}
	game               proxy.Game
	begin              bool
	curRound, maxRound uint32
	extend             map[string]interface{}
	log                *log.Logger
}

func New(ctx context.Context, maxRound uint32, game proxy.Game, extend map[string]interface{}) Happy {
	if ctx == nil {
		ctx = context.Background()
	}
	return &_Happy{
		ctx:      ctx,
		delay:    delay.New(),
		pMgr:     pmgr.New(),
		msgChan:  make(chan *proxy.Msg, 100),
		game:     game,
		maxRound: maxRound,
		extend:   extend,
	}
}

func (h *_Happy) Context() context.Context {
	return h.ctx
}

func (h *_Happy) Resume(begin bool, curRound uint32) {
	h.begin = begin
	h.curRound = curRound
}

func (h *_Happy) Init() error {
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
			p.SetAutoTs(0)
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
			h.game.PlayerAuto(key)
		},
	})

	h.pMgr.Watch(pmgr.WatchKindLine, func(key interface{}, p *player.Player) {
		if h.event != nil && h.event.PlayerLine != nil {
			h.event.PlayerLine(h, key, h.pMgr, h.extend)
		}

		// ????????????
		h.PlayerKickOut(key)
	})
	h.pMgr.Watch(pmgr.WatchKindReady, func(key interface{}, p *player.Player) {
		if h.event != nil && h.event.PlayerReady != nil {
			h.event.PlayerReady(h, key, h.pMgr, h.extend)
		}
	})
	h.pMgr.Watch(pmgr.WatchKindOp, func(key interface{}, p *player.Player) {
		if h.event != nil && h.event.PlayerOp != nil {
			h.event.PlayerOp(h, key, h.pMgr, h.extend)
		}
		if p.Op() {
			h.AutoPlayer(key)
		}
	})
	h.pMgr.Watch(pmgr.WatchKindAuto, func(key interface{}, p *player.Player) {
		if h.event != nil && h.event.PlayerAuto != nil {
			h.event.PlayerAuto(h, key, h.pMgr, h.extend)
		}
		if p.Op() {
			switch p.Auto() {
			case true:
				h.AutoPlayer(key)
			default:
				h.AutoPlayerDel(key)
				h.AutoPlayer(key)
			}
		}
	})
	h.pMgr.Watch(pmgr.WatchKindScore, func(key interface{}, p *player.Player) {
		if h.event != nil && h.event.PlayerScore != nil {
			h.event.PlayerScore(h, key, h.pMgr, h.extend)
		}
	})
	h.pMgr.Watch(pmgr.WatchKindSite, func(key interface{}, p *player.Player) {
		if h.event != nil && h.event.PlayerSite != nil {
			h.event.PlayerSite(h, key, h.pMgr, h.extend)
		}
	})

	if h.log == nil {
		if err := h.LogSetting("", log.LevelDebug); err != nil {
			return err
		}
	}
	return h.game.Init(h.ctx, h, h.pMgr, h.playerMsg, h.log)
}

func (h *_Happy) Run(resume bool) {
	defer func() {
		err := recover()
		switch err {
		case nil:
		case PanicDoneExit:
			h.log.Info("done exit")
		default:
			var buf = make([]byte, 4096)
			n := runtime.Stack(buf, false)
			buf = buf[:n]
			h.log.Err("run fail",
				log.Bool("resume", resume),
				log.String("err", fmt.Sprintln(err)),
				log.ByteString("stack", buf))
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
			h.RoundBegin(true, false)
		}
	}
	h.Loop()
}

func (h *_Happy) Loop() {
	var msg *proxy.Msg
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
		case <-h.delay.Done():
			h.delay.Handler()
		case <-h.ctx.Done():
			h.Finish(false)
		}
	}
}
