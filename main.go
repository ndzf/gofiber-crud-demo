package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ndzf/gofiber-crud-demo.git/config"
	"github.com/ndzf/gofiber-crud-demo.git/routes"
)

func main() {

	app := fiber.New()
	config.ConnectDatabase()
	routes.SetupApiRoutes(app)
	app.Listen(":3000")
}
