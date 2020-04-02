package main

import (
	"inn/internal/user/conf"
	"inn/internal/user/repository/persistence/orm"
	"inn/internal/user/rpc"
	"inn/pkg/tracer"
	"log"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"

	userpb "inn/pb/user"

	traceplugin "github.com/micro/go-plugins/wrapper/trace/opentracing"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/transport/grpc"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
)

const name = "go.micro.srv.user"

func main() {
	//初始化
	conf.Init()

	orm.Init()
	defer orm.Close()

	etcdRegistry := etcdv3.NewRegistry(
		func(opt *registry.Options) {
			opt.Addrs = []string{viper.GetString("ETCD.ADDR")}
		},
	)

	t, io, err := tracer.NewTracer(name, viper.GetString("JAEGER.ADDR"))
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	// 创建服务
	srv := micro.NewService(
		micro.Name(name),
		micro.Registry(etcdRegistry),
		micro.Transport(grpc.NewTransport()),
		micro.WrapHandler(traceplugin.NewHandlerWrapper(t)),
		micro.WrapCall(traceplugin.NewCallWrapper(t)),
	)

	srv.Init()

	userpb.RegisterUserHandler(
		srv.Server(),
		rpc.NewUserRpc(orm.NewUserRepository()),
	)

	go func() {
		if err := srv.Run(); err != nil {
			panic(err)
		}
	}()

	log.Println("user service start")
	//退出服务
	var state int32 = 1
	sc := make(chan os.Signal)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	sig := <-sc
	atomic.StoreInt32(&state, 0)
	log.Printf("received exit signal:[%s]", sig.String())

	log.Println("user service shutdown")
	os.Exit(int(atomic.LoadInt32(&state)))

}
