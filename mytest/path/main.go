package main

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func main() {
	// println(os.Args[0])
	wd, _ := os.Getwd()
	println(wd)
	// println(path.Dir(wd))
	// println(GetCurrentDirectory())
}
func GetRootDirectory() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return path.Dir(wd)
}

// 返回的是可执行文件所在的目录
// go run 的可执行文件地址是 tmp 目录: /var/folders/zt/55wjgcps3qn_mc_3ymr88d5r0000gn/T/go-build170132262/b001/exe
func GetCurrentDirectory() string {
	dic := filepath.Dir(os.Args[0]) //filepath.Dir(os.Args[0])去除最后一个元素的路径
	dir, err := filepath.Abs(dic)   //返回绝对路径
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}
