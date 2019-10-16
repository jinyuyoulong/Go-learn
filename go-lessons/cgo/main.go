package main

// #include <stdio.h>
// #include <stdlib.h>
/*
#cgo CFLAGS: -I./
#cgo LDFLAGS: -L./ -lhello
#include "hello.h" //非标准c头文件，所以用引号

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

	C.hello()
}

//参考来源：https://tonybai.com/2012/09/26/interoperability-between-go-and-c/