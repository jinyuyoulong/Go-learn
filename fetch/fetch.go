package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {

	for _, url := range os.Args[1:] {
		// 容错 http:// 前缀
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		response, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// body, err := ioutil.ReadAll(response.Body)
		var status string = response.Status
		response.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v \n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", status)
	}
}
