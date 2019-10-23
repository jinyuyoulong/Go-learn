package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// 依赖注入 模拟
const timeCount = 3
const finalWorld = "GO!"

func CountDown(w io.Writer) {
	for i := 0; i < timeCount; i++ {
		time.Sleep(time.Second * 1)
		fmt.Fprintf(w, "%d\n", timeCount-i)

	}
	time.Sleep(time.Second * 1)
	fmt.Fprint(w, finalWorld)
}
func main() {
	CountDown(os.Stdout)
}
