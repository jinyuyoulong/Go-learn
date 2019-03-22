package controller

import (
	"fmt"
	"log"
	"time"

	"github.com/kataras/iris"
	"v5u.win/golearn/iris/projectapi/src/app/model"
	"v5u.win/golearn/iris/projectapi/src/app/service"
)

type AdminController struct {
	Ctx     iris.Context
	Service service.ProjectapiService
}

func (c *AdminController) Get(ctx iris.Context) {
	var datalist []model.StarInfo
	datalist = c.Service.GetAll()
	ctx.JSON(ApiResult(true, datalist, ""))
}

// remove, not used
func (c *AdminController) GetEdit(ctx iris.Context) {
	id, err := c.Ctx.URLParamInt("id")
	var data *model.StarInfo
	if err == nil {
		data = c.Service.Get(id)
	}
	fmt.Println(id, data)
	ctx.JSON(ApiResult(true, data, ""))
}

// 上传数据
func (c *AdminController) PostSave(ctx iris.Context) {
	info := model.StarInfo{}
	err := c.Ctx.ReadForm(&info) // 结合 model 中填写的 form 信息 使用
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
	ctx.JSON(ApiResult(true, "post data success!", ""))
}

// 假删除 sys_status:1
func (c *AdminController) GetDelete(ctx iris.Context) {
	id, err := c.Ctx.URLParamInt("id")

	if err != nil {
		log.Fatal(err)
		ctx.JSON(ApiResult(false, "", "argument wrong!"))
	}
	err = c.Service.Delete(id)
	fmt.Println(id)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(ApiResult(false, "delete data faild!", err.Error()))
	} else {
		ctx.JSON(ApiResult(true, "delete data success!", ""))
	}
}
