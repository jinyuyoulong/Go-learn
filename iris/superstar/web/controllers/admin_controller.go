package controllers

import (
	"fmt"
	"log"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"v5u.win/golearn/iris/superstar/models"
	"v5u.win/golearn/iris/superstar/services"
)

type AdminController struct {
	Ctx     iris.Context
	Service services.SuperstarService
}

func (c *AdminController) Get() mvc.Result {
	datalist := c.Service.GetAll()

	return mvc.View{
		Name: "admin/index.html",
		Data: iris.Map{
			"Title":    "管理后台",
			"Datalist": datalist,
		},
		Layout: "admin/layout.html",
	}
}

func (c *AdminController) GetEdit() mvc.Result {
	id, err := c.Ctx.URLParamInt("id")
	var data *models.StarInfo
	if err == nil {
		data = c.Service.Get(id)
	} else {
		// 添加，id = -1，必须要设置 Id.否则模板无法判断
		data = &models.StarInfo{
			Id: 0,
		}
	}
	// if id == -1 {
	// 	fmt.Println("id is null")
	// 	// data.Id = 0 data= <nil>
	// }
	fmt.Println("edit data: ", id, data)

	return mvc.View{
		Name: "admin/edit.html",
		Data: iris.Map{
			"Title": "管理后台",
			"info":  data,
		},
		Layout: "admin/layout.html", // 不要跟前端的layout混用
	}
}

// 上传数据
func (c *AdminController) PostSave() mvc.Result {
	info := models.StarInfo{}
	err := c.Ctx.ReadForm(&info) // 结合 models 中填写的 form 信息 使用
	if err != nil {
		log.Fatal(err)
	}
	if info.Id > 0 { // 更新
		info.SysUpdated = int(time.Now().Unix())
		c.Service.Update(&info, []string{"name_zh", "name_en", "avatar", "birthday", "height", "weight", "club", "jersy", "country", "moreinfo"})
	} else { // 创建
		info.SysCreated = int(time.Now().Unix())
		c.Service.Create(&info)
	}

	return mvc.Response{
		Path: "/admin/",
	}
}

func (c *AdminController) GetDelete() mvc.Result {
	id, err := c.Ctx.URLParamInt("id")
	if err == nil {
		c.Service.Delete(id)
	}

	return mvc.Response{
		Path: "/admin/",
	}
}
