## 准备
### 安装 gRPC / protobuf
安装 [protoc compiler](https://github.com/google/protobuf/releases)

安装 gRPC 和 protoc go 插件
```
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go
```

安装 go-micro 编译器插件 
```
go get -u github.com/micro/protobuf/proto
go get -u github.com/micro/protobuf/protoc-gen-go
```