package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ndzf/gofiber-crud-demo.git/controllers"
)

func SetupApiRoutes(app *fiber.App) {
	api := app.Group("/api")

	tasks := api.Group("/tasks")
	tasks.Get("/", controllers.GetTasks)
	tasks.Post("/", controllers.CreateTask)
	tasks.Get("/:taskId", controllers.GetTask)
}
