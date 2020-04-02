package main

import (
	"context"
	"fmt"
	"inn/internal/gateway/conf"
	"inn/internal/gateway/repository/persistence/redis"
	"inn/internal/gateway/subscriber"
	"inn/pkg/tracer"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"inn/internal/gateway/route"

	"inn/internal/gateway/repository/persistence/syncmap"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-plugins/broker/rabbitmq"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
)

const name = "go.web.srv.gateway"

func main() {
	//初始化
	conf.Init()

	redis.Init()
	defer redis.Close()

	t, io, err := tracer.NewTracer(name, viper.GetString("JAEGER.ADDR"))
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	engine := gin.New()
	route.Router(engine)

	server := &http.Server{
		Addr:         viper.GetString("SERVER.ADDR"),
		Handler:      engine,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("HTTP server listen: %s\n", err)
		}
	}()

	rabbitmq.DefaultRabbitURL = "amqp://guest:guest@" + viper.GetString("MQ.ADDR")
	pubSub := rabbitmq.NewBroker(func(opts *broker.Options) {
		opts.Addrs = []string{viper.GetString("MQ.ADDR")}
	})

	pubSub.Init()
	err = pubSub.Connect()
	if err != nil {
		panic(err)
	}
	userConnRepo := syncmap.NewUserConnRepository()
	sub := subscriber.NewMessageSubscriber(userConnRepo)
	_, err = pubSub.Subscribe(viper.GetString("TOPIC"), sub.Handler)
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("gateway service shutdown")
	os.Exit(int(atomic.LoadInt32(&state)))

}
