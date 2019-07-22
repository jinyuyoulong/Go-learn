package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	// 通过channel 共享内存
	ch := make(chan string) // 创建通道 channel
	// 循环调用 fetch 函数
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // gp function 表示创建一个新的 gorouting，并在这个新的 goroutine 中执行这个函数
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
		/*
			当一个goroutine尝试在一个channel上做send或者receive操作时，
			这个goroutine会阻塞在调用处，直到另一个goroutine往这个channel里写入、或者接收值，
			这样两个goroutine才会继续执行channel操作之后的逻辑。
			在这个例子中，每一个fetch函数在执行时都会往channel里发送一个值(ch <- expression)，主函数负责接收这些值(<-ch)。
			这个程序中我们用main函数来接收所有fetch函数传回的字符串，
			可以避免在goroutine异步执行还没有完成时main函数提前退出。
		*/
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	response, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	// nbytes, err := io.Copy(ioutil.Discard, response.Body) // ioutil.Discard 向里面写一些不需要的数据，我们需要他的字节数，但是不要其内容
	// 1.10 练习
	nbytes, err := io.Copy(os.Stdout, response.Body)
	response.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	// fmt.Printf("%s\n", os)
	seconds := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", seconds, nbytes, url)
}

/*
How to use
go build gopl.io/ch1/fetchall
$ ./fetchall https://golang.org http://gopl.io https://godoc.org
*/
