// file: controllers/user_controller.go

package controllers

import (
	"log"

	"github.com/kataras/iris/mvc"
	"v5u.win/golearn/iris/project_api/datasource"
	"v5u.win/golearn/iris/project_api/models"
	"v5u.win/golearn/iris/project_api/services"

	"github.com/kataras/iris"
)

type IndexController struct {
	Ctx     iris.Context
	Service services.project_apiService
}

func (c *IndexController) Get() mvc.Result {
	datalist := c.Service.GetAll()
	// var datalist []models.StarInfo
	// return mvc.Response{Text: "ok\n"}
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

// 手动清除缓存
func (c *IndexController) GetClearcache() mvc.Result {
	err := datasource.InstanceMaster().ClearCache(&models.StarInfo{})
	if err != nil {
		log.Fatal(err)
	}

	return mvc.Response{
		Text: "xorm 缓存清除成功",
	}
}
