package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ndzf/gofiber-crud-demo.git/config"
	"github.com/ndzf/gofiber-crud-demo.git/controllers"
)

func main() {

	app := fiber.New()

	config.ConnectDatabase()

	app.Get("/", controllers.GetTasks)

	app.Listen(":3000")
}
