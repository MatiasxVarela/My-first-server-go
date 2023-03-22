package routes

import "github.com/gofiber/fiber/v2"

func MainRoute(app *fiber.App) {

	users := app.Group("/api/v1/users")
	UserRoutes(users)

	tasks := app.Group("/api/v1/tasks")
	TasksRoutes(tasks)

}
