package main

import (
	"fmt"
	"github.com/pelletier/go-toml"
)

func main() {
	file := "../config/config.toml"
	config, err := toml.LoadFile(file)
	if err != nil {
		fmt.Println("Toml Error!", err.Error())
	}
	category := config.Get("image.image_categroy.car_logo.paths").(string)
	fmt.Println(category)

	// for _, value := range category {
	// 	fmt.Println(value)
	// 	// fmt.Printf("%d %T", len(value), value)
	// }
}
