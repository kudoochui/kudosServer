package user

import (
	"context"
	"github.com/kudoochui/kudos/log"
	"github.com/kudoochui/kudos/rpc"
	"github.com/kudoochui/kudosServer/app/user/msg"
)

type Hi struct {

}

func (h *Hi) Say(ctx context.Context, args *rpc.Args, replay *msg.HiResp) error {
	var req msg.HiReq
	args.GetObject(&req)

	log.Info("session id is: %s", args.Session.GetSessionId())
	args.Session.Bind(123)
	args.Session.Set("Login", "1")
	args.Session.Push()
	replay.Words = "hello " + req.Words
	return nil
}