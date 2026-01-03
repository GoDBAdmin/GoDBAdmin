package services

import (
	"database/sql"
	"fmt"
	"mysql-admin-tool/internal/database"
	"mysql-admin-tool/internal/models"
	"strings"
)

// ExecuteQuery executes a SQL query and returns results
func ExecuteQuery(dbName, query string, host string, port int, user, password string) (*models.QueryResponse, error) {
	dbConn, err := database.ConnectToDatabase(dbName, host, port, user, password)
	if err != nil {
		return nil, err
	}
	defer dbConn.Close()

	// Trim and check query type
	query = strings.TrimSpace(query)
	upperQuery := strings.ToUpper(query)

	// Handle SELECT queries
	if strings.HasPrefix(upperQuery, "SELECT") || strings.HasPrefix(upperQuery, "SHOW") || strings.HasPrefix(upperQuery, "DESCRIBE") || strings.HasPrefix(upperQuery, "DESC") || strings.HasPrefix(upperQuery, "EXPLAIN") {
		return executeSelectQuery(dbConn, query)
	}

	// Handle INSERT, UPDATE, DELETE, CREATE, DROP, ALTER
	result, err := dbConn.Exec(query)
	if err != nil {
		return &models.QueryResponse{
			Error: err.Error(),
		}, nil
	}

	affected, _ := result.RowsAffected()
	return &models.QueryResponse{
		Affected: affected,
	}, nil
}

// executeSelectQuery executes a SELECT query and returns rows
func executeSelectQuery(dbConn *sql.DB, query string) (*models.QueryResponse, error) {
	rows, err := dbConn.Query(query)
	if err != nil {
		return &models.QueryResponse{
			Error: err.Error(),
		}, nil
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return &models.QueryResponse{
			Error: err.Error(),
		}, nil
	}

	var resultRows [][]interface{}
	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range values {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			return &models.QueryResponse{
				Error: err.Error(),
			}, nil
		}

		// Convert []byte to string for JSON serialization
		row := make([]interface{}, len(values))
		for i, v := range values {
			if b, ok := v.([]byte); ok {
				row[i] = string(b)
			} else if v == nil {
				row[i] = nil
			} else {
				row[i] = v
			}
		}
		resultRows = append(resultRows, row)
	}

	return &models.QueryResponse{
		Columns: columns,
		Rows:    resultRows,
	}, nil
}

// CreateTable creates a new table
func CreateTable(dbName, tableName, createSQL string, host string, port int, user, password string) error {
	dbConn, err := database.ConnectToDatabase(dbName, host, port, user, password)
	if err != nil {
		return err
	}
	defer dbConn.Close()

	_, err = dbConn.Exec(createSQL)
	return err
}

// DropTable drops a table
func DropTable(dbName, tableName string, host string, port int, user, password string) error {
	dbConn, err := database.ConnectToDatabase(dbName, host, port, user, password)
	if err != nil {
		return err
	}
	defer dbConn.Close()

	query := fmt.Sprintf("DROP TABLE IF EXISTS `%s`.`%s`", dbName, tableName)
	_, err = dbConn.Exec(query)
	return err
}

