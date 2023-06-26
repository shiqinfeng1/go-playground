package main

import (
	"flag"
	"fmt"
	"net/http"

	"go-playground/cmd/zero-demo/api/internal/config"
	"go-playground/cmd/zero-demo/api/internal/handler"
	"go-playground/cmd/zero-demo/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	xhttp "github.com/zeromicro/x/http"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithUnauthorizedCallback(func(w http.ResponseWriter, r *http.Request, err error) {
		xhttp.JsonBaseResponseCtx(r.Context(), w, err)
	}))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
