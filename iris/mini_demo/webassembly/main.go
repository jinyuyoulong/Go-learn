package main

import (
	"net/http"
	"strings"

	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.HTML("<b>Resource Not found</b>")
	})
	app.Get("/", func(ctx iris.Context) {
		ctx.ServeFile("./public/index.html", false)
	})
	app.Get("/profile/{username}", func(ctx iris.Context) {
		ctx.Writef("Hello %s", ctx.Params().Get("username"))
	})
	//提供来自根“/”的文件，如果我们使用.StaticWeb它可以覆盖
	//由于下划线需要通配符，所有路由。
	//在这里，我们将看到如何绕过这种行为
	//通过创建一个新的文件服务器处理程序和
	//为路由器设置包装器（如“低级”中间件）
	//为了手动检查我们是否想要正常处理路由器
	//或者执行文件服务器处理程序。
	//使用.StaticHandler
	//它与StaticWeb相同，但不是
	//注册路由，它只返回处理程序。
	fileServer := app.StaticHandler("./public", false, false)
	//使用本机net / http处理程序包装路由器。
	//如果url不包含任何“。” （即：.css，.js ......）
	//（取决于应用程序，您可能需要添加更多文件服务器异常），
	//然后处理程序将执行负责的路由器
	//注册路线（看“/”和“/ profile / {username}”）
	//如果没有，那么它将根据根“/”路径提供文件。
	app.WrapRouter(func(w http.ResponseWriter, r *http.Request, router http.HandlerFunc) {
		path := r.URL.Path
		//请注意，如果path的后缀为“index.html”，则会自动重定向到“/”，
		//所以我们的第一个处理程序将被执行。
		if !strings.Contains(path, ".") {
			//如果它不是资源，那么就像正常情况一样继续使用路由器. <-- IMPORTANT
			router(w, r)
			return
		}
		//获取并释放上下文，以便使用它来执行我们的文件服务器
		//记得：我们使用net / http.Handler，因为我们在路由器本身之前处于“低级别”。
		ctx := app.ContextPool.Acquire(w, r)
		fileServer(ctx)
		app.ContextPool.Release(ctx)
	})
	// http://localhost:8080
	// http://localhost:8080/index.html
	// http://localhost:8080/app.js
	// http://localhost:8080/css/main.css
	// http://localhost:8080/profile/anyusername
	app.Run(iris.Addr(":8080"))
	//注意：在这个例子中我们只看到一个用例，
	//你可能想要.WrapRouter或.Downgrade以绕过虹膜的默认路由器，即：
	//您也可以使用该方法设置自定义代理。
	//如果您只想在除root之外的其他路径上提供静态文件
	//你可以使用StaticWeb, i.e:
	//                          .StaticWeb("/static", "./public")
	// ________________________________requestPath, systemPath
}
