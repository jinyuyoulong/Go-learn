package routes

import (
	"github.com/kataras/iris/mvc"
	"v5u.win/golearn/iris/projectweb/bootstrap"
	"v5u.win/golearn/iris/projectweb/services"
	"v5u.win/golearn/iris/projectweb/web/controllers"
	"v5u.win/golearn/iris/projectweb/web/middleware"
)

// Configure registers the necessary routes to the app.
func Configure(b *bootstrap.Bootstrapper) {
	projectwebService := services.NewprojectwebService()

	index := mvc.New(b.Party("/"))
	index.Register(projectwebService)
	index.Handle(new(controllers.IndexController))

	admin := mvc.New(b.Party("/admin"))
	admin.Router.Use(middleware.BasicAuth)
	admin.Register(projectwebService)
	admin.Handle(new(controllers.AdminController))

	// b.Get("/", GetIndexHandler)
	// b.Get("/follower/{id:long}", GetFollowerHandler)
	// b.Get("/following/{id:long}", GetFollowingHandler)
	// b.Get("/like/{id:long}", GetLikeHandler)
}
