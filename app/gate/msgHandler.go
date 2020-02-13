package gate

import (
	"github.com/kudoochui/kudos/rpc"
	"github.com/kudoochui/kudos/service/msgService"
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
	// register service
	RegisterHandler(new(Arith))

	// register msg type
	msgService.GetMsgService().Register("Arith.Mul", &Args{}, &Reply{})
}