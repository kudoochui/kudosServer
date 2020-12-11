package gate

import (
	"fmt"
	"github.com/kudoochui/kudos/app"
	"github.com/kudoochui/kudos/component/connector/protobuf"
	rpcServer "github.com/kudoochui/kudos/component/remote"
	"github.com/kudoochui/kudos/log"
	"github.com/kudoochui/kudos/rpc"
	"github.com/kudoochui/kudos/service/rpcClientService"
	"github.com/kudoochui/kudosServer/config"
)

type Gate struct {
	*app.ServerDefault
}

func init()  {
	app.RegisterCreateServerFunc("gatebp", func(serverId string) app.Server {
		return &Gate{
			ServerDefault: app.NewServerDefault(serverId),
		}
	})
}

func (g *Gate) OnStart(){
	settings, err := config.ServersConfig.GetMap("gatebp")
	if err != nil {
		log.Error("%s", err)
	}
	serverSetting := settings[g.ServerId].(map[string]interface{})
	wsAddr := fmt.Sprintf("%s:%.f",serverSetting["host"], serverSetting["clientPort"])
	remoteAddr := fmt.Sprintf("%s:%.f",serverSetting["host"], serverSetting["port"])
	conn := protobuf.NewConnector(
		protobuf.WSAddr(wsAddr),
		)
	g.Components["connector"] = conn

	remote := rpcServer.NewRemote(
		rpcServer.Addr(remoteAddr),
		rpcServer.RegistryType(config.RegistryConfig.String("registry")),
		rpcServer.RegistryAddr(config.RegistryConfig.String("addr")),
		rpcServer.BasePath(config.RegistryConfig.String("basePath")))
	g.Components["remote"] = remote

	g.OnInit()

	conn.SetConnectionListener(g)
}

func (g *Gate) Run(closeSig chan bool){
	g.OnRun(closeSig)

	<-closeSig
	//closing
	log.Info("gate closing")
}

func (g *Gate) OnStop(){
	g.OnDestroy()
}

func (g *Gate) OnDisconnect(session *rpc.Session) {
	args := &rpc.Args{
		Session: *session,
	}
	reply := &rpc.Reply{}
	rpcClientService.GetRpcClientService().Call("User", "OnOffline", args, reply)
}