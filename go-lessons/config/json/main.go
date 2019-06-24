package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// json解析，结构体字段必须首字母大写
type image struct {
	ImageLib      string         `json:"image_lib"`
	ImagePath     string         `json:"image_path"`
	ImageURL      string         `json:"image_ur"`
	ImageOrg      string         `json:"image_org"`
	ImageTmp      string         `json:"image_tmp"`
	ImageTypes    []string       `json:"image_types"`
	WaterMark     string         `json:"waterMark"`
	ImageCategroy *ImageCategory `json:"image_categroy"`
	CarImages     []string       `json:"car_images"`
}
type ImageCategory struct {
	Paths string
	Sizes []string
}

func main() {
	img := image{}

	data, _ := ioutil.ReadFile("config.json")
	// fmt.Println(data)
	//读取的数据为json格式，需要进行解码
	err := json.Unmarshal(data, &img)
	// err := json.MarshalIndent(data,"","    ")
	if err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Printf("%v\n", img)
	fmt.Printf("%v %p", img.ImageCategroy, &img.ImageCategroy)
}
