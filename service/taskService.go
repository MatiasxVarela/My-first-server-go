package service

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

/* Create new task */
func CreateTask(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON("")
	}
}

/* Find all tasks */
func FindTask(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON("Find all task")
	}
}

/* Find tasks by id */
func FindOneTask(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON("Find task by id")
	}
}

/* Update task */
func UpdateTask(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON("Update Task")
	}
}

/* Delete Task */
func DeleteTask(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON("Delete Task")
	}
}

/* (db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON("")
	}
} */
