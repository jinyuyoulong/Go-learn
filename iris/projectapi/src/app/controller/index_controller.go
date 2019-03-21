// file: controller/user_controller.go

package controller

import (
	"fmt"
	"log"

	"github.com/kataras/iris/mvc"
	"v5u.win/golearn/iris/projectapi/src/app/datasource"
	"v5u.win/golearn/iris/projectapi/src/app/model"
	"v5u.win/golearn/iris/projectapi/src/app/service"

	"github.com/kataras/iris"
)

type IndexController struct {
	Ctx     iris.Context
	Service service.ProjectapiService
}

func (c *IndexController) Get(ctx iris.Context) {
	datalist := c.Service.GetAll()
	// var datalist []model.StarInfo
	// return mvc.Response{Text: "ok\n"}
	ctx.JSON(iris.Map{"message": "Hello iris web framework.",
		"data": datalist})
}

// /?id=2
func (c *IndexController) GetBy(ctx iris.Context, id int) {
	if id < 0 {
		ctx.JSON(iris.Map{"code": 1, "result": ""})
	}
	data := c.Service.Get(id)
	ctx.JSON(iris.Map{"code": 0, "result": data})
}

// /search?country=瑞士
func (c *IndexController) GetSearch(ctx iris.Context) {
	country := c.Ctx.URLParam("country")
	datalist := c.Service.Search(country)
	fmt.Println("search:", datalist)
	ctx.JSON(iris.Map{"code": 0, "result": datalist})
}

// 手动清除缓存 分布式数据同步
func (c *IndexController) GetClearcache() mvc.Result {
	err := datasource.InstanceMaster().ClearCache(&model.StarInfo{})
	if err != nil {
		log.Fatal(err)
	}

	return mvc.Response{
		Text: "xorm 缓存清除成功",
	}
}
