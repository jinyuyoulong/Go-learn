package routes

import (
	"github.com/kataras/iris/mvc"
	"v5u.win/golearn/iris/project_web/bootstrap"
	"v5u.win/golearn/iris/project_web/services"
	"v5u.win/golearn/iris/project_web/web/controllers"
	"v5u.win/golearn/iris/project_web/web/middleware"
)

// Configure registers the necessary routes to the app.
func Configure(b *bootstrap.Bootstrapper) {
	project_webService := services.Newproject_webService()

	index := mvc.New(b.Party("/"))
	index.Register(project_webService)
	index.Handle(new(controllers.IndexController))

	admin := mvc.New(b.Party("/admin"))
	admin.Router.Use(middleware.BasicAuth)
	admin.Register(project_webService)
	admin.Handle(new(controllers.AdminController))

	// b.Get("/", GetIndexHandler)
	// b.Get("/follower/{id:long}", GetFollowerHandler)
	// b.Get("/following/{id:long}", GetFollowingHandler)
	// b.Get("/like/{id:long}", GetLikeHandler)
}
