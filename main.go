package main

import (
	"preview/conf"
	"preview/models"
	"preview/routers"
)

func main() {
	conf.NewServer("./conf.yaml")

	models.Init(conf.Server.Database)

	routers.SetupRouter().Run(conf.Server.System.Address())
}
