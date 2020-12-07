package user

import (
	"context"
	"github.com/kudoochui/kudos/rpc"
)

type HelloReq struct {
	Words	string
}

type HelloResp struct {
	Words	string
}

type Hello struct {

}

func (h *Hello) Say(ctx context.Context, args *rpc.Args, replay *HelloResp) error {
	var req HelloReq
	args.GetObject(&req)

	//log.Info("hello" + req.Words)
	replay.Words = "hello " + req.Words

	//go func(){
	//	sessionService.GetSessionService().KickBySid(args.Session.NodeId, args.Session.GetSessionId(), "not allowed")
	//}()

	//route := "onNotify"
	//msg := &HelloResp{
	//	Words: "welcome",
	//}
	//channelService.GetChannelService().PushMessageBySid(args.Session.NodeId, route, msg, []int64{args.Session.SessionId})
	return nil
}