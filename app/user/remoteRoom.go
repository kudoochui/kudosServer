package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/kudoochui/kudos/rpc"
	"github.com/kudoochui/kudos/service/channelService"
)

type RoomJoin struct {
	Route 	string
	Name	string
}

type RoomLeave struct {
	Route 	string
	Name	string
}

type RoomResp struct {
	Code 	int
	ErrMsg	string
}

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

	msg := &HelloResp{
		Words: fmt.Sprintf("%d leave", session.GetSessionId()),
	}

	c.PushMessage("onLeave", msg)
	return nil
}

func (r *RoomRemote) Join(ctx context.Context, args *rpc.Args, replay *RoomResp) error {
	var req RoomJoin
	args.GetObject(&req)

	replay.Code = 200

	c := channelService.GetChannelService().GetChannel(ROOM)
	if c == nil {
		c = channelService.GetChannelService().CreateChannel(ROOM)
	}
	c.Add(&args.Session)

	msg := &HelloResp{
		Words: "welcome",
	}

	c.PushMessage("onJoin", msg)
	return nil
}

func (r *RoomRemote) Leave(ctx context.Context, args *rpc.Args, replay *RoomResp) error {
	var req RoomLeave
	args.GetObject(&req)

	replay.Code = 200

	c := channelService.GetChannelService().GetChannel(ROOM)
	if c == nil {
		replay.Code = 100
		return errors.New("not in room")
	}
	c.Leave(args.Session.GetUserId())

	msg := &HelloResp{
		Words: req.Name + " leave",
	}

	c.PushMessage("onLeave", msg)
	return nil
}

func (r *RoomRemote) Say(ctx context.Context, args *rpc.Args, replay *HelloResp) error {
	var req HelloReq
	args.GetObject(&req)

	replay.Words = "success"

	c := channelService.GetChannelService().GetChannel(ROOM)
	if c == nil {
		return errors.New("not in room")
	}

	msg := &HelloResp{
		Words: req.Words,
	}

	c.PushMessage("onSay", msg)
	return nil
}