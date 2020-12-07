package user

import (
	"github.com/kudoochui/kudos/component"
	"github.com/kudoochui/kudos/component/remote"
	msgService "github.com/kudoochui/kudos/service/msgService"
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
	room := &Room{
		roomRemote: &RoomRemote{},
	}
	room.roomRemote.room = room
	RegisterHandler(room.roomRemote)
	RegisterHandler(&User{room:room})

	// register msg type
	msgService.GetMsgService().Register("User.Login", &LoginReq{}, &LoginResp{})
	msgService.GetMsgService().Register("Hello.Say", &HelloReq{}, &HelloResp{})
	msgService.GetMsgService().Register("RoomRemote.Join", &RoomJoin{}, &RoomResp{})
	msgService.GetMsgService().Register("RoomRemote.Leave", &RoomLeave{}, &RoomResp{})
	msgService.GetMsgService().Register("RoomRemote.Say", &HelloReq{}, &HelloResp{})
	msgService.GetMsgService().RegisterPush("onNotify")
	msgService.GetMsgService().RegisterPush("onLeave")
	msgService.GetMsgService().RegisterPush("onJoin")
	msgService.GetMsgService().RegisterPush("onSay")
}
