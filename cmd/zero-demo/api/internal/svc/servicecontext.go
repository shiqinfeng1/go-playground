package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"go-playground/cmd/zero-demo/api/internal/config"
	"go-playground/cmd/zero-demo/api/internal/middleware"
)

type ServiceContext struct {
	Config          config.Config
	AuthInterceptor rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		AuthInterceptor: middleware.NewAuthInterceptorMiddleware().Handle,
	}
}
