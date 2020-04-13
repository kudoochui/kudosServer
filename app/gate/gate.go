package gate

import (
	"fmt"
	"github.com/kudoochui/kudosServer/config"
	"github.com/kudoochui/kudos/app"
	"github.com/kudoochui/kudos/component/connector"
	rpcClient "github.com/kudoochui/kudos/component/proxy"
	rpcServer "github.com/kudoochui/kudos/component/remote"
	"github.com/kudoochui/kudos/component/timers"
	"github.com/kudoochui/kudos/log"
	"github.com/kudoochui/kudos/rpc"
)

type Gate struct {
	*app.ServerDefault

	msgHandler *MsgHandler
}

func init()  {
	app.RegisterCreateServerFunc("gate", func(serverId string) app.Server {
		return &Gate{
			ServerDefault: app.NewServerDefault(serverId),
		}
	})
}

func (g *Gate) OnStart(){
	settings, err := config.ServersConfig.GetMap("gate")
	if err != nil {
		log.Error("%s", err)
	}
	serverSetting := settings[g.ServerId].(map[string]interface{})
	wsAddr := fmt.Sprintf("%s:%.f",serverSetting["host"], serverSetting["clientPort"])
	remoteAddr := fmt.Sprintf("%s:%.f",serverSetting["host"], serverSetting["port"])
	conn := connector.NewConnector(
		connector.WSAddr(wsAddr),
		)
	g.Components["connector"] = conn

	remote := rpcServer.NewRemote(
		rpcServer.Addr(remoteAddr),
		rpcServer.RegistryType(config.RegistryConfig.String("registry")),
		rpcServer.RegistryAddr(config.RegistryConfig.String("addr")),
		rpcServer.BasePath(config.RegistryConfig.String("basePath")))
	g.Components["remote"] = remote
	g.msgHandler = &MsgHandler{r:remote}

	proxy := rpcClient.NewProxy(
		rpcClient.RegistryType(config.RegistryConfig.String("registry")),
		rpcClient.RegistryAddr(config.RegistryConfig.String("addr")),
		rpcClient.BasePath(config.RegistryConfig.String("basePath")))
	g.Components["proxy"] = proxy

	timer := timers.NewTimer()
	g.Components["timer"] = timer

	for _,com := range g.Components {
		com.OnInit()
	}

	// register service.  Note: must behind remote OnInit
	g.msgHandler.RegisterHandler()

	conn.SetRouter(proxy)
	conn.SetRegisterServiceHandler(remote)
	proxy.SetRpcResponder(conn)
	conn.SetConnectionListener(g)
}

func (g *Gate) Run(closeSig chan bool){
	for _,com := range g.Components {
		go com.Run(closeSig)
	}

	<-closeSig
	//closing
	log.Info("gate closing")
}

func (g *Gate) OnStop(){
	for _,com := range g.Components {
		com.OnDestroy()
	}
}

func (g *Gate) OnDisconnect(session *rpc.Session) {
	proxy := g.GetComponent("proxy").(*rpcClient.Proxy)
	args := &rpc.Args{
		Session: *session,
	}
	reply := &rpc.Reply{}
	proxy.RpcCall("User", "OnOffline", args, reply)
}