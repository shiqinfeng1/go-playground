package main

import (
	"go-playground/cmd/rpc/server/service"
	"log"
	"net"
	"net/rpc"
)

func main() {
	service.RegisterHelloService(new(service.HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		log.Println("accept a client conn")
		go rpc.ServeConn(conn)
	}

}
