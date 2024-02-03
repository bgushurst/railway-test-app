package main

import (
	"os"

	"github.com/bgushurst/railway-test-app/handler"
	"github.com/gofiber/fiber/v2"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}

	return port
}

func main() {
	app := fiber.New()

	indexHandler := handler.IndexHandler{}

	app.Get("/", indexHandler.HandleIndexShow)

	app.Listen(getPort())
}
