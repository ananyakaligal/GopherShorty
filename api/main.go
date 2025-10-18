package main

import (
	"/routes"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
}

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	app := fiber.New()
	app.Use(logger.New()) // Logs incoming requests

	setupRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
