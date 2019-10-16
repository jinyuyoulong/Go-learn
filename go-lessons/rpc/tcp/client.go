// client.go
package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
)

type Args struct {
	A, B int
}
type Quotient struct {
	Quo, Rem int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "server")
		os.Exit(1)
	}

	serverAddr := os.Args[1]
	client, err := rpc.Dial("tcp", serverAddr+":1234")
	if err != nil {
		log.Fatal("dialing", err)
	}

	args := Args{7, 9}
	var reply int
	err = client.Call("Math.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Math error:", err)
	}

	fmt.Printf("Math: %d * %d= %d\n", args.A, args.B, reply)

	var quo Quotient
	err = client.Call("Math.Divide", args, &quo)
	if err != nil {
		log.Fatal("Math error:", err)
	}

	fmt.Printf("Math: %d ÷ %d= %d 余 %d\n", args.A, args.B, quo.Quo, quo.Rem)
}
