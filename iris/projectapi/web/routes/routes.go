package routes

import (
	"github.com/kataras/iris/mvc"
	"v5u.win/golearn/iris/projectapi/bootstrap"
	"v5u.win/golearn/iris/projectapi/services"
	"v5u.win/golearn/iris/projectapi/web/controllers"
	"v5u.win/golearn/iris/projectapi/web/middleware"
)

// Configure registers the necessary routes to the app.
func Configure(b *bootstrap.Bootstrapper) {
	projectapiService := services.NewprojectapiService()

	index := mvc.New(b.Party("/"))
	index.Register(projectapiService)
	index.Handle(new(controllers.IndexController))

	admin := mvc.New(b.Party("/admin"))
	admin.Router.Use(middleware.BasicAuth)
	admin.Register(projectapiService)
	admin.Handle(new(controllers.AdminController))

	// -------------------------------------------------------
	// b.Get("/", GetIndexHandler)
	// b.Get("/follower/{id:long}", GetFollowerHandler)
	// b.Get("/following/{id:long}", GetFollowingHandler)
	// b.Get("/like/{id:long}", GetLikeHandler)
}
