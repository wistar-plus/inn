package route

import (
	"inn/internal/gateway/controller"
	"inn/internal/gateway/route/middlewares/jaeger"
	"inn/internal/gateway/service"
	"net/http"

	"inn/internal/gateway/repository/persistence/redis"
	"inn/internal/gateway/repository/persistence/syncmap"

	msgpb "inn/pb/message"
	userpb "inn/pb/user"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/transport/grpc"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/spf13/viper"
)

func Router(e *gin.Engine) {
	e.LoadHTMLGlob("*.html")
	e.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	e.Use(jaeger.SetUp())

	e.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//user
	etcdRegistry := etcdv3.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{viper.GetString("ETCD.ADDR")}
	})

	userSrvice := micro.NewService(
		micro.Name("go.micro.cli.user"),
		micro.Registry(etcdRegistry),
		micro.Transport(grpc.NewTransport()),
	)
	userSrvice.Init()
	userSrv := userpb.NewUserService("go.micro.srv.user", userSrvice.Client())

	userController := controller.NewUserController(userSrv)
	e.POST("/register", userController.Register)
	e.POST("/login", userController.Login)

	//ws
	msgSrvice := micro.NewService(
		micro.Name("go.micro.cli.message"),
		micro.Registry(etcdRegistry),
		micro.Transport(grpc.NewTransport()),
	)
	msgSrvice.Init()
	msgSrv := msgpb.NewMessageService("go.micro.srv.message", msgSrvice.Client())
	userConnRepo := syncmap.NewUserConnRepository()
	userTopicRepo := redis.NewUserTopicRepository()
	gwService := service.NewGateWayService(userConnRepo, userTopicRepo, msgSrv)
	wsController := controller.NewWSController(gwService)
	e.GET("/ws", wsController.ServeWs)

}
