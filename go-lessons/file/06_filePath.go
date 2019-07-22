package main

import (
	"fmt"
	"path"
	// "path/filepath"
)

func main() {
	// for {
	// 	fmt.Println("请输入 path")
	// 	var input string
	// 	fmt.Scan(&input)
	// 	subpath := getPathFirstItem(input)
	// 	fmt.Println(subpath)
	// }
	input := "/org/category/2019-12-09/a3/a38b.jpg"
	// subpath := path.Base(fpath) // a38bjpg
	// subpath2 := filepath.Base(fpath) // filepath包，兼容各操作系统的文件路径
	// subpath := path.Dir(fpath) // /org/category/2019-12-09/a3
	subpath, _ := path.Split(input)
	fmt.Println(subpath)
}

func getPathFirstItem(fpath string) string {
	// subpath := path.Base(fpath) // a38bjpg
	// subpath2 := filepath.Base(fpath) // filepath包，兼容各操作系统的文件路径
	subpath := path.Dir(fpath) // /org/category/2019-12-09/a3
	// subpath := path.Split(fpath)
	// fmt.Println(subpath)
	return subpath
}
