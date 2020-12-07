package gate

import (
	"github.com/kudoochui/kudos/component"
	"github.com/kudoochui/kudos/component/remote"
	"github.com/kudoochui/kudos/service/msgService"
)

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
	// register service
	m.rpcServer.RegisterHandler(new(Arith), "")
}

func init() {
	// register msg type
	msgService.GetMsgService().Register("Arith.Mul", &Args{}, &Reply{})
}