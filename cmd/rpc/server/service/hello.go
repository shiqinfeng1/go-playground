package service

import (
	"log"
	"net/rpc"
)

const HelloServiceName = "HelloService"

type HelloServiceIntf = interface {
	Hello(request string, reply *string) error
	Hi(request string, reply *string) error
}

func RegisterHelloService(svc HelloServiceIntf) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	log.Println("reply:", *reply)
	return nil
}
func (p *HelloService) Hi(request string, reply *string) error {
	*reply = "hi:" + request
	log.Println("reply:", *reply)
	return nil
}
