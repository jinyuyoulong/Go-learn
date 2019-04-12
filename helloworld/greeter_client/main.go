package main

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	pd "v5u.win/add/helloworld"
)

const (
	address     = "localhost:8000"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("grpc connction %v", err)
		return
	}
	defer conn.Close()
	c := pd.NewGreeterClient(conn)
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pd.HelloRequest{Name: name})
	if err != nil {
		log.Fatal("can not get :%v", err)
		return
	}
	log.Printf("greeting %v", r.Message)
}
