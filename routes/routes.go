package routes

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func MainRoute(app *fiber.App, db *sql.DB) {

	users := app.Group("/api/v1/users")
	UserRoutes(users, db)

	tasks := app.Group("/api/v1/tasks")
	TasksRoutes(tasks, db)

}
