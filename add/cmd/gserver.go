package cmd

import (
	"net"

	"v5u.win/add"
	"v5u.win/add/server"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run the RPC server",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Fatal(runServe())
	},
}

func runServe() error {
	lis, err := net.Listen("tcp", ":8000")
	logrus.Infof("add serve at %s", "0.0.0.0:8000")
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	add.RegisterAddServiceServer(s, &server.AddServer{})
	return s.Serve(lis)
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
