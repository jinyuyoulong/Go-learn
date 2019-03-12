// file: controllers/user_controller.go

package controllers

import (
	"github.com/kataras/iris/mvc"
	"v5u.win/golearn/iris/superstar/services"

	"github.com/kataras/iris"
)

type IndexController struct {
	Ctx     iris.Context
	Service services.SuperstarService
}

func (c *IndexController) Get() mvc.Result {
	datalist := c.Service.GetAll()

	return mvc.View{
		Name: "index.html",
		Data: iris.Map{
			"Title":    "球星库",
			"Datalist": datalist,
		},
	}
}

func (c *IndexController) GetBy(id int) mvc.Result {
	if id < 0 {
		return mvc.Response{
			Path: "/",
		}
	}
	data := c.Service.Get(id)
	return mvc.View{
		Name: "info.html",
		Data: iris.Map{
			"Title": "球星库",
			"info":  data,
		},
	}
}

func (c *IndexController) GetSearch() mvc.Result {
	country := c.Ctx.URLParam("country")
	datalist := c.Service.Search(country)
	return mvc.View{
		Name: "index.html",
		Data: iris.Map{
			"Title":    "球星库",
			"Datalist": datalist,
		},
	}
}
