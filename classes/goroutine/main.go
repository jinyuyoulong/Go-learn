package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// goroutine()
	// gorange()
	// cacheChannel()
	// cacheChannel1()
	// goselect()
	goserverClient()
}

// goroutine 最简单的使用
func goroutine() {
	// go 通过channel 共享内存
	// channel 通过make创建，close关闭
	// channel 是引用类型
	c := make(chan bool)
	go func() {
		fmt.Println("Go GO GO!")
		c <- true
	}()
	<-c // 程序执行到这，开始阻塞。知道 c 接收到值，即上面的 go func fmt 执行完
}

// 通过for range迭代操作 channel
func gorange() {
	c := make(chan bool)
	go func() {
		fmt.Println("Go GO GO!")
		c <- true
		close(c) // 关闭 channel，防止死锁
	}()
	for v := range c {
		fmt.Println(v)
	}
}

// 单向通道为了防止误操作（读写）
func cacheChannel() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 设置利用CPU多核

	c := make(chan bool, 10) //设置缓存
	for index := 0; index < 10; index++ {
		go Go(c, index)
	}

	for index := 0; index < 10; index++ {
		<-c
	}
	// 无缓存是，此处堵塞，必须等到值被读出。有缓存时，就不管了
	// 有缓存是异步的，无缓存是同步阻塞的
}
func Go(c chan bool, index int) {
	a := 1
	for index := 0; index < 100000; index++ {
		a += index
	}
	fmt.Println(index, c)
	c <- true
}

// 另一种解决方式 wait done
// 通过同步包，实现多个goroutine打印内容
func cacheChannel1() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 设置利用CPU多核
	wg := sync.WaitGroup{}
	wg.Add(10)

	for index := 0; index < 10; index++ {
		go Go1(&wg, index)
	}

	wg.Wait()

	// 无缓存是，此处堵塞，必须等到值被读出。有缓存时，就不管了
	// 有缓存是异步的，无缓存是同步阻塞的
}
func Go1(wg *sync.WaitGroup, index int) {
	a := 1
	for index := 0; index < 100000; index++ {
		a += index
	}
	fmt.Println(index, a)
	wg.Done() // 每调一次，标记完成一次，信号量 -1
}

func goselect() {
	c1, c2 := make(chan int), make(chan string)
	o := make(chan bool, 2)
	go func() {
		for {
			select {
			case v, ok := <-c1:
				if !ok {
					o <- true
					fmt.Println("c1", v)
					break
				}
				fmt.Println(v)

			case v, ok := <-c2:
				if !ok {
					o <- true
					fmt.Println("c2", v)
					break
				}
				fmt.Println(v)
			}
		}
	}()

	c1 <- 1
	c2 <- "hello"
	c1 <- 20
	c2 <- "xxx"
	close(c1)
	// close(c2)
	// 如果判断两个都是否关闭是不可行的，我们只能对其中一个进行判断。在这个🌰中
	for i := 0; i < 2; i++ {
		<-o
	}

}
func goserverClient() {
	c = make(chan string)
	go Pingpong()
	for i := 0; i < 10; i++ {
		c <- fmt.Sprintf("from main:Hi \t #%d\n", i)
		fmt.Printf(<-c)
	}
}

var c chan string

func Pingpong() {

	i := 0
	for {
		fmt.Println(<-c)
		c <- fmt.Sprintf("from pingpong: Hi \t #%d\n", i)
		i++
	}
}
