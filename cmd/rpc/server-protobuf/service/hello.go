package service

import "net/rpc"

const HelloServiceName = "HelloService"

type HelloServiceIntf = interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(svc HelloServiceIntf) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

type HelloService struct{}

func (p *HelloService) Hello(request *String, reply *String) error {
	reply.Value = "hello:" + request.GetValue()
	return nil
}
