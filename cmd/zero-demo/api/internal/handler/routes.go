// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"
	"time"

	user "go-playground/cmd/zero-demo/api/internal/handler/user"
	"go-playground/cmd/zero-demo/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: user.LoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1/user"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthCheck},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/info",
					Handler: user.InfoHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/info",
					Handler: user.SetHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/v1/user"),
		rest.WithTimeout(3000*time.Millisecond),
		rest.WithMaxBytes(1024),
	)
}
