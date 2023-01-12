package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ndzf/gofiber-crud-demo.git/config"
	"github.com/ndzf/gofiber-crud-demo.git/models"
)

func GetTasks(c *fiber.Ctx) error {
	db := config.DB
	var task []models.Task

	db.Find(&task)
	return c.JSON(fiber.Map{
		"success": true,
		"message": "Success get all tasks data",
		"data":    task,
	})
}
