package middleware

import (
	"mysql-admin-tool/internal/services"

	"github.com/gofiber/fiber/v2"
)

// AuthMiddleware validates JWT token and stores MySQL credentials in context
func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(401).JSON(fiber.Map{
				"error": "Authorization header required",
			})
		}

		// Extract token from "Bearer <token>"
		tokenString := authHeader
		if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			tokenString = authHeader[7:]
		}

		claims, err := services.ValidateToken(tokenString)
		if err != nil {
			return c.Status(401).JSON(fiber.Map{
				"error": "Invalid or expired token",
			})
		}

		// Store MySQL credentials in context
		c.Locals("db_host", claims.DBHost)
		c.Locals("db_port", claims.DBPort)
		c.Locals("db_user", claims.DBUser)
		c.Locals("db_password", claims.DBPassword)
		c.Locals("db_database", claims.DBDatabase)
		return c.Next()
	}
}

