package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		response, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		// _, err := io.Copy(os.Stdout, response.Body)
		_, err = io.Copy(os.Stdout, resp.Body) //在这里输出内容到标准输出
		response.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v \n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", os.Stdout)
	}
}
