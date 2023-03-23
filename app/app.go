package app

import (
	"database/sql"
	"myFirstServerWithGo/routes"

	"github.com/gofiber/fiber/v2"
)

func CreateApp(db *sql.DB) *fiber.App {
	app := fiber.New()

	routes.MainRoute(app, db)

	return app
}
