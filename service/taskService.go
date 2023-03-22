package service

import (
	"github.com/gofiber/fiber/v2"
)

/* Create new task */
func CreateTask(c *fiber.Ctx) error {
	return c.JSON("Create task")
}

/* Find all tasks */
func FindTask(c *fiber.Ctx) error {
	return c.JSON("Find all tasks")
}

/* Find tasks by id */
func FindOneTask(c *fiber.Ctx) error {
	return c.JSON("Find one task")
}

/* Update task */
func UpdateTask(c *fiber.Ctx) error {
	return c.JSON("Update task")
}

/* Delete Task */
func DeleteTask(c *fiber.Ctx) error {
	return c.JSON("Delete task")
}
