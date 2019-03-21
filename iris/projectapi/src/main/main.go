package main

import (
	"v5u.win/golearn/iris/projectapi/src/app/bootstrap"
	"v5u.win/golearn/iris/projectapi/src/app/web/middleware/identity"
	"v5u.win/golearn/iris/projectapi/src/app/web/routes"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("projectapi database", "fanjinlong")
	app.Bootstrap()
	// 设置identity 为全局中间件
	app.Configure(identity.Configure, routes.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen(":8080")
}
