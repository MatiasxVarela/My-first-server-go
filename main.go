package main

import (
	"myFirstServerWithGo/app"
	"myFirstServerWithGo/db"

	_ "github.com/lib/pq"
)

func main() {
	db := db.OpenDb(true)
	server := app.CreateApp(db)

	server.Listen(":3000")
	defer db.Close()
}
