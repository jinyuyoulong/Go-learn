package main

// cgo 直接引用 c/c++ 文件的形式
// 测试通过, .c .h 文件不能和 .go 文件在同一个目录下

import (
	"github.com/jinyuyoulong/Go-learn/go-lessons/cgo/use-cfile/util"
)

func main() {
	util.GoSum(4, 5)
}

