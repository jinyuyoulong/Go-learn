package main

import (
	"fmt"
	"time"
)

// 依赖注入 模拟

func CountDown() {
	for i := 0; i < 3; i++ {
		fmt.Println(3 - i)
		time.Sleep(time.Second * 1)
		if i == 2 {
			fmt.Println("GO!")
		}
	}
}
func main() {
	CountDown()
}
