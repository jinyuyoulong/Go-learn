package timer

// 定时任务 time.Timer

import (
	"fmt"
	"time"
)

func TimeCount() {
	//timer := time.After()
	timer := time.NewTimer(5 * time.Second)
	fmt.Println("计时开始", time.Now())

	time.Sleep(1 * time.Second)

	//// 终止定时器
	//ok := timer.Stop()
	//if ok {
	//	os.Exit(1)
	//}

	// 以当前时间为基准（而不是timer被创建的时间），将定时器重置为2秒
	timer.Reset(2 * time.Second)

	endTime := <-timer.C
	fmt.Println("炸弹引爆", endTime)
}
