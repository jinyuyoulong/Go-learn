package main

import (
	bootstrap2 "github.com/jinyuyoulong/Go-learn/src/iris/superstar/bootstrap"
	identity2 "github.com/jinyuyoulong/Go-learn/src/iris/superstar/web/middleware/identity"
	routes2 "github.com/jinyuyoulong/Go-learn/src/iris/superstar/web/routes"
)

func newApp() *bootstrap2.Bootstrapper {
	app := bootstrap2.New("Superstar database", "fanjinlong")
	app.Bootstrap()
	// 设置identity 为全局中间件
	app.Configure(identity2.Configure, routes2.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen(":8080")
}
