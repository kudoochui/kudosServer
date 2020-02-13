package config

import conf "github.com/kudoochui/kudos/config"

var (
	GameConfig *conf.AppConfig
	RedisConfig *conf.AppConfig
	MysqlConfig *conf.AppConfig
	MongoConfig *conf.AppConfig
	ServersConfig *conf.AppConfig
	RegistryConfig *conf.AppConfig
)

func init()  {
	GameConfig, _ = conf.NewAppConfig("game.json")
	RedisConfig, _ = conf.NewAppConfig("redis.json")
	MongoConfig, _ = conf.NewAppConfig("mongodb.json")
	MysqlConfig, _ = conf.NewAppConfig("mysql.json")
	ServersConfig, _ = conf.NewAppConfig("servers.json")
	RegistryConfig, _ = conf.NewAppConfig("registry.json")

	conf.NodeId,_ = GameConfig.Int64("nodeId")
}

