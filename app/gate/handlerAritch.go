package gate


import (
	"context"
	"fmt"
	"github.com/kudoochui/kudos/log"
	"github.com/kudoochui/kudos/rpc"
)

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}

type Arith int


func (t *Arith) Mul(ctx context.Context, args *rpc.Args, reply *Reply) error {
	var msgReq Args
	err := args.GetObject(&msgReq)
	if err != nil {
		log.Error("%+v", err)
	}

	reply.C = msgReq.A * msgReq.B
	fmt.Printf("call: %d * %d = %d\n", msgReq.A, msgReq.B, reply.C)
	return nil
}

func (t *Arith) Add(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A + args.B
	fmt.Printf("call: %d + %d = %d\n", args.A, args.B, reply.C)
	return nil
}

func (t *Arith) Say(ctx context.Context, args *string, reply *string) error {
	*reply = "hello " + *args
	return nil
}