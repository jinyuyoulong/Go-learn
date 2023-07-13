// file: controllers/user_controller.go

package controllers

import (
	"log"

	datasource2 "github.com/jinyuyoulong/Go-learn/src/iris/superstar/datasource"
	models2 "github.com/jinyuyoulong/Go-learn/src/iris/superstar/models"
	services2 "github.com/jinyuyoulong/Go-learn/src/iris/superstar/services"

	"github.com/kataras/iris/v12/mvc"

	"github.com/kataras/iris/v12"
)

type IndexController struct {
	Ctx     iris.Context
	Service services2.SuperstarService
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
	err := datasource2.InstanceMaster().ClearCache(&models2.StarInfo{})
	if err != nil {
		log.Fatal(err)
	}

	return mvc.Response{
		Text: "xorm 缓存清除成功",
	}
}
