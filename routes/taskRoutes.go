package routes

import (
	"myFirstServerWithGo/service"

	"github.com/gofiber/fiber/v2"
)

type TaskBody struct {
	Name string `json:"name"`
}

func TasksRoutes(tasks fiber.Router) {

	/* Get all tasks */
	tasks.Get("/", service.Find)

	/* Post tasks */
	tasks.Post("/", service.Create)

	/* Get tasks by id */
	tasks.Get("/:id", service.FindOne)
}
