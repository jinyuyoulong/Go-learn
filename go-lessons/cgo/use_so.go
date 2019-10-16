package main

// LDFLAGS 路径问题并未测试通过

/*
#cgo CFLAGS: -I./
#cgo LDFLAGS: -L./ -lhello
#include "hello.h"
*/
import "C"

func main() {
	C.hello()
}

// CFLAGS: -I路径 这句话指明头文件所在路径，-Iinclude 指明 当前项目根目录的 include 文件夹
// LDFLAGS: -L路径 -l名字 指明动态库的所在路径，-Llib -llibhello，指明在 lib 下面以及它的名字 hello
// 如果动态库不存在，将会爆: 找不到定义之类的错误信息


// 作者：林冠宏
// 链接：https://juejin.im/post/5a62f7cff265da3e4c07e0ab
// 来源：掘金
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
