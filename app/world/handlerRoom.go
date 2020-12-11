package world

import (
	"context"
	"errors"
	"fmt"
	"github.com/kudoochui/kudos/rpc"
	"github.com/kudoochui/kudos/service/channelService"
	"github.com/kudoochui/kudosServer/app/world/msg"
)

type Room struct {
	roomRemote *RoomRemote
}

type RoomRemote struct {
	room 	*Room
}

const (
	ROOM = "roomA"
)

func (r *Room) OnLeave(session *rpc.Session) error {
	c := channelService.GetChannelService().GetChannel(ROOM)
	if c == nil {
		return nil
	}
	c.Leave(session.GetUserId())

	msg := &msg.HelloResp{
		Words: fmt.Sprintf("%d leave", session.GetSessionId()),
	}

	c.PushMessage("onLeave", msg, nil)
	return nil
}

func (r *RoomRemote) Join(ctx context.Context, args *rpc.Args, replay *msg.RoomResp) error {
	var req msg.RoomJoin
	args.GetObject(&req)

	replay.Code = 200

	c := channelService.GetChannelService().GetChannel(ROOM)
	if c == nil {
		c = channelService.GetChannelService().CreateChannel(ROOM)
	}
	c.Add(&args.Session)

	msg := &msg.HelloResp{
		Words: "welcome",
	}

	c.PushMessage("onJoin", msg, nil)
	return nil
}

func (r *RoomRemote) Leave(ctx context.Context, args *rpc.Args, replay *msg.RoomResp) error {
	var req msg.RoomLeave
	args.GetObject(&req)

	replay.Code = 200

	c := channelService.GetChannelService().GetChannel(ROOM)
	if c == nil {
		replay.Code = 100
		return errors.New("not in room")
	}
	c.Leave(args.Session.GetUserId())

	msg := &msg.HelloResp{
		Words: req.Name + " leave",
	}

	c.PushMessage("onLeave", msg, nil)
	return nil
}

func (r *RoomRemote) Say(ctx context.Context, args *rpc.Args, replay *msg.HelloResp) error {
	var req msg.HelloReq
	args.GetObject(&req)

	replay.Words = "success"

	c := channelService.GetChannelService().GetChannel(ROOM)
	if c == nil {
		return errors.New("not in room")
	}

	msg := &msg.HelloResp{
		Words: req.Words,
	}

	c.PushMessage("onSay", msg, nil)
	return nil
}