package main

import (
	"inn/internal/user/conf"
	"inn/internal/user/controller"
	"inn/internal/user/repository/persistence/orm"
	"inn/internal/user/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

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

	srv := web.NewService(
		web.Name(viper.GetString("SERVER.NAME")),
		web.Registry(etcdRegistry),
		web.Address(viper.GetString("SERVER.ADDR")),
	)

	engine := gin.New()
	route(engine)

	srv.Handle("/", engine)
	go func() {
		if err := srv.Run(); err != nil {
			panic(errors.Wrap(err, "http.user.server run err"))
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

func route(e *gin.Engine) {
	e.LoadHTMLGlob("*.html")
	e.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	e.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	userController := controller.NewUserController(service.NewUserService(orm.NewUserRepository()))

	e.POST("/register", userController.Register)
	e.POST("/login", userController.Login)

	e.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
