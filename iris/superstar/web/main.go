package main

import (
	"github.com/jinyuyoulong/Go-learn/iris/superstar/bootstrap"
	"github.com/jinyuyoulong/Go-learn/iris/superstar/web/middleware/identity"
	"github.com/jinyuyoulong/Go-learn/iris/superstar/web/routes"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("Superstar database", "fanjinlong")
	app.Bootstrap()
	// 设置identity 为全局中间件
	app.Configure(identity.Configure, routes.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen(":8080")
}
