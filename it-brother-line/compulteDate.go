package main

import (
	"fmt"
)

// 封装一个函数，判断某一年份是不是闰年
//闰年的算法：整4不整百，或者整400
func isLeapYear(theYear int) bool {
	var isLeapYear bool
	if theYear%4 == 0 && theYear%100 != 0 {
		isLeapYear = true
		return isLeapYear
	}
	if theYear%400 == 0 {
		isLeapYear = true
		return isLeapYear
	}
	return isLeapYear
}
func main() {
	var scanf int
	for {
		fmt.Scanln(&scanf)
		fmt.Println(isLeapYear(scanf))
	}
}
