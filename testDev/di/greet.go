package main

import (
	"fmt"
	"io"
	"net/http"
)

//Greet 问候
// 参数 bytes.Buffer 改为 io.Writer 使方法更通用
func Greet(b io.Writer, name string) {
	fmt.Fprintf(b, "hello, %s", name)
}

func myGreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "fan")
}
func main() {
	// Greet(os.Stdout, "dd")
	http.ListenAndServe(":5000", http.HandlerFunc(myGreetHandler))
}
