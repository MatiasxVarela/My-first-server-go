package main

import (
	"myFirstServerWithGo/app"
)

func main() {
	server := app.CreateApp()

	server.Listen(":3000")
}
