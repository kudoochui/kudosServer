# kudosServer
基于kudos游戏开发框架开发的脚手架。

## 安装
###1. 下载
```shell script
git clone https://github.com/kudoochui/kudosServer.git
```
###2. 启动注册中心
[安装consul](https://learn.hashicorp.com/consul/getting-started/install)
```
consul agent --dev
```
###3. 运行游戏
```shell script
go build app/main.go
./main
```

###4. 切换服务器
这里有两组服务器，gate、user一组是使用的pomelo连接器（connector),它支持的是websocket的pomelo通信协议。gatepbg与world是另一组支持
protobuf连接器（connector)的服务器。它同时支持websocket与tcp，使用的是protobuf数据压缩协议。

4.1 pomelo服务器组对应的客户端是bin/clientPomelo

4.2 protobuf服务器组对应客户端：
websocket对应的是bin/client
tcp对应的是bin/goclient
```go
go run main.go
```

4.3 从pomelo服务器组切换到protobuf服务器组
默认是pomelo服务器组。做以下修改切换到protobuf服务器上来

4.3.1 打开main.go上的注释
```go
    //_ "github.com/kudoochui/kudosServer/app/gate"
    //_ "github.com/kudoochui/kudosServer/app/user"
    _ "github.com/kudoochui/kudosServer/app/gatebp"
	_ "github.com/kudoochui/kudosServer/app/world"

...

    //切换到使用protobuf, 默认使用json
	codecService.SetCodecType(codecService.TYPE_CODEC_PROTOBUF)
...
```

4.3.2 打开config.go上的注释
```go
	//ServersConfig, _ = conf.NewAppConfig("servers.json")				// pomelo server
	ServersConfig, _ = conf.NewAppConfig("pbservers.json")		// pb server
```

4.3.3 完成

4.4 websocket切换成tcp
找到gateServer.go
```go
conn := protobuf.NewConnector(
		//protobuf.WSAddr(wsAddr),
		protobuf.TCPAddr(wsAddr),
		)
```
切换。

## 目录介绍
```
/app
    |-  gate                gate服务节点：网络前端
        |- gate.go  
        |- msgHandler.go    注册路由和服务
        |- remoteAritch.go  Aritch服务：一些基础运算
    |-  gatebp              gate服务节点：使用的是pb连接器
    |-  user                user服务节点：后端服务
        |- msg              protobuf消息目录
        |- msgHandler.go    注册路由和服务
        |- remoteHello.go  Hello服务示例
        |- remoteHi.go      Hi服务
        |- remoteRoom.go    房间服务：加入、离开，群发消息
        |- userServer.go  
    |-  world                pb连接对应的后台服务
    main.go                 入口
/bin
    |-  client              protobuf网页测试客户端
    |-  clientPomelo        网页测试客户端，浏览器直接打开，点Test Game Server运行
    |-  conf                游戏配置
    |-  goclient            protobuf tcp连接测试
/config
    |-  config.go           配置文件初始化
```
   
## 分布式部署
```shell script
//启动一个连接服
./main -type gate -id gate-1
//再启动一个
./main -type gate -id gate-2
//启动一个后端服务
./main -type user -id user-1
```
可部署在不同服务器上，在servers.json中配置好ip和端口即可。