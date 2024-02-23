package main

import (
	"fmt"
	"github.com/DRJ31/shorturl-go/controller"
	"github.com/DRJ31/shorturl-go/util"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/template/django/v3"
)

func initRouter(app *fiber.App) {
	app.Static("/static", "./static")
	app.Static("/404", "./static/404")
	app.Get("/", controller.Home)
	app.Get("/:short", controller.Redirect)
	app.Post("/short", controller.Generate)
}

func main() {
	engine := django.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(compress.New())
	util.InitSnowflake()
	initRouter(app)
	cf := util.GetConfig()
	_ = app.Listen(fmt.Sprintf("%v:%v", cf.Server.Host, cf.Server.Port))
}
