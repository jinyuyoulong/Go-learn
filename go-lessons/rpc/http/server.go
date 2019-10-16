// server.go
package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/rpc"
)

type Args struct {
	A, B int
}

// func () NAME(ARGS TYPE,REPLY TYPE)error
type Math int

// 乘法
func (m *Math) Multiply(args Args, reply *int) error {
	*reply = args.A * args.B
	return nil

}

type Quotient struct {
	Quo, Rem int
}

func (m *Math) Divide(args Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil

}
func main() {
	math := new(Math)
	rpc.Register(math)
	rpc.HandleHTTP()
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
