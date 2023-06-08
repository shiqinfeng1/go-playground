package main

import (
	"fmt"
	"go-playground/cmd/rpc/server/service"
	"log"
	"net/rpc"
)

type HelloServiceClient struct {
	client *rpc.Client
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.client.Call(service.HelloServiceName+".Hello", request, reply)
}

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{client: c}, nil
}

var _ service.HelloServiceIntf = (*HelloServiceClient)(nil)

func main() {
	client, err := DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Hello("hello", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
