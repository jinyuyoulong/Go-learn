package make_file

import (
	"os"
)

func main() {
	// path := "./cfile"
	// err := os.Mkdir(path, 0777)

	// 创建
	err := os.MkdirAll("./afile/test", 0666)

	if err != nil {
		println(err.Error())
	}

}
