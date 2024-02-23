package main

import (
	"fmt"
	"github.com/DRJ31/shorturl-go/controller"
	"github.com/DRJ31/shorturl-go/util"
	"github.com/gofiber/fiber/v2"
)

func initRouter(app *fiber.App) {
	app.Static("/static", "./static")
	app.Static("/404", "./static/404")
	app.Get("/", controller.HomePage)
	app.Get("/manual", controller.ManualPage)
	app.Get("/:short", controller.Redirect)
	app.Post("/manual", controller.Manual)
	app.Post("/short", controller.Generate)
}

func main() {
	app := util.InitApp()
	server := util.GetServerInfo()
	initRouter(app)
	_ = app.Listen(fmt.Sprintf("%v:%v", server.Host, server.Port))
}
