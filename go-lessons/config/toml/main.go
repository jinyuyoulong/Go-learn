package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Config struct {
	App   App
	Image Image
}
type App struct {
	Name  string
	Url   string
	Port  string
	Debug bool
}
type Image struct {
	ImageLib      string        `toml:"image_lib"`
	ImagePath     string        `toml:"image_path"`
	ImageURL      string        `toml:"image_ur"`
	ImageOrg      string        `toml:"image_org"`
	ImageTmp      string        `toml:"image_tmp"`
	ImageTypes    []string      `toml:"image_types"`
	WaterMark     string        `toml:"water_mark"`
	ImageCategory imageCategory `toml:"imageCategory"`
	// ImageCategory struct {
	// 	carLogo struct {
	// 		Paths string `toml:"paths"`
	// 		Sizes string `toml:sizes`
	// 	} `toml:"carLogo"`
	// 	imgLogo struct {
	// 		Paths string `toml:"paths"`
	// 	} `toml:"imgLogo"`
	// } `toml:"imageCategory"`
}

type imageCategory struct {
	ImgLogo imgLogo `toml:"imgLogo"`
	CarLogo carLogo `toml:"carLogo"`
}

type carLogo struct {
	Paths string   `toml:"paths"`
	Sizes []string `toml:"sizes"`
}
type imgLogo struct {
	Paths string `toml:paths`
}

func main() {
	file := "../config/config.toml"
	var conf Config
	if _, err := toml.DecodeFile(file, &conf); err != nil {
		fmt.Println("err: ", err.Error())
	}
	fmt.Printf("app: %s\n", conf.App.Url)
	fmt.Printf("ImageLib: %s\n", conf.Image.ImageTypes)
	fmt.Println(conf.Image.ImageCategory)
}
