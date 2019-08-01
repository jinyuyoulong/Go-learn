package main

import (
	"github.com/jinyuyoulong/Go-learn/iris/superstar/bootstrap"
	"github.com/jinyuyoulong/Go-learn/iris/superstar/web/middleware/identity"
	"github.com/jinyuyoulong/Go-learn/iris/superstar/web/routes"
)

func main() {
	app := bootstrap.New("Superstar database", "fanjinlong")
	app.Bootstrap()
	// 设置identity 为全局中间件
	app.Configure(identity.Configure, routes.Configure)
	app.Listen(":8081")
}
