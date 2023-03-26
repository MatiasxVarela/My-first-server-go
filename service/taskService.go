package service

import (
	"database/sql"
	"myFirstServerWithGo/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

/* Create new task */
func CreateTask(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		task := new(models.Task)
		if err := c.BodyParser(task); err != nil {
			return err
		}
		userIdInt, _ := strconv.Atoi(task.UserId)

		query := "INSERT INTO tasks (name, userid) VALUES ($1, $2);"
		_, err := db.Exec(query, task.TaskName, userIdInt)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"message": "Task Created",
		})
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
