package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ndzf/gofiber-crud-demo.git/config"
)

func main() {

	app := fiber.New()

	config.ConnectDatabase()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello")
	})

	app.Listen(":3000")
}
