package timer

// 循环任务 time.ticker

import (
	"fmt"
	"time"
)

func EnterFunction()  {
	ticker := time.NewTicker(1 * time.Second)
	isStop := false

	go func() {
		time.Sleep(5*time.Second)
		ticker.Stop()
		isStop = true
	}()

	for  {
		if isStop {
			break
		}
		//每次阻塞一秒
		<-ticker.C
		fmt.Println("我要去浪")
	}
}
