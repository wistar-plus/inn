package main

import (
	"inn/internal/message/conf"
	"inn/internal/message/repository/persistence/redis"
	"inn/internal/message/rpc"
	"inn/internal/message/service"
	msgpb "inn/pb/message"
	"log"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/broker/rabbitmq"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/spf13/viper"
)

func main() {
	conf.Init()

	redis.Init()
	defer redis.Close()

	etcdRegisty := etcdv3.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{viper.GetString("ETCD.ADDR")}
	})

	rabbitmq.DefaultRabbitURL = "amqp://guest:guest@" + viper.GetString("MQ.ADDR")
	mq := rabbitmq.NewBroker(func(options *broker.Options) {
		options.Addrs = []string{viper.GetString("MQ.ADDR")}
	})

	mq.Init()
	err := mq.Connect()
	if err != nil {
		panic(err)
	}

	// 创建服务
	srv := micro.NewService(
		micro.Name(viper.GetString("SERVER.NAME")),
		micro.Registry(etcdRegisty),
		micro.Broker(mq),
	)

	// Init will parse the command line flags.
	srv.Init()

	// Register handler
	msgpb.RegisterMessageHandler(
		srv.Server(),
		rpc.NewMessageRpc(service.NewMessageService(mq, redis.NewConnAddrRepository())),
	)

	// Run the server
	go func() {
		if err := srv.Run(); err != nil {
			panic(err)
		}
	}()

	log.Println("message service start")
	//退出服务
	var state int32 = 1
	sc := make(chan os.Signal)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	sig := <-sc
	atomic.StoreInt32(&state, 0)
	log.Printf("received exit signal:[%s]", sig.String())
	log.Println("message service shutdown")
	os.Exit(int(atomic.LoadInt32(&state)))
}
