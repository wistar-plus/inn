package conf

import (
	"strings"

	"github.com/spf13/viper"
)

func Init() {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	viper.SetDefault("JAEGER.ADDR", "localhost:6831")
	viper.SetDefault("ETCD.ADDR", "127.0.0.1:2379")
	viper.SetDefault("MQ.ADDR", "127.0.0.1:5672")
	viper.SetDefault("REDIS.ADDR", "127.0.0.1:6379")

	viper.SetDefault("DB.DRIVER", "mysql")
	viper.SetDefault("DB.USER", "root")
	viper.SetDefault("DB.PASSWORD", "root")
	viper.SetDefault("DB.NAME", "message")
	viper.SetDefault("DB.HOST", "127.0.0.1")
	viper.SetDefault("DB.PORT", "3307")
}
