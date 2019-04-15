package cmd

import (
	"context"

	"github.com/kataras/iris"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	"v5u.win/add"
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Run the RPC client",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Fatal(runClient())
	},
}

func runClient() error {

	// 建立一个grpc连接
	cc, err := grpc.Dial("0.0.0.0:8000", grpc.WithInsecure())
	if err != nil {
		return err
	}

	// 创建AddService的客户端
	cl := add.NewAddServiceClient(cc)

	app := iris.New()

	app.Get("/:a/:b", func(ctx iris.Context) {
		a, _ := ctx.Params().GetInt64("a")
		b, _ := ctx.Params().GetInt64("b")

		c := context.Background() // top level context
		rs, err := cl.Add(c, &add.AddRequest{A: uint64(a), B: uint64(b)})
		if err != nil {
			ctx.Text(err.Error())
			return
		}
		ctx.JSON(rs)
	})
	return app.Run(iris.Addr(":8100"))
}

func init() {
	rootCmd.AddCommand(clientCmd)
}
