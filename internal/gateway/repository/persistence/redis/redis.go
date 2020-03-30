package redis

import (
	"fmt"
	"log"

	"github.com/go-redis/redis/v7"
	"github.com/spf13/viper"
)

var db *redis.Client

func Init() {
	db = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("REDIS.ADDR"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := db.Ping().Result()
	if err != nil {
		log.Fatal("db.Ping() err: ", err)
	}
	fmt.Println(pong)
}

func Close() {
	if db != nil {
		db.Close()
	}
}
