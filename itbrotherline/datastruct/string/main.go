package main 
import "fmt"
func main() {
	var str = `package main 
import "fmt"
func main() {`
	// 反引号 按字符串原样输出
	fmt.Println(str)

// 长字符串则行，+ 号留在一行最后，否则报错。因为go 默认会在最后添加分号
str2 := "hello" + "world"+"hello" + "world"+
"hello" + "world"+
"hello" + "world"
fmt.Println(str2)
}