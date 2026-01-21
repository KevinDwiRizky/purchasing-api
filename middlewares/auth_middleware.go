package middlewares

import (
	"purchasing-api/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return utils.ResponseError(c, 401, "Token tidak ditemukan")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return utils.ResponseError(c, 401, "Format token tidak valid")
	}

	tokenString := parts[1]
	userID, err := utils.VerifyToken(tokenString)
	if err != nil {
		return utils.ResponseError(c, 401, "Token tidak valid atau sudah expired")
	}

	c.Locals("user_id", userID)
	return c.Next()
}
