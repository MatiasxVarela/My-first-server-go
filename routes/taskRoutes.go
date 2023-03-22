package routes

import (
	"myFirstServerWithGo/service"

	"github.com/gofiber/fiber/v2"
)

type TaskBody struct {
	Name string `json:"name"`
}

func TasksRoutes(tasks fiber.Router) {

	/* Post tasks */
	tasks.Post("/", service.CreateTask)

	/* Get all tasks */
	tasks.Get("/", service.FindTask)

	/* Get task by id */
	tasks.Get("/:id", service.FindOneTask)

	/* Delete task */
	tasks.Delete("/:id", service.DeleteTask)

	/* Update */
	tasks.Patch("/:id", service.UpdateTask)
}
