package routes

import (
	bootstrap2 "github.com/jinyuyoulong/Go-learn/src/iris/superstar/bootstrap"
	services2 "github.com/jinyuyoulong/Go-learn/src/iris/superstar/services"
	controllers2 "github.com/jinyuyoulong/Go-learn/src/iris/superstar/web/controllers"
	middleware2 "github.com/jinyuyoulong/Go-learn/src/iris/superstar/web/middleware"
	"github.com/kataras/iris/v12/mvc"
)

// Configure registers the necessary routes to the app.
func Configure(b *bootstrap2.Bootstrapper) {
	superstarService := services2.NewSuperstarService()

	index := mvc.New(b.Party("/"))
	index.Register(superstarService)
	index.Handle(new(controllers2.IndexController))

	admin := mvc.New(b.Party("/admin"))
	admin.Router.Use(middleware2.BasicAuth)
	admin.Register(superstarService)
	admin.Handle(new(controllers2.AdminController))

	// b.Get("/", GetIndexHandler)
	// b.Get("/follower/{id:long}", GetFollowerHandler)
	// b.Get("/following/{id:long}", GetFollowingHandler)
	// b.Get("/like/{id:long}", GetLikeHandler)
}
