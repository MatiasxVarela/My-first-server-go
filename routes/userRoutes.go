package routes

import "github.com/gofiber/fiber/v2"

func UserRoutes(users fiber.Router) {

	users.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Get all users")
	})

	users.Get("/:id", func(c *fiber.Ctx) error {
		return c.SendString("Get user by id")
	})

}
