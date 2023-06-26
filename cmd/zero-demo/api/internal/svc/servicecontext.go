package svc

import (
	"go-playground/cmd/zero-demo/api/internal/config"
	"go-playground/cmd/zero-demo/api/internal/middleware"
	"go-playground/cmd/zero-demo/rpc/greet"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config          config.Config
	AuthInterceptor rest.Middleware
	//定义rpc类型
	Greet greet.Greet
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		AuthInterceptor: middleware.NewAuthInterceptorMiddleware().Handle,
		//引入gprc服务
		Greet: greet.NewGreet(zrpc.MustNewClient(c.GreetRPC)),
	}
}
