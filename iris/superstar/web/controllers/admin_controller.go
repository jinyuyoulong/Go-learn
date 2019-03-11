package controllers

import (
	"v5u.win/GoLearn/iris/superstar/services"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type AdminController struct {
	Ctx     iris.Context
	Service services.SuperstarService
}

func (c *AdminController) Get() mvc.Result {
	// TODO:
	return nil
}

func (c *AdminController) GetEdit() mvc.Result {
	// TODO:
	return nil
}

func (c *AdminController) PostSave() mvc.Result {
	// TODO:
	return nil
}

func (c *AdminController) GetDelete() mvc.Result {
	// TODO;
	return nil
}
