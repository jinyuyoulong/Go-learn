package main

import "github.com/kataras/iris/v12"

type User struct {
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	City      string `json:"city"`
	Age       int    `json:"age"`
}

func main() {
	app := iris.New()
	app.RegisterView(iris.HTML("./views", "html").Reload(true))
	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		errMessage := ctx.Values().GetString("error")
		if errMessage != "" {
			ctx.Writef("Internal server error: %s", errMessage)
			return
		}
		ctx.Writef("(Unexpected) internal server error")
	})

	app.Use(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("Begin request for path: %s", ctx.Path())
		ctx.Next()
	})

	app.Subdomain("mysubdomain.").Post("/decode", func(ctx iris.Context) {
	})

	app.Post("/decode", func(ctx iris.Context) {
		var user User
		ctx.ReadJSON(&user)
		ctx.Writef("%s %s is %d years old and comes from %s", user.Firstname, user.Lastname, user.Age, user.City)
	})

	app.Get("/encode", func(ctx iris.Context) {
		doe := User{
			Username:  "fanjinlong",
			Firstname: "fan",
			Lastname:  "jinlong",
			City:      "Beijing",
			Age:       25,
		}
		ctx.JSON(doe)
	})
	app.Get("/profile/{username:string}", profileByUsername)

	usersRoutes := app.Party("/users", logThisMiddleware)
	{
		// method get: http://localhost:8080/users/42
		usersRoutes.Get("/{id:int min(1)}", getUserByID)
		// method post: http://localhost:8080/users/create
		usersRoutes.Post("/create", createUser)
	}

	// Listen for incoming HTTP/1.x & HTTP/2 clients on localhost port 8080.
	app.Run(iris.Addr(":8080"), iris.WithCharset("UTF-8"))
}

func logThisMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Path: %s | IP: %s", ctx.Path(), ctx.RemoteAddr())

	// .Next is required to move forward to the chain of handlers,
	// if missing then it stops the execution at this handler.
	ctx.Next()
}

func profileByUsername(ctx iris.Context) {
	// .Params are used to get dynamic path parameters.
	username := ctx.Params().Get("username")
	ctx.ViewData("Username", username)
	// renders "./views/user/profile.html"
	// with {{ .Username }} equals to the username dynamic path parameter.
	ctx.View("user/profile.html")
}

func getUserByID(ctx iris.Context) {
	userID := ctx.Params().Get("id") // Or convert directly using: .Values().GetInt64 etc...
	// your own db fetch here instead of user :=...
	user := User{Username: "username" + userID}

	ctx.XML(user)
}

func createUser(ctx iris.Context) {
	var user User
	err := ctx.ReadForm(&user)
	if err != nil {
		ctx.Values().Set("error", "create user, read and parse form faild. "+err.Error())
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
	// renders "./views/user/create_verification.html"
	// with {{ . }} equals to the User object, i.e {{ .Username }}, {{.Firstname}} etc...
	ctx.ViewData("", user)
	ctx.View("user/create_verification.html")
}
