package controller

import "github.com/gofiber/fiber/v2"

func InitRouter(app *fiber.App) {
	app.Static("/static", "./static")
	app.Get("/", HomePage)
	app.Get("/404", PageNotFound)
	app.Get("/manual", ManualPage)
	app.Get("/:short", Redirect)
	app.Post("/manual", Manual)
	app.Post("/short", Generate)
}
