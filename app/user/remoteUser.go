package user

import (
	"context"
	"github.com/kudoochui/kudos/rpc"
)

type User struct {
	//server 	*UserServer
	room 	*Room
}

func (u *User) Login(ctx context.Context, args *rpc.Args, replay *rpc.Reply) error {
	session := args.Session
	var UserId int64 = 123121	//这里模拟从数据库获得UserId
	session.Bind(UserId)	//登入成功，连接绑定UserId

	return nil
}

func (u *User) OnOffline(ctx context.Context, args *rpc.Args, replay *rpc.Reply) error {
	u.room.OnLeave(&args.Session)
	return nil
}