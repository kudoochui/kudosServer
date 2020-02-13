package main

import (
	"flag"
	"github.com/kudoochui/kudos/app"
	"github.com/kudoochui/kudos/log"
	"github.com/kudoochui/kudosServer/app/gate"
	"github.com/kudoochui/kudosServer/app/user"
	"github.com/kudoochui/kudosServer/config"
)

var (
	stype = flag.String("type", "", "server type")
	sid = flag.String("id", "", "server id")
)

func main() {
	flag.Parse()

	//codec.SetCodecType(codec.TYPE_CODEC_PROTOBUF)

	switch *stype {
	case "gate":
		if *sid == "" {
			// startup all gate
			settings, err := config.ServersConfig.GetMap("gate")
			if err != nil {
				log.Error("%s", err)
			}
			servers := make([]app.Server, 0)
			for k,_ := range settings {
				servers = append(servers, gate.Server(k))
			}
			app.Run(servers...)
		} else {
			// startup specified gate
			app.Run(gate.Server(*sid))
		}

	case "user":
		if *sid == "" {
			// startup all user server
			settings, err := config.ServersConfig.GetMap("user")
			if err != nil {
				log.Error("%s", err)
			}
			servers := make([]app.Server, 0)
			for k,_ := range settings {
				servers = append(servers, user.Server(k))
			}
			app.Run(servers...)
		} else {
			// startup specified gate
			app.Run(user.Server(*sid))
		}

	default:
		// startup all server

		//gate
		gateSettings, err := config.ServersConfig.GetMap("gate")
		if err != nil {
			log.Error("%s", err)
		}
		userSettings, err := config.ServersConfig.GetMap("user")
		if err != nil {
			log.Error("%s", err)
		}
		servers := make([]app.Server, 0)
		for k,_ := range gateSettings {
			servers = append(servers, gate.Server(k))
		}
		for k,_ := range userSettings {
			servers = append(servers, user.Server(k))
		}
		app.Run(servers...)
	}
}
