package main

import (
	"v5u.win/golearn/iris/projectapi/src/app/bootstrap"
	"v5u.win/golearn/iris/projectapi/src/app/config"
	"v5u.win/golearn/iris/projectapi/src/app/middleware/identity"
	"v5u.win/golearn/iris/projectapi/src/app/route"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("projectapi database", "fanjinlong")
	app.Bootstrap()
	// 设置identity 为全局中间件
	app.Configure(identity.Configure, route.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen(config.Port)
}
