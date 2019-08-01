package main

import (
	"html/template"
	"log"
	"os"
	"time"

	"github.com/jinyuyoulong/Go-learn/ch4/github"
)

// 定义模板显示格式
const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

// 时间转换
func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

// 模板配置函数
var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	// 使用定义好的 模板 输出到 控制台
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
