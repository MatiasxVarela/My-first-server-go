package routes

import (
	"database/sql"
	"myFirstServerWithGo/service"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(users fiber.Router, db *sql.DB) {

	/* Post user */
	users.Post("/", service.CreateUser(db))

	/* Get all users */
	users.Get("/", service.FindUser)

	/* Get user by id */
	users.Get("/:id", service.FindOneUser)

	/* Delete user */
	users.Delete("/:id", service.DeleteUser)

	/* Update user */
	users.Patch("/:id", service.UpdateUser)

}
