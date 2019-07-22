package main

import (
	"fmt"
	// "regexp"
	"strings"
)

func main() {
	for {
		fmt.Println("请输入 path")
		var input string
		fmt.Scan(&input)
		fmt.Println(input)
		category := GetImageCategoryByPath(input)
		fmt.Println(category)
	}
}

func GetImageCategoryByPath(path string) string {
	// /org/category/2019-12-09/a3/a38b.jpg
	// categorys := regexp.MustCompile(`/[a-zA-Z0-9]+`).FindAllString(path, -1)
	categorys := strings.Split(path, "/")
	fmt.Println(len(categorys))
	for _, value := range categorys {
		fmt.Println("value:", value)
	}
	if len(categorys) > 2 {
		return categorys[1]
	}
	return ""
}
