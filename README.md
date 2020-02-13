# kudosServer
基于kudos游戏开发框架开发的脚手架。

## 安装
1. 下载
```shell script
git clone https://github.com/kudoochui/kudosServer.git
```
2. 启动注册中心
[安装consul](https://learn.hashicorp.com/consul/getting-started/install)
```
consul agent --dev
```
3. 运行游戏
```shell script
go build app/main.go
./main
```



## 目录介绍
```
/app
    |-  gate                gate服务节点：网络前端
        |- gate.go  
        |- msgHandler.go    注册路由和服务
        |- remoteAritch.go  Aritch服务：一些基础运算
    |-  user                user服务节点：后端服务
        |- msg              protobuf消息目录
        |- msgHandler.go    注册路由和服务
        |- remoteHello.go  Hello服务示例
        |- remoteHi.go      Hi服务
        |- remoteRoom.go    房间服务：加入、离开，群发消息
        |- userServer.go  
    main.go                 入口
/bin
    |-  clientPomelo        网页测试客户端，浏览器直接打开，点Test Game Server运行
    |-  conf                游戏配置
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