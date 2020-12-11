package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/kudoochui/kudos/network"
	"github.com/kudoochui/kudosServer/bin/goclient/msg"
	"os"
	"os/signal"
)

type Agent struct {
	client 	*network.TCPClient
	conn      network.Conn
}

func (a *Agent) Run() {
	go func() {
		// login
		req := &msg.LoginReq{Account:"kudoo"}
		buffer, _ := proto.Marshal(req)
		a.client.MsgParser.Write(a.conn.(*network.TCPConn), uint32(msg.ECustomerMsgType_REQ_LOGIN), buffer)
	}()

	for {
		pkgType, body, err := a.client.MsgParser.Read(a.conn.(*network.TCPConn))
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		switch pkgType {
		case uint32(msg.ECustomerMsgType_RESP_LOGIN):
			resp := &msg.LoginResp{}
			proto.Unmarshal(body,resp)
			fmt.Println("login response:", resp)

			a.JoinRoom()
			break;
		case uint32(msg.ECustomerMsgType_RESP_JOIN_ROOM):
			resp := &msg.RoomResp{}
			proto.Unmarshal(body,resp)
			fmt.Println("join response:", resp)

			a.Say()
			a.LeaveRoom()
			break;
		case uint32(msg.ECustomerMsgType_RESP_LEAVE_ROOM):
			resp := &msg.RoomResp{}
			proto.Unmarshal(body,resp)
			fmt.Println("leave response:", resp)
			break;
		case uint32(msg.ECustomerMsgType_RESP_HELLO),uint32(msg.ECustomerMsgType_PUSH_SAY),uint32(msg.ECustomerMsgType_PUSH_JOIN),uint32(msg.ECustomerMsgType_PUSH_LEAVE):
			resp := &msg.HelloResp{}
			proto.Unmarshal(body,resp)
			fmt.Println(pkgType, resp)
			break;
		}

	}
}

func (a *Agent) OnClose() {

}

func (a *Agent) JoinRoom() {
	req := &msg.RoomJoin{Name:"kudoo"}
	buffer, _ := proto.Marshal(req)
	a.client.MsgParser.Write(a.conn.(*network.TCPConn), uint32(msg.ECustomerMsgType_REQ_JOIN_ROOM), buffer)
}

func (a *Agent) LeaveRoom() {
	req := &msg.RoomLeave{Name:"kudoo"}
	buffer, _ := proto.Marshal(req)
	a.client.MsgParser.Write(a.conn.(*network.TCPConn), uint32(msg.ECustomerMsgType_REQ_LEAVE_ROOM), buffer)
}

func (a *Agent) Say() {
	req := &msg.HelloReq{Words:"hello"}
	buffer, _ := proto.Marshal(req)
	a.client.MsgParser.Write(a.conn.(*network.TCPConn), uint32(msg.ECustomerMsgType_REQ_HELLO), buffer)
}

func main ()  {
	var client *network.TCPClient
	client = &network.TCPClient{
		Addr: "192.168.38.70:5020",
		ConnNum: 1,
		LenMsgLen: 4,
		MinMsgLen: 4,
		MaxMsgLen: 4096,
		ConnectInterval: 2,
		NewAgent: func(conn *network.TCPConn) network.Agent {
			a := &Agent{
				client: client,
				conn: conn,
			}
			return a
		},
	}

	client.Start()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
	client.Close()
}