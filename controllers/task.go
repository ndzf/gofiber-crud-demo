package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ndzf/gofiber-crud-demo.git/config"
	"github.com/ndzf/gofiber-crud-demo.git/helpers"
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

func GetTask(c *fiber.Ctx) error {
	db := config.DB
	var task models.Task

	id := c.Params("taskId")
	db.Find(&task, "id = ?", id)

	// Empty task
	if task.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "no task present",
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Task Found",
		"data":    task,
	})
}

func CreateTask(c *fiber.Ctx) error {
	task := new(models.Task)

	if err := c.BodyParser(task); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "review your inputs",
			"data":    err,
		})
	}

	validationErrors := helpers.Validate(*task)
	if validationErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "validation errors",
			"data":    validationErrors,
		})
	}

	// creating notes
	db := config.DB
	if err := db.Create(&task).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "could not create task",
			"data":    err,
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Created task",
		"data":    task,
	})
}
