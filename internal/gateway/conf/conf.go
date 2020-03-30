package conf

import (
	"strings"

	"github.com/spf13/viper"
)

func Init() {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	viper.SetDefault("SERVER.NAME", "websocket.gateway.server")
	viper.SetDefault("SERVER.ADDR", ":8080")
	viper.SetDefault("ETCD.ADDR", "127.0.0.1:2379")
	viper.SetDefault("MQ.ADDR", "127.0.0.1:5672")
	viper.SetDefault("REDIS.ADDR", "127.0.0.1:6379")

	viper.SetDefault("TOPIC", "gw.1")
}