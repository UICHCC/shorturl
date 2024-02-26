package main

import (
	"github.com/DRJ31/shorturl-go/controller"
	"github.com/DRJ31/shorturl-go/service"
	"github.com/DRJ31/shorturl-go/util"
)

func main() {
	app := service.InitApp()
	util.InitSnowflake()
	controller.InitRouter(app)
	server := service.GetServerInfo()
	_ = app.Listen(server.String())
}
