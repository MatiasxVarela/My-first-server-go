package app

import (
	"database/sql"
	"myFirstServerWithGo/routes"

	"github.com/gofiber/fiber/v2"
)

func CreateApp(db *sql.DB) *fiber.App {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	routes.MainRoute(app)

	return app
}
