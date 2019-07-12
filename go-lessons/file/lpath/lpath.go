package lpath

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var RootPath string = GetRootDirectory()

// 返回程序开始执行(命令执行所在目录，比如 shall 执行所在的目录) 所在目录的上一级目录
// 例：/Users/xxx/dev/go/projectweb ps.execute 文件在 projectweb/bin 目录下
func GetRootDirectory() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return path.Dir(wd)
}

// 返回的是可执行文件所在的目录
func GetCurrentDirectory() string {
	args0 := os.Args[0]
	argsDir := filepath.Dir(args0)
	dir, err := filepath.Abs(argsDir)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}
