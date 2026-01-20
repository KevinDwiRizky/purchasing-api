package main

import (
	"log"
	"os"
	"purchasing-api/config"
	"purchasing-api/models"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	app := fiber.New()

	config.ConnectDB()

	config.DB.AutoMigrate(
		&models.User{},
		&models.Supplier{},
		&models.Item{},
		&models.Purchasing{},
		&models.PurchasingDetail{},
	)

	
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("API + DB connected")
	})

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	app.Listen(":" + port)
}
