package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// 依赖注入 模拟

func CountDown(w io.Writer) {
	for i := 0; i < 3; i++ {
		fmt.Fprintf(w, "%d\n", 3-i)
		time.Sleep(time.Second * 1)
		if i == 2 {
			fmt.Fprint(w, "GO!")
		}
	}
}
func main() {
	CountDown(os.Stdout)
}
