package app

import (
	"myFirstServerWithGo/routes"

	"github.com/gofiber/fiber/v2"
)

func CreateApp() *fiber.App {
	app := fiber.New()
	routes.MainRoute(app)

	return app
}
