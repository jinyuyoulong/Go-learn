 package main
    import (
        "github.com/kataras/iris"
        "github.com/kataras/iris/mvc"
        "github.com/kataras/iris/middleware/logger"
        "github.com/kataras/iris/middleware/recover"
    )
    func main() {
        app := iris.New()
        //（可选）添加两个内置处理程序
        // 可以从任何http相关的恐慌中恢复
        // 并将请求记录到终端。
        app.Use(recover.New())
        app.Use(logger.New())

        // 基于根路由器服务控制器， "/".
        mvc.New(app).Handle(new(ExampleController))

        // http://localhost:8080
        // http://localhost:8080/ping
        // http://localhost:8080/hello
        // http://localhost:8080/custom_path
        app.Run(iris.Addr(":8080"))
    }
    //ExampleController服务于 "/", "/ping" and "/hello".
    type ExampleController struct{}
    // Get serves
    // Method:   GET
    // Resource: http://localhost:8080
    func (c *ExampleController) Get() mvc.Result {
        return mvc.Response{
            ContentType: "text/html",
            Text:        "<h1>Welcome</h1>",
        }
    }
    // GetPing serves
    // Method:   GET
    // Resource: http://localhost:8080/ping
    func (c *ExampleController) GetPing() string {
        return "pong"
    }
    // GetHello serves
    // Method:   GET
    // Resource: http://localhost:8080/hello
    func (c *ExampleController) GetHello() interface{} {
        return map[string]string{"message": "Hello Iris!"}
    }
    //在控制器适应主应用程序之前调用一次BeforeActivation
    //当然在服务器运行之前。
    //在版本9之后，您还可以为特定控制器的方法添加自定义路由。
    //在这里您可以注册自定义方法的处理程序
    //使用带有`ca.Router`的标准路由器做一些你可以做的事情，没有mvc，
    //并添加将绑定到控制器的字段或方法函数的输入参数的依赖项。
    func (c *ExampleController) BeforeActivation(b mvc.BeforeActivation) {
        anyMiddlewareHere := func(ctx iris.Context) {
            ctx.Application().Logger().Warnf("Inside /custom_path")
            ctx.Next()
        }
        b.Handle("GET", "/custom_path", "CustomHandlerWithoutFollowingTheNamingGuide", anyMiddlewareHere)

        //甚至添加基于此控制器路由器的全局中间件，
        //在这个例子中是根“/”：
        // b.Router（）。使用（myMiddleware）
    }
    // CustomHandlerWithoutFollowingTheNamingGuide serves
    // Method:   GET
    // Resource: http://localhost:8080/custom_path
    func (c *ExampleController) CustomHandlerWithoutFollowingTheNamingGuide() string {
        return "hello from the custom handler without following the naming guide"
    }
    // GetUserBy serves
    // Method:   GET
    // Resource: http://localhost:8080/user/{username:string}
    //是一个保留的“关键字”来告诉框架你要去的
    //在函数的输入参数中绑定路径参数，它也是
    //有助于在同一控制器中使用“Get”和“GetBy”。
    //
    // func (c *ExampleController) GetUserBy(username string) mvc.Result {
    //     return mvc.View{
    //         Name: "user/username.html",
    //         Data: username,
    //     }
    // }
    /*
    可以使用多个，工厂会确定
    为每条路线注册了正确的http方法
    对于此控制器，如果需要，请取消注释：
    */
    func (c *ExampleController) Post() {}
    func (c *ExampleController) Put() {}
    func (c *ExampleController) Delete() {}
    func (c *ExampleController) Connect() {}
    func (c *ExampleController) Head() {}
    func (c *ExampleController) Patch() {}
    func (c *ExampleController) Options() {}
    func (c *ExampleController) Trace() {}
    */
    /*
    func (c *ExampleController) All() {}
    //        OR
    func (c *ExampleController) Any() {}

    func (c *ExampleController) BeforeActivation(b mvc.BeforeActivation) {
        // 1 -> the HTTP Method
        // 2 -> the route's path
        // 3 -> this controller's method name that should be handler for that route.
        b.Handle("GET", "/mypath/{param}", "DoIt", optionalMiddlewareHere...)
    }
    // After activation, all dependencies are set-ed - so read only access on them
    // but still possible to add custom controller or simple standard handlers.
    func (c *ExampleController) AfterActivation(a mvc.AfterActivation) {}
    */