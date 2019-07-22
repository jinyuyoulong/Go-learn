package main

import (
	"github.com/pelletier/go-toml"
)

func main() {
	file := "../config/config.toml"
	config, err := toml.LoadFile(file)
	if err != nil {
		fmt.Println("Toml Error!", err.Error())
	}

	category := config.Get("image.array").([]interface{})
	for value := range category {
		println(value, " aa")
	}
	fmt.Printf("%d %T", len(category), category)
	path := category[0]
}
