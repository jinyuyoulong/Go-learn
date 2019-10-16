package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

func main() {
	// for {
	// 	fmt.Println("请输入 path")
	// 	var input string
	// 	fmt.Scan(&input)
	// 	subpath := getPathFirstItem(input)
	// 	fmt.Println(subpath)
	// }

	//input := "/org/category/2019-12-09/a3/a38b.jpg"
	//filePath := GetFilePath(input)
	//filePath := getPathFirstItem(input)
	//fmt.Println(filePath)


	s, _ := exec.LookPath(os.Args[0])// 获取可执行文件的路径 执行程序所在路径： 执行程序文件相对路径：
	path := s
	fmt.Println(path)
}

func GetFilePath(fpath string) string {
	// subpath := path.Base(fpath) // a38bjpg
	// subpath2 := filepath.Base(fpath) // filepath包，兼容各操作系统的文件路径
	// subpath := path.Dir(fpath) // /org/category/2019-12-09/a3

	subpath, _ := path.Split(fpath)

	return subpath
}

func getPathFirstItem(fpath string) string {
	// subpath := path.Base(fpath) // a38bjpg
	// subpath2 := filepath.Base(fpath) // filepath包，兼容各操作系统的文件路径
	subpath := path.Dir(fpath) // 结果为：/org/category/2019-12-09/a3
	// subpath := path.Split(fpath)
	// fmt.Println(subpath)
	return subpath
}
