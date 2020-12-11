package user

import (
	"github.com/kudoochui/kudos/component"
	"github.com/kudoochui/kudos/component/remote"
	msgService "github.com/kudoochui/kudos/service/msgService"
	"github.com/kudoochui/kudosServer/app/user/msg"
)

// register server service to remote
var msgArray = []interface{}{}

func RegisterHandler(msg interface{}){
	msgArray = append(msgArray, msg)
}

type MsgHandler struct {
	server		component.ServerImpl
	rpcServer 		*remote.Remote
}

func NewMsgHandler(s component.ServerImpl) *MsgHandler {
	h := &MsgHandler{server:s}
	h.rpcServer = h.server.GetComponent("remote").(*remote.Remote)
	return h
}

func (m *MsgHandler)RegisterHandler()  {
	for _,v := range msgArray {
		m.rpcServer.RegisterHandler(v,"")
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
	msgService.GetMsgService().Register("User.Login", 0, 0, &LoginReq{}, &LoginResp{})
	msgService.GetMsgService().Register("Hello.Say", 0, 0, &HelloReq{}, &HelloResp{})
	msgService.GetMsgService().Register("Hi.Say",0, 0, &msg.HiReq{}, &msg.HiResp{})		//test pomelo pb
	msgService.GetMsgService().Register("RoomRemote.Join", 0, 0, &RoomJoin{}, &RoomResp{})
	msgService.GetMsgService().Register("RoomRemote.Leave", 0, 0, &RoomLeave{}, &RoomResp{})
	msgService.GetMsgService().Register("RoomRemote.Say", 0, 0, &HelloReq{}, &HelloResp{})
	msgService.GetMsgService().RegisterPush("onNotify", 0)
	msgService.GetMsgService().RegisterPush("onLeave", 0)
	msgService.GetMsgService().RegisterPush("onJoin", 0)
	msgService.GetMsgService().RegisterPush("onSay", 0)
}
