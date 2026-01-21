package main

import (
	"log"
	"os"
	"purchasing-api/config"
	"purchasing-api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	app := fiber.New()

	config.ConnectDB()

	// config.DB.AutoMigrate(
	// 	&models.User{},
	// 	&models.Supplier{},
	// 	&models.Item{},
	// 	&models.Purchasing{},
	// 	&models.PurchasingDetail{},
	// )

	routes.SetupRoutes(app)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	app.Listen(":" + port)
}
