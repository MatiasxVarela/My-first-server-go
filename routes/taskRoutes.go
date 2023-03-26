package routes

import (
	"database/sql"
	"myFirstServerWithGo/service"

	"github.com/gofiber/fiber/v2"
)

type TaskBody struct {
	Name string `json:"name"`
}

func TasksRoutes(tasks fiber.Router, db *sql.DB) {

	/* Post tasks */
	tasks.Post("/", service.CreateTask(db))

	/* Get all tasks */
	tasks.Get("/", service.FindTask(db))

	/* Get task by id */
	tasks.Get("/:id", service.FindOneTask(db))

	/* Delete task */
	tasks.Delete("/:id", service.DeleteTask(db))

	/* Update */
	tasks.Patch("/:id", service.UpdateTask(db))
}
