package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	 srw := sync.RWMutex{}
	 sw := sync.WaitGroup{}
	sw.Add(1)
	for i := 0; i < 2; i++ {
		srw.Lock()
			go func() {
				time.Sleep(1*time.Second)
				fmt.Println("write something")
			}()

		srw.Unlock()
	}
	sw.Done()
	time.Sleep(1*time.Second)
	sw.Add(1)

	for i := 0; i < 3; i++ {
		srw.RLock()
		go func() {
			time.Sleep(1*time.Second)
			fmt.Println("read some")
		}()
		defer srw.RUnlock()
	}

	sw.Done()
	sw.Wait()
	time.Sleep(3*time.Second)
	fmt.Println("main over")
}

//var mu sync.RWMutex
//var balance int
//func Balance2() int {
//	mu.RLock() // readers lock
//	defer mu.RUnlock()
//	return balance
//}