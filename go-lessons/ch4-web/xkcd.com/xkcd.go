package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// json 结构体 命名必须首字母大写，成员属性也要首字母大写，注意
type XkcdResult struct {
	Month      int
	Num        int
	Link       string
	Year       int
	News       string
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Alt        string
	Img        string
	Title      string
	Day        int
}
type Xvcd struct {
	Title      string
	Img        string
	Transcript string
}

/*
练习 4.12： 流行的web漫画服务xkcd也提供了JSON接口。例如，一个 https://xkcd.com/571/info.0.json 请求将返回一个很多人喜爱的571编号的详细描述。下载每个链接（只下载一次）然后创建一个离线索引。编写>
一个xkcd工具，使用这些离线索引，打印和命令行输入的检索词相匹配的漫画的URL。
*/
func main() {
	xkcdurl := "https://xkcd.com/571/info.0.json"
	fetch(xkcdurl)
}
func fetch(url string) {
	var result XkcdResult
	resp, _ := http.Get(url)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		os.Exit(1)
	}
	json.NewDecoder(resp.Body).Decode(&result)
	// titles := strings.Split(result.Transcript, " ")
	fmt.Printf("%s - %s\n\n", result.Img, result.Title)
	// fmt.Println(result.Title)
}
