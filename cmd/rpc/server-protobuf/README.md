# 编译protobuf文件
```
cd cmd/rpc/server-protobuf/service/
protoc --go_out=. hello.proto
```
编译后生成hello.pb.go
