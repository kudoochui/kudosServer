package world

import (
	"github.com/kudoochui/kudos/component"
	"github.com/kudoochui/kudos/component/remote"
	msgService "github.com/kudoochui/kudos/service/msgService"
	"github.com/kudoochui/kudosServer/app/world/msg"
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
	room := &Room{
		roomRemote: &RoomRemote{},
	}
	room.roomRemote.room = room
	RegisterHandler(room.roomRemote)
	RegisterHandler(&User{room:room})

	// register msg type
	msgService.GetMsgService().Register("User.Login",
		uint32(msg.ECustomerMsgType_REQ_LOGIN),
		uint32(msg.ECustomerMsgType_RESP_LOGIN),
		&msg.LoginReq{}, &msg.LoginResp{})
	msgService.GetMsgService().Register("RoomRemote.Join",
		uint32(msg.ECustomerMsgType_REQ_JOIN_ROOM),
		uint32(msg.ECustomerMsgType_RESP_JOIN_ROOM),
		&msg.RoomJoin{}, &msg.RoomResp{})
	msgService.GetMsgService().Register("RoomRemote.Leave",
		uint32(msg.ECustomerMsgType_REQ_LEAVE_ROOM),
		uint32(msg.ECustomerMsgType_RESP_LEAVE_ROOM),
		&msg.RoomLeave{}, &msg.RoomResp{})
	msgService.GetMsgService().Register("RoomRemote.Say",
		uint32(msg.ECustomerMsgType_REQ_HELLO),
		uint32(msg.ECustomerMsgType_RESP_HELLO),
		&msg.HelloReq{}, &msg.HelloResp{})
	msgService.GetMsgService().RegisterPush("onLeave", uint32(msg.ECustomerMsgType_PUSH_LEAVE))
	msgService.GetMsgService().RegisterPush("onJoin", uint32(msg.ECustomerMsgType_PUSH_JOIN))
	msgService.GetMsgService().RegisterPush("onSay", uint32(msg.ECustomerMsgType_PUSH_SAY))
}
