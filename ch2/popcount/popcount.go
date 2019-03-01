package main

import "fmt"

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// var pc [256]byte = func()(pc [256]byte){
// 	for i:= range pc {
// 		pc[i] = pc[1/2] + byte(i&1)
// 	}
// 	return
// }()
func main() {
	// count := PopCount(30)
	// count := PopCount_cycle(3)
	// count := PopCount_rightBit(4)
	count := PopCount_lowestBit(2)

	fmt.Printf("%d", count)
}

// Popcount returns the population count(number of set bits) of x.
// “下面的代码定义了一个PopCount函数，用于返回一个数字中含二进制1bit的个数。”
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

/*
“练习 2.3： 重写PopCount函数，用一个循环代替单一的表达式。比较两个版本的性能。”
*/
func PopCount_cycle(x uint64) (n int) {
	// 为什么 i 是 uint 类型
	// 为什么 int(pc[byte(x>>(i*8))]+...) 里面的运算可以用 int() 拆分
	for i := uint(0); i < 8; i++ {
		n += int(pc[byte(x>>(i*8))])
	}
	return n
}

/*
“练习 2.4： 用移位算法重写PopCount函数，每次测试最右边的1bit，然后统计总数。比较和查表算法的性能差异。”
*/
func PopCount_rightBit(x uint64) (n int) {
	for i := uint(0); i < 64; i++ {
		if x&1 != 0 {
			n++
		}
		x >>= 1
	}
	return n
}

/*
“练习 2.5： 表达式x&(x-1)用于将x的最低的一个非零的bit位清零。使用这个算法重写PopCount函数，然后比较性能。”
*/
func PopCount_lowestBit(x uint) (n int) {
	for x != 0 {
		x = x & (x - 1)
		n++
	}
	// 最后一位为1，后面 &(x-1) = 0,由此求得最后一位非0 bit 位
	return n
}
