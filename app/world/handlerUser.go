package world

import (
	"context"
	"github.com/kudoochui/kudos/rpc"
	"github.com/kudoochui/kudosServer/app/world/msg"
	"math/rand"
)

type User struct {
	//server 	*UserServer
	room 	*Room
}

func (u *User) Login(ctx context.Context, args *rpc.Args, replay *msg.LoginResp) error {
	var req msg.LoginReq
	args.GetObject(&req)

	session := args.Session
	var UserId int64 = rand.Int63()	//这里模拟从数据库获得UserId
	session.Bind(UserId)	//登入成功，连接绑定UserId
	args.Session.Set("Login", "1")	//session里保存额外信息
	args.Session.Push()

	replay.Result = "success"
	return nil
}

func (u *User) OnOffline(ctx context.Context, args *rpc.Args, replay *rpc.Reply) error {
	u.room.OnLeave(&args.Session)
	return nil
}