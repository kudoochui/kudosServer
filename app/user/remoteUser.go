package user

import (
	"context"
	"github.com/kudoochui/kudos/rpc"
)

type User struct {
	//server 	*UserServer
	room 	*Room
}

func (u *User) OnOffline(ctx context.Context, args *rpc.Args, replay *rpc.Reply) error {
	u.room.OnLeave(&args.Session)
	return nil
}