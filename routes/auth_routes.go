package routes

import (
	"purchasing-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	auth := app.Group("/api/v1/auth")
	auth.Post("/register", controllers.Register)
	auth.Post("/login", controllers.Login)
}
