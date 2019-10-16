package main

// so 动态链接库调用
// 参考 https://studygolang.com/articles/10163

import "fmt"

/*
#cgo CFLAGS: -I./
#cgo LDFLAGS: -L./ -lhi
#include "hi.h" //非标准c头文件，所以用引号
*/
import "C"

func main() {
	C.hi()
	fmt.Println("Hi, vim-go")
}

// CFLAGS: -I路径 这句话指明头文件所在路径
// LDFLAGS: -L路径 -l名字 指明动态库的所在路径，-Llib -lhi，指明在 lib 下面以及它的名字 hi