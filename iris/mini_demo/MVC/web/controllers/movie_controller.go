package controllers

import (
	"errors"
	datamodels2 "github.com/jinyuyoulong/Go-learn/iris/mini_demo/MVC/datamodels"
	services2 "github.com/jinyuyoulong/Go-learn/iris/mini_demo/MVC/services"

	"github.com/kataras/iris"
)

// MovieController
type MovieController struct {
	//我们的MovieService，它是一个界面
	//从主应用程序绑定。
	Service services2.MovieService
}

// 获取电影列表
// curl -i http://localhost:8080/movies
// 如果您有敏感数据，这是正确的方法：
// func (c *MovieController) Get() (results []viewmodels.Movie) {
//     data := c.Service.GetAll()
//     for _, movie := range data {
//         results = append(results, viewmodels.Movie{movie})
//     }
//     return
// }
// 否则只返回数据模型
func (c *MovieController) Get() (results []datamodels2.Movie) {
	return c.Service.GetAll()
}

//获取一部电影
// Demo:
// curl -i http://localhost:8080/movies/1
func (c *MovieController) GetBy(id int64) (movie datamodels2.Movie, found bool) {
	return c.Service.GetByID(id) // it will throw 404 if not found.
}

// 用put请求更新一部电影
// Demo:
// curl -i -X PUT -F "genre=Thriller" -F "poster=@/Users/kataras/Downloads/out.gif" http://localhost:8080/movies/1
func (c *MovieController) PutBy(ctx iris.Context, id int64) (datamodels2.Movie, error) {
	// get the request data for poster and genre
	file, info, err := ctx.FormFile("poster")
	if err != nil {
		return datamodels2.Movie{}, errors.New("failed due form file 'poster' missing")
	}
	// 不需要文件所以关闭他
	file.Close()
	//想象一下，这是上传文件的网址......
	poster := info.Filename
	genre := ctx.FormValue("genre")
	return c.Service.UpdatePosterAndGenreByID(id, poster, genre)
}

// Delete请求删除一部电影
// curl -i -X DELETE -u admin:password http://localhost:8080/movies/1
func (c *MovieController) DeleteBy(id int64) interface{} {
	wasDel := c.Service.DeleteByID(id)
	if wasDel {
		// 返回删除的id
		return iris.Map{"deleted": id}
	}
	//在这里我们可以看到方法函数可以返回这两种类型中的任何一种（map或int），
	//我们不必将返回类型指定为特定类型。
	return iris.StatusBadRequest
}
