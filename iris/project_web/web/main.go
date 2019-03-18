package main

import (
	"v5u.win/golearn/iris/project_web/bootstrap"
	"v5u.win/golearn/iris/project_web/web/middleware/identity"
	"v5u.win/golearn/iris/project_web/web/routes"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("project_web database", "fanjinlong")
	app.Bootstrap()
	// 设置identity 为全局中间件
	app.Configure(identity.Configure, routes.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen(":8080")
}
