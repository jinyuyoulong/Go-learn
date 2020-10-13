# golang 中调用 c/c++ 代码有三种方式

- 直接`嵌套`在go文件中使用，最简单直观的
- 导入`动态库 .so 或 dll` 的形式，最安全但是很不爽也比较慢的
- 直接引用 c/c++ 文件的形式，层次分明，容易随时修改看结果的

## 需要的环境支持

- Linux 具备 gcc 与 g++ 即可
- Mac 参考 Linux
- Windows 需要安装 [mingw](https://link.juejin.im/?target=http%3A%2F%2Fwww.mingw.org%2F)，否则编译时会有这类错：`cannot find -lmingwex`

## 嵌套

```go
package main

// 嵌入式 调用

// #include <stdio.h>
// #include <stdlib.h>
/*
void print(char *str) {
    printf("%s\n", str);
}
*/
import "C"//这里可看作封装的伪包C, 这条语句要紧挨着上面的注释块，不可在它俩之间间隔空行！

import "unsafe"

func main() {
	s := "Hello Cgo"
	cs := C.CString(s)
	C.print(cs)
	C.free(unsafe.Pointer(cs))
}

//参考来源：https://tonybai.com/2012/09/26/interoperability-between-go-and-c/
```

## 调用动态链接库 .so

目录结构

```sh
├── hi.c
├── hi.go
├── hi.h
├── hi.o
└── libhi.so
// hi.c
#include <stdio.h>

void hi() {
    printf("Hello Cgo!\n");
}
void hi(); // file: hi.h
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
// 注：LDFLAGS 路径问题并未测试通过
// CFLAGS: -I路径 这句话指明头文件所在路径
// LDFLAGS: -L路径 -l名字 指明动态库的所在路径，-Llib -lhi，指明在 lib 下面以及它的名字 hi
# 编译生成 .so 动态库
gcc hi.c -fPIC -shared -o libhi.so
```

## 直接引用 c/c++ 文件的形式

```sh
.
├── main.go
└── util
    ├── util.go
    └── utilc
        ├── util.c
        └── util.h
```

目录结构

```go
// util.go
package util

/*
#include "utilc/util.c"
*/
import "C"

import "fmt"

func GoSum(a, b int) {
	s := C.sum(C.int(a), C.int(b))
	fmt.Println(s)
}
// main.go
package main

// cgo 直接引用 c/c++ 文件的形式
// 测试通过, .c .h 文件不能和 .go 文件在同一个目录下

import (
	"github.com/xxx/util"
)

func main() {
	util.GoSum(4, 5)
}
// utilc/util.c
#include "util.h"

int sum(int a,int b){
    return (a+b);
}
int sum(int a,int b);// utilc/util.h
```

直接运行`go run main.go`

## 注意事项

- 但凡要引用与 c/c++ 相关的内容，写到 go 文件的头部`注释`里面
- 嵌套的 c/c++ 代码必须符合其语法，不与 go 一样
- `import "C"` 这句话要紧随，注释后，不要换行，否则报错
- go 代码中调用 c/c++ 的格式是: `C.xxx()`，例如 C.add(2, 1)

------

`CFLAGS: -I路径` 这句话指明头文件所在路径，-Iinclude 指明 当前项目根目录的 include 文件夹

`LDFLAGS: -L路径 -l名字` 指明动态库的所在路径，-Llib -llibvideo，指明在 lib 目录下面以及它的名字 video

注：此处的 LDFLAGS 更改动态库的路径后，-L指明路径，无法验证通过