package main

// 不定参数 方法定义
import "fmt"

type myFunc func(int) myFunc

var total int

func Sum(nums ...int) myFunc {
	total = 0

	for _, num := range nums {
		total += num
	}

	var f1 myFunc
	f1 = func(a int) myFunc {
		total += a
		return f1
	}

	return f1
}

func main() {
	Sum(2, 3)
	fmt.Println(total)
	Sum(2)(3)
	fmt.Println(total)
	Sum(2)(3)(4)
	fmt.Println(total)
	Sum(2)(3)(4)(5)
	fmt.Println(total)
	Sum(2)(3)(4)(5)(6)
	fmt.Println(total)

}
