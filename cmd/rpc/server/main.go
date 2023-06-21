package main

import (
	"go-playground/cmd/rpc/server/service"
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	// 实例化一个具体的rpc服务
	svc := &service.HelloService{}
	// 注册该服务
	service.RegisterHelloService(svc)

	// listener, err := net.Listen("tcp", ":1234")
	// if err != nil {
	// 	log.Fatal("ListenTCP error:", err)
	// }

	// for {
	// 	conn, err := listener.Accept()
	// 	if err != nil {
	// 		log.Fatal("Accept error:", err)
	// 		continue
	// 	}
	// 	log.Println("accept a client conn")
	// 	// go rpc.ServeConn(conn)
	// 	go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	// }
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}

		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	http.ListenAndServe(":1234", nil)

}
