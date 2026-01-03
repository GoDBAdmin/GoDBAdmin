package api

import (
	"mysql-admin-tool/internal/api/handlers"
	"mysql-admin-tool/internal/api/middleware"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes configures all API routes
func SetupRoutes(app *fiber.App) {
	// Public routes
	api := app.Group("/api")
	api.Post("/auth/login", handlers.Login)

	// Protected routes
	protected := api.Group("", middleware.AuthMiddleware())
	{
		// Database routes
		protected.Get("/databases", handlers.GetDatabases)
		protected.Get("/databases/:db/tables", handlers.GetTables)
		protected.Get("/databases/:db/tables/:table", handlers.GetTableStructure)
		protected.Get("/databases/:db/tables/:table/data", handlers.GetTableData)

		// Query routes
		protected.Post("/query", handlers.ExecuteQuery)

		// Table management routes
		protected.Post("/databases/:db/tables", handlers.CreateTable)
		protected.Delete("/databases/:db/tables/:table", handlers.DropTable)
	}
}

