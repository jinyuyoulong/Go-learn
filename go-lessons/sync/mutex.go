package main

import (
	"fmt"
	"sync"
)

func main() {
	sw := sync.WaitGroup{}
	sw.Add(1)
	go func() {
		fmt.Println("do it")
		sw.Done()
	}()
	sw.Wait()
	fmt.Println("main over")

	//-----------------
	Deposit(3)
	b := Balance()
	fmt.Println("b =", b)
}

var (
	mu      sync.Mutex // guards balance
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	balance = balance + amount
	mu.Unlock()
}

func Balance() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}