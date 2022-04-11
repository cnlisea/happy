package happy

import (
	"context"
	"github.com/cnlisea/happy/delay"
	"github.com/cnlisea/happy/heartbeat"
	"github.com/cnlisea/happy/pmgr"
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

	msgChan chan *proxy.MsgKind

	event  *Event
	plugin *Plugin

	playerMsg PlayerMsg

	vote *vote.Vote

	ownerUserKey interface{}

	game               proxy.Game
	begin              bool   // 开始状态
	curRound, maxRound uint32 // 局数
}

func New(ctx context.Context) *Happy {
	if ctx == nil {
		ctx = context.Background()
	}
	return &Happy{
		ctx:   ctx,
		delay: delay.New(),
		pMgr:  pmgr.New(),
	}
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
		// TODO resume
	}

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
