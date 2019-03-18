package main

import (
	"v5u.win/golearn/iris/projectweb/bootstrap"
	"v5u.win/golearn/iris/projectweb/web/middleware/identity"
	"v5u.win/golearn/iris/projectweb/web/routes"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("projectweb database", "fanjinlong")
	app.Bootstrap()
	// 设置identity 为全局中间件
	app.Configure(identity.Configure, routes.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen(":8080")
}
