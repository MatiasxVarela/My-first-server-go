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
		var task = []models.Task{}

		query := "SElECT * FROM tasks"

		rows, err := db.Query(query)
		if err != nil {
			return err
		}
		defer rows.Close()

		for rows.Next() {
			var responseTasks models.Task
			if err := rows.Scan(&responseTasks.Id, &responseTasks.TaskName, &responseTasks.UserId); err != nil {
				return err
			}
			task = append(task, responseTasks)
		}

		if err := rows.Err(); err != nil {
			return err
		}

		if len(task) > 0 {
			return c.JSON(task)
		} else {
			return c.Status(404).JSON(fiber.Map{
				"message": "Tasks not found",
			})
		}
	}
}

/* Find tasks by id */
func FindOneTask(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		idInt, _ := strconv.Atoi(id)
		var task models.Task

		query := "SElECT * FROM tasks WHERE id=$1"
		err := db.QueryRow(query, idInt).Scan(&task.Id, &task.TaskName, &task.UserId)
		if err != nil {
			return c.Status(404).JSON(fiber.Map{
				"message": "Task not found",
			})
		}

		return c.JSON(task)
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
