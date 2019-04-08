go modules

预先处理

正确配置 GOROOT

export GOPROXY="https://athens.azurefd.net" # 微软提供的公共代理，解决golang/x/tools 下载失败

#### 初始化 同时 确定了项目的绝对目录路径
go mod init v5u.win/projectapi
生成 go.mod
导入包：`import "v5u.win/projectapi/src/app/service"`
#### 添加依赖
go mod tidy
#### 生成 go.sum 自动添加依赖关系到 go.mod
go run main.go

下载依赖包
go mod download github.com/pelletier/go-toml


1. 不能在 golib/src 下创建项目
2.  在 Go1.11 版本下，GOPATH 目录中的项目默认是禁用 Go Module 的，需要手动开启
3. export GOPROXY="https://athens.azurefd.net" // 微软提供的公共代理
4. export export GO111MODULE=on // 必须使用module依赖，在src/ 目录下可以使用module

##### 常用命令
```
go mod tidy //拉取缺少的模块，移除不用的模块。
go mod download //下载依赖包
go mod graph //打印模块依赖图
go mod vendor //将依赖复制到vendor下
go mod verify //校验依赖
go mod why //解释为什么需要依赖
go list -m -json all //依赖详情
```