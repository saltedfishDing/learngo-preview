package conf

import (
	"os"
	"preview/pkg/helper"

	"gopkg.in/yaml.v2"
)

type server struct {
	System   System    `yaml:"system"`
	Database *Database `yaml:"database"`
}

var Server *server

func NewServer(path string) {
	var err error

	_, err = helper.PathExists(path)

	if err != nil {
		panic(err)
	}

	body, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	if err = yaml.Unmarshal(body, &Server); err != nil {
		panic(err)
	}
}
