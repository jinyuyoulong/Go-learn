// 用于参考写命令行工具 如何处理命令行参数，
package gostart

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")

var sep = flag.String("s", " ", "separator")

func echo4() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
