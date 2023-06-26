生成框架代码：
```
// api
cd ~/workspace/go-playground
goctl api go -api cmd/zero-demo/api/user.api -dir cmd/zero-demo/api

// grpc 
cd ~/workspace/go-playground
goctl rpc protoc cmd/zero-demo/rpc/greet.proto --go_out=cmd/zero-demo/rpc --go-grpc_out=cmd/zero-demo/rpc --zrpc_out=cmd/zero-demo/rpc
```

启动
```
// api
cd ~/workspace/go-playground
go run cmd/zero-demo/api/*.go -f cmd/zero-demo/api/etc/user.yaml

// grpc 
go run cmd/zero-demo/rpc/*.go -f cmd/zero-demo/rpc/etc/greet.yaml
```