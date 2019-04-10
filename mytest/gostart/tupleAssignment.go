// 元组赋值
package gostart

import "fmt"

// 求最大公约数
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
		fmt.Printf("x=%d y=%d\n", x, y)
	}
	return x
}

// 寻找第n个斐波那契数 Fibonacci
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func test() {
	// x := gcd(6, 9)
	n := fib(9)
	fmt.Println(n)
}
