package main

import (
	"go-playground/cmd/rpc/server/service"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	// fmt.Printf("print a null string:%v", nil)
	service.RegisterHelloService(new(service.HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
			continue
		}
		log.Println("accept a client conn")
		// go rpc.ServeConn(conn)
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

}
