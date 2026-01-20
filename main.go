package main

import (
	"log"
	"os"
	"purchasing-api/config"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	app := fiber.New()

	config.ConnectDB()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("API + DB connected")
	})

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	app.Listen(":" + port)
}
