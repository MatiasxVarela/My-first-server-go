package routes

import (
	"myFirstServerWithGo/service"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(users fiber.Router) {

	/* Post user */
	users.Post("/", service.CreateUser)

	/* Get all users */
	users.Get("/", service.FindUser)

	/* Get user by id */
	users.Get("/:id", service.FindOneUser)

	/* Delete user */
	users.Delete("/:id", service.DeleteUser)

	/* Update user */
	users.Patch("/:id", service.UpdateUser)

}
