package config

import (
	"fmt"

	"github.com/pelletier/go-toml"
	// "github.com/BurntSushi/toml"
)

var (
	Conf = New()
)

func New() *toml.Tree {
	config, err := toml.LoadFile("./src/app/config/config.toml")
	if err != nil {
		fmt.Println("Toml Error!", err.Error())
	}
	return config
}
