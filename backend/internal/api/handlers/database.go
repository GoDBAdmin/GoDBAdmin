package handlers

import (
	"mysql-admin-tool/internal/models"
	"mysql-admin-tool/internal/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// GetDatabases returns all databases
func GetDatabases(c *fiber.Ctx) error {
	host := c.Locals("db_host").(string)
	port := c.Locals("db_port").(int)
	user := c.Locals("db_user").(string)
	password := c.Locals("db_password").(string)
	database := c.Locals("db_database").(string)

	databases, err := services.GetDatabases(host, port, user, password, database)
	if err != nil {
		return c.Status(500).JSON(models.ErrorResponse{
			Error: err.Error(),
		})
	}
	return c.JSON(databases)
}

// GetTables returns all tables in a database
func GetTables(c *fiber.Ctx) error {
	dbName := c.Params("db")
	if dbName == "" {
		return c.Status(400).JSON(models.ErrorResponse{
			Error: "Database name required",
		})
	}

	host := c.Locals("db_host").(string)
	port := c.Locals("db_port").(int)
	user := c.Locals("db_user").(string)
	password := c.Locals("db_password").(string)

	tables, err := services.GetTables(dbName, host, port, user, password)
	if err != nil {
		return c.Status(500).JSON(models.ErrorResponse{
			Error: err.Error(),
		})
	}
	return c.JSON(tables)
}

// GetTableStructure returns the structure of a table
func GetTableStructure(c *fiber.Ctx) error {
	dbName := c.Params("db")
	tableName := c.Params("table")

	if dbName == "" || tableName == "" {
		return c.Status(400).JSON(models.ErrorResponse{
			Error: "Database and table name required",
		})
	}

	host := c.Locals("db_host").(string)
	port := c.Locals("db_port").(int)
	user := c.Locals("db_user").(string)
	password := c.Locals("db_password").(string)

	structure, err := services.GetTableStructure(dbName, tableName, host, port, user, password)
	if err != nil {
		return c.Status(500).JSON(models.ErrorResponse{
			Error: err.Error(),
		})
	}
	return c.JSON(structure)
}

// GetTableData returns paginated table data
func GetTableData(c *fiber.Ctx) error {
	dbName := c.Params("db")
	tableName := c.Params("table")

	if dbName == "" || tableName == "" {
		return c.Status(400).JSON(models.ErrorResponse{
			Error: "Database and table name required",
		})
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize", "50"))
	sortBy := c.Query("sortBy", "")
	sortDir := c.Query("sortDir", "asc")

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 1000 {
		pageSize = 50
	}

	host := c.Locals("db_host").(string)
	port := c.Locals("db_port").(int)
	user := c.Locals("db_user").(string)
	password := c.Locals("db_password").(string)

	data, err := services.GetTableData(dbName, tableName, page, pageSize, sortBy, sortDir, host, port, user, password)
	if err != nil {
		return c.Status(500).JSON(models.ErrorResponse{
			Error: err.Error(),
		})
	}
	return c.JSON(data)
}

// ExecuteQuery executes a SQL query
func ExecuteQuery(c *fiber.Ctx) error {
	var req models.QueryRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(models.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	if req.Query == "" {
		return c.Status(400).JSON(models.ErrorResponse{
			Error: "Query is required",
		})
	}

	if req.Database == "" {
		return c.Status(400).JSON(models.ErrorResponse{
			Error: "Database name is required",
		})
	}

	host := c.Locals("db_host").(string)
	port := c.Locals("db_port").(int)
	user := c.Locals("db_user").(string)
	password := c.Locals("db_password").(string)

	result, err := services.ExecuteQuery(req.Database, req.Query, host, port, user, password)
	if err != nil {
		return c.Status(500).JSON(models.ErrorResponse{
			Error: err.Error(),
		})
	}
	return c.JSON(result)
}

// CreateTable creates a new table
func CreateTable(c *fiber.Ctx) error {
	dbName := c.Params("db")
	if dbName == "" {
		return c.Status(400).JSON(models.ErrorResponse{
			Error: "Database name required",
		})
	}

	var req struct {
		TableName string `json:"tableName"`
		CreateSQL string `json:"createSQL"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(models.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	host := c.Locals("db_host").(string)
	port := c.Locals("db_port").(int)
	user := c.Locals("db_user").(string)
	password := c.Locals("db_password").(string)

	if err := services.CreateTable(dbName, req.TableName, req.CreateSQL, host, port, user, password); err != nil {
		return c.Status(500).JSON(models.ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Table created successfully",
	})
}

// DropTable drops a table
func DropTable(c *fiber.Ctx) error {
	dbName := c.Params("db")
	tableName := c.Params("table")

	if dbName == "" || tableName == "" {
		return c.Status(400).JSON(models.ErrorResponse{
			Error: "Database and table name required",
		})
	}

	host := c.Locals("db_host").(string)
	port := c.Locals("db_port").(int)
	user := c.Locals("db_user").(string)
	password := c.Locals("db_password").(string)

	if err := services.DropTable(dbName, tableName, host, port, user, password); err != nil {
		return c.Status(500).JSON(models.ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Table dropped successfully",
	})
}

