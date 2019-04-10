package mpackage

import (
	"log"
	"os"
	"path"
)

func GetRootDirectory() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return path.Dir(wd)
}
