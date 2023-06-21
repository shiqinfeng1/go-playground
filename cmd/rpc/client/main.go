package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

const HelloServiceName = "HelloService"

type HelloServiceIntf = interface {
	Hello(request string, reply *string) error
	Hi(request string, reply *string) error
}

type HelloServiceClient struct {
	client *rpc.Client
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.client.Call(HelloServiceName+".Hello", request, reply)
}
func (p *HelloServiceClient) Hi(request string, reply *string) error {
	return p.client.Call(HelloServiceName+".Hi", request, reply)
}

func DialHelloService(network, address string) (HelloServiceIntf, error) {
	// c, err := rpc.Dial(network, address)
	// if err != nil {
	// 	return nil, err
	// }
	// return &HelloServiceClient{client: c}, nil

	conn, err := net.Dial(network, address)
	if err != nil {
		log.Fatal("net.Dial:", err)
	}
	return &HelloServiceClient{client: rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))}, nil
}

var _ HelloServiceIntf = (*HelloServiceClient)(nil)

func main() {
	client, err := DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Hello("世界", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
	err = client.Hi("世界", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
