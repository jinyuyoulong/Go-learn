go modules

预先处理

正确配置 GOROOT

export GOPROXY="https://athens.azurefd.net" # 微软提供的公共代理，解决golang/x/tools 下载失败

初始化
go mod init v5u.win/projectapi
生成 go.mod

go mod tidy // 添加依赖

go run main.go
生成 go.sum 自动添加依赖关系到 go.mod 



1. 不能在 golib/src 下创建项目
2.  在 Go1.11 版本下，GOPATH 目录中的项目默认是禁用 Go Module 的，需要手动开启
3. export GOPROXY="https://athens.azurefd.net" // 微软提供的公共代理
4. export export GO111MODULE=on // 必须使用module依赖，在src/ 目录下可以使用module