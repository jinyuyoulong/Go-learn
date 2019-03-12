package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()
	appEngine := iris.HTML("/Users/fanjinlong/dev/go/golib/src/v5u.win/golearn/iris/superstar/_demo/iris/", ".html")
	app.RegisterView(appEngine)

	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString("Hello World ----by Iris ")
	})
	app.Get("hello", func(ctx iris.Context) {
		ctx.ViewData("Title", "测试页面")
		ctx.ViewData("message", "Hello World ----by Iris template")
		ctx.View("hello.html")
	})
	app.Run(iris.Addr(":8080"), iris.WithCharset("UTF-8"))
}
