// file: controllers/user_controller.go

package controllers

import (
	"github.com/kataras/iris/mvc"
	"v5u.win/GoLearn/iris/superstar/services"

	"github.com/kataras/iris"
)

type IndexController struct {
	Ctx     iris.Context
	Service services.SuperstarService
}

func (c *IndexController) Get() mvc.Result {
	// TODO:
	return nil
}

func (c *IndexControlller) GetBy(id int) mvc.Result {
	// TODO:
	return nil
}

func (c *IndexController) GetSearch() mvc.Result {
	// TODO:
	return nil
}
