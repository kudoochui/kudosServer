package user

import (
	"github.com/kudoochui/kudos/rpc"
	msgService "github.com/kudoochui/kudos/service/msgService"
	"github.com/kudoochui/kudosServer/app/user/msg"
)

// register server service to remote
var msgArray = []interface{}{}

func RegisterHandler(msg interface{}){
	msgArray = append(msgArray, msg)
}

type MsgHandler struct {
	r rpc.HandlerRegister
}

func (m *MsgHandler)RegisterHandler()  {
	for _,v := range msgArray {
		m.r.RegisterHandler(v,"")
	}
}

func init() {
	RegisterHandler(new(Hello))
	RegisterHandler(new(Hi))
	room := &Room{
		roomRemote: &RoomRemote{},
	}
	room.roomRemote.room = room
	RegisterHandler(room.roomRemote)
	RegisterHandler(&User{room:room})

	// register msg type
	msgService.GetMsgService().Register("Hello.Say", &HelloReq{}, &HelloResp{})
	msgService.GetMsgService().Register("Hi.Say", &msg.HiReq{}, &msg.HiResp{})
	msgService.GetMsgService().Register("RoomRemote.Join", &RoomJoin{}, &RoomResp{})
	msgService.GetMsgService().Register("RoomRemote.Leave", &RoomLeave{}, &RoomResp{})
	msgService.GetMsgService().RegisterPush("onNotify")
	msgService.GetMsgService().RegisterPush("onLeave")
	msgService.GetMsgService().RegisterPush("onJoin")
	msgService.GetMsgService().RegisterPush("onSay")
}
