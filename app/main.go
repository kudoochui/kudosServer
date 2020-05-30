package main

import (
	"flag"
	"github.com/kudoochui/kudos/app"
	"github.com/kudoochui/kudos/log"
	_ "github.com/kudoochui/kudosServer/app/gate"
	_ "github.com/kudoochui/kudosServer/app/user"
	"github.com/kudoochui/kudosServer/config"
)

var (
	stype = flag.String("type", "", "server type")
	sid = flag.String("id", "", "server id")
)

func main() {
	flag.Parse()

	//切换到使用protobuf, 默认使用json
	//codecService.SetCodecType(codecService.TYPE_CODEC_PROTOBUF)

	if *stype != "" {
		f := app.GetCreateServerFunc(*stype)
		if *sid != "" {
			app.Run(f(*sid))
		} else {
			setting, err := config.ServersConfig.GetMap(*stype)
			if err != nil {
				log.Error("%s", err)
			}
			servers := make([]app.Server, 0)
			for k,_ := range setting {
				servers = append(servers, f(k))
			}
			app.Run(servers...)
		}
	} else {
		settings,_ := config.ServersConfig.GetEnvMap()
		servers := make([]app.Server, 0)
		for stype, setting := range settings {
			f := app.GetCreateServerFunc(stype)
			for k,_ := range setting.(map[string]interface{}) {
				servers = append(servers, f(k))
			}
		}
		app.Run(servers...)
	}
}
