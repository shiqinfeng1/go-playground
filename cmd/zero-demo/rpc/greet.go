package main

import (
	"flag"
	"fmt"
	"net"

	"go-playground/cmd/zero-demo/rpc/internal/config"
	"go-playground/cmd/zero-demo/rpc/internal/server"
	"go-playground/cmd/zero-demo/rpc/internal/svc"
	"go-playground/cmd/zero-demo/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func getFreePort() (int, error) {
	// 动态获取可用端口
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:0")
	if err != nil {
		return 0, err
	}
	l, err := net.Listen("tcp", addr.String())
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}

var configFile = flag.String("f", "etc/greet.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	//获取动态接口口
	port, _ := getFreePort()
	// 替换yaml里面的host和端口
	c.ListenOn = fmt.Sprintf("0.0.0.0:%d", port)

	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterGreetServer(grpcServer, server.NewGreetServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
