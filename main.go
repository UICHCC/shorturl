package main

import (
	"github.com/DRJ31/shorturl-go/controller"
	"github.com/DRJ31/shorturl-go/service"
	"github.com/gofiber/fiber/v2"
)

func initRouter(app *fiber.App) {
	app.Static("/static", "./static")
	app.Get("/", controller.HomePage)
	app.Get("/404", controller.PageNotFound)
	app.Get("/manual", controller.ManualPage)
	app.Get("/:short", controller.Redirect)
	app.Post("/manual", controller.Manual)
	app.Post("/short", controller.Generate)
}

func main() {
	app := service.InitApp()
	initRouter(app)
	server := service.GetServerInfo()
	_ = app.Listen(server.String())
}
