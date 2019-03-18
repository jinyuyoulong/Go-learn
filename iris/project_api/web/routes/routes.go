package routes

import (
	"github.com/kataras/iris/mvc"
	"v5u.win/golearn/iris/project_api/bootstrap"
	"v5u.win/golearn/iris/project_api/services"
	"v5u.win/golearn/iris/project_api/web/controllers"
	"v5u.win/golearn/iris/project_api/web/middleware"
)

// Configure registers the necessary routes to the app.
func Configure(b *bootstrap.Bootstrapper) {
	project_apiService := services.Newproject_apiService()

	index := mvc.New(b.Party("/"))
	index.Register(project_apiService)
	index.Handle(new(controllers.IndexController))

	admin := mvc.New(b.Party("/admin"))
	admin.Router.Use(middleware.BasicAuth)
	admin.Register(project_apiService)
	admin.Handle(new(controllers.AdminController))

	// b.Get("/", GetIndexHandler)
	// b.Get("/follower/{id:long}", GetFollowerHandler)
	// b.Get("/following/{id:long}", GetFollowingHandler)
	// b.Get("/like/{id:long}", GetLikeHandler)
}
