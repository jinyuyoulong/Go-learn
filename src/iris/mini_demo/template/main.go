package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()
	// 必须绝对路径，大坑。 相对路径报 html/template: "hello.html" is undefined
	app.RegisterView(iris.HTML("/Users/fanjinlong/dev/go/golib/src/github.com/jinyuyoulong/Go-learn/iris/template/views", ".html"))

	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("message", "Hello world!")
		ctx.View("hello.html")
	})

	app.Run(iris.Addr(":8080"))
}
