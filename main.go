package main

import (
	"myFirstServerWithGo/app"
	"myFirstServerWithGo/db"

	_ "github.com/lib/pq"
)

func main() {
	db.OpenDb(true)
	server := app.CreateApp()

	server.Listen(":3000")
}
