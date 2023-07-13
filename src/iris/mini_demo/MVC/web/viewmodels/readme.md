视图View Models
应该有视图模型，客户端将能够看到的结构。例：

    import (
        "github.com/kataras/iris/v12/_examples/mvc/overview/datamodels"
        "github.com/kataras/iris/v12/context"
    )
    type Movie struct {
        datamodels.Movie
    }
    func (m Movie) IsValid() bool {
        /* 做一些检查，如果有效则返回true。.. */
        return m.ID > 0
    }
Iris能够将任何自定义数据结构转换为HTTP响应调度程序，因此从理论上讲，如果确实需要，则允许使用以下内容：

    // Dispatch完成`kataras / iris / mvc＃Result`界面。
    //将“电影”作为受控的http响应发送。
    //如果其ID为零或更小，则返回404未找到错误
    //否则返回其json表示，
    //（就像控制器的函数默认为自定义类型一样）。
    //
    //不要过度，应用程序的逻辑不应该在这里。
    //在响应之前，这只是验证的又一步，
    //可以在这里添加简单的检查。
    //
    //这只是一个展示
    //想象一下设计更大的应用程序时此功能给出的潜力。
    //
    //调用控制器方法返回值的函数
    //是“电影”的类型。
    //例如`controllers / movie_controller.go＃GetBy`.
    func (m Movie) Dispatch(ctx context.Context) {
        if !m.IsValid() {
            ctx.NotFound()
            return
        }
        ctx.JSON(m, context.JSON{Indent: " "})
    }
但是，我们将使用“datamodels”作为唯一的模型包，因为Movie结构不包含任何敏感数据，客户端能够查看其所有字段，并且我们不需要任何额外的功能或验证。