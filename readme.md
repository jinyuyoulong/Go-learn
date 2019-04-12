支持环境配置(Mac)

安装protobuf 这是 Google 开源的一套成熟的结构数据序列化机制[protocol buffer]() 

```
brew info protobuf
brew install protobuf
```

检验protobuf安装结果

```
protoc --version
libprotoc 3.5.1
```

安装第三方包

```
go mod download github.com/golang/protobuf/proto
go mod download github.com/golang/protobuf/protoc-gen-go
go mod download google.golang.org/grpc
go install github.com/golang/protobuf/protoc-gen-go //编译 protoc-gen-go 可执行文件
```

创建 protobuf 文件

```
vi add.proto
add some date
```

生成 gRPC 代码

```
protoc -I ./protos ./protos/helloworld.proto --go_out=plugins=grpc:helloworld
或
protoc -I . add.proto --go_out=plugins=grpc:.
```

这生成了 `helloworld.pb.go` ，包含了我们生成的客户端和服务端类，此外还有用于填充、序列化、提取 `HelloRequest` 和 `HelloResponse` 消息类型的类。



在server.go 实现AddServiceServer 的接口方法

实现server & client

打开两个终端，分别启动 serve & client

浏览器访问输出返回结果。