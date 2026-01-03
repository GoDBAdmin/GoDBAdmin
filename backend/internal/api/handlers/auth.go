package handlers

import (
	"mysql-admin-tool/internal/models"
	"mysql-admin-tool/internal/services"

	"github.com/gofiber/fiber/v2"
)

// Login handles user login with MySQL credentials
func Login(c *fiber.Ctx) error {
	var req models.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(models.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	// Set defaults
	if req.Host == "" {
		req.Host = "localhost"
	}
	if req.Port == 0 {
		req.Port = 3306
	}
	if req.Database == "" {
		req.Database = "mysql"
	}

	// Validate MySQL connection
	if err := services.ValidateMySQLConnection(req.Host, req.Port, req.Username, req.Password, req.Database); err != nil {
		return c.Status(401).JSON(models.ErrorResponse{
			Error: "Invalid MySQL credentials: " + err.Error(),
		})
	}

	// Generate token with MySQL credentials
	token, err := services.GenerateToken(req.Host, req.Port, req.Username, req.Password, req.Database)
	if err != nil {
		return c.Status(500).JSON(models.ErrorResponse{
			Error: "Failed to generate token",
		})
	}

	return c.JSON(models.LoginResponse{
		Token: token,
	})
}

