package main

import (
	"fmt"
	"inn/internal/gateway/conf"
	"inn/internal/gateway/repository/persistence/redis"
	"inn/internal/gateway/repository/persistence/syncmap"
	"inn/internal/gateway/service"
	"inn/internal/gateway/subscriber"
	"inn/internal/gateway/websocket"
	msgpb "inn/pb/message"
	"log"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/broker/rabbitmq"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func main() {
	//初始化
	conf.Init()

	redis.Init()
	defer redis.Close()

	//运行websocket
	etcdRegistry := etcdv3.NewRegistry(
		func(opt *registry.Options) {
			opt.Addrs = []string{viper.GetString("ETCD.ADDR")}
		},
	)

	srv := web.NewService(
		web.Name(viper.GetString("SERVER.NAME")),
		web.Registry(etcdRegistry),
		web.Address(viper.GetString("SERVER.ADDR")),
	)

	msgSrv := micro.NewService(
		micro.Name("grpc.msg.client"),
		micro.Registry(etcdRegistry),
	)
	msgSrv.Init()
	// Create new greeter client
	msg := msgpb.NewMessageService("grpc.msg.server", msgSrv.Client())

	userConnRepo := syncmap.NewUserConnRepository()
	userTopicRepo := redis.NewUserTopicRepository()
	ws := websocket.NewGateWayHandler(service.NewGateWayService(userConnRepo, userTopicRepo, msg))

	srv.Handle("/ws", ws)
	//srv.HandleFunc(, wsHandler)

	go func() {
		if err := srv.Run(); err != nil {
			panic(errors.Wrap(err, "ws.gw.server run err"))
		}
	}()

	rabbitmq.DefaultRabbitURL = "amqp://guest:guest@" + viper.GetString("MQ.ADDR")
	pubSub := rabbitmq.NewBroker(func(opts *broker.Options) {
		opts.Addrs = []string{viper.GetString("MQ.ADDR")}
	})

	pubSub.Init()
	err1 := pubSub.Connect()
	if err1 != nil {
		panic(err1)
	}

	sub := subscriber.NewMessageSubscriber(userConnRepo)
	_, err := pubSub.Subscribe(viper.GetString("TOPIC"), sub.Handler)

	if err != nil {
		fmt.Printf("sub error: %v\n", err)
	}

	log.Println("gateway service start")
	//退出服务
	var state int32 = 1
	sc := make(chan os.Signal)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	sig := <-sc
	atomic.StoreInt32(&state, 0)
	log.Printf("received exit signal:[%s]", sig.String())

	log.Println("gateway service shutdown")
	os.Exit(int(atomic.LoadInt32(&state)))

}
