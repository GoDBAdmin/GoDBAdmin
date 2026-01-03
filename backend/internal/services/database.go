package services

import (
	"database/sql"
	"fmt"
	"mysql-admin-tool/internal/database"
	"mysql-admin-tool/internal/models"
)

// GetDatabases returns a list of all databases
func GetDatabases(host string, port int, user, password, defaultDB string) ([]models.Database, error) {
	dbConn, err := database.ConnectToDatabase(defaultDB, host, port, user, password)
	if err != nil {
		return nil, err
	}
	defer dbConn.Close()

	rows, err := dbConn.Query("SHOW DATABASES")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var databases []models.Database
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			continue
		}
		// Skip system databases
		if name == "information_schema" || name == "performance_schema" || name == "sys" {
			continue
		}
		databases = append(databases, models.Database{Name: name})
	}

	return databases, nil
}

// GetTables returns a list of tables in a database
func GetTables(dbName string, host string, port int, user, password string) ([]models.Table, error) {
	dbConn, err := database.ConnectToDatabase(dbName, host, port, user, password)
	if err != nil {
		return nil, err
	}
	defer dbConn.Close()

	query := fmt.Sprintf("SELECT TABLE_NAME, ENGINE, TABLE_ROWS, ROUND((DATA_LENGTH + INDEX_LENGTH) / 1024 / 1024, 2) AS SIZE_MB, TABLE_COMMENT FROM information_schema.TABLES WHERE TABLE_SCHEMA = '%s' ORDER BY TABLE_NAME", dbName)
	rows, err := dbConn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tables []models.Table
	for rows.Next() {
		var table models.Table
		var sizeMB sql.NullFloat64
		if err := rows.Scan(&table.Name, &table.Engine, &table.Rows, &sizeMB, &table.Comment); err != nil {
			continue
		}
		if sizeMB.Valid {
			table.Size = fmt.Sprintf("%.2f MB", sizeMB.Float64)
		} else {
			table.Size = "0 MB"
		}
		tables = append(tables, table)
	}

	return tables, nil
}

// GetTableStructure returns the structure of a table
func GetTableStructure(dbName, tableName string, host string, port int, user, password string) (*models.TableStructure, error) {
	dbConn, err := database.ConnectToDatabase(dbName, host, port, user, password)
	if err != nil {
		return nil, err
	}
	defer dbConn.Close()

	// Get columns
	query := fmt.Sprintf("SELECT COLUMN_NAME, COLUMN_TYPE, IS_NULLABLE, COLUMN_KEY, COLUMN_DEFAULT, EXTRA, COLUMN_COMMENT FROM information_schema.COLUMNS WHERE TABLE_SCHEMA = '%s' AND TABLE_NAME = '%s' ORDER BY ORDINAL_POSITION", dbName, tableName)
	rows, err := dbConn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var columns []models.Column
	for rows.Next() {
		var col models.Column
		var defaultVal sql.NullString
		if err := rows.Scan(&col.Name, &col.Type, &col.Null, &col.Key, &defaultVal, &col.Extra, &col.Comment); err != nil {
			continue
		}
		if defaultVal.Valid {
			col.Default = &defaultVal.String
		}
		columns = append(columns, col)
	}
	rows.Close()

	// Get indexes
	indexQuery := fmt.Sprintf("SELECT DISTINCT INDEX_NAME, NON_UNIQUE FROM information_schema.STATISTICS WHERE TABLE_SCHEMA = '%s' AND TABLE_NAME = '%s' ORDER BY INDEX_NAME", dbName, tableName)
	indexRows, err := dbConn.Query(indexQuery)
	if err != nil {
		return nil, err
	}
	defer indexRows.Close()

	indexMap := make(map[string]*models.Index)
	for indexRows.Next() {
		var indexName string
		var nonUnique int
		if err := indexRows.Scan(&indexName, &nonUnique); err != nil {
			continue
		}
		if indexName == "PRIMARY" || indexName != "" {
			indexMap[indexName] = &models.Index{
				Name:   indexName,
				Unique: nonUnique == 0,
			}
		}
	}
	indexRows.Close()

	// Get index columns
	for indexName := range indexMap {
		colQuery := fmt.Sprintf("SELECT COLUMN_NAME FROM information_schema.STATISTICS WHERE TABLE_SCHEMA = '%s' AND TABLE_NAME = '%s' AND INDEX_NAME = '%s' ORDER BY SEQ_IN_INDEX", dbName, tableName, indexName)
		colRows, err := dbConn.Query(colQuery)
		if err != nil {
			continue
		}
		for colRows.Next() {
			var colName string
			if err := colRows.Scan(&colName); err != nil {
				continue
			}
			indexMap[indexName].Columns = append(indexMap[indexName].Columns, colName)
		}
		colRows.Close()
	}

	var indexes []models.Index
	for _, idx := range indexMap {
		indexes = append(indexes, *idx)
	}

	return &models.TableStructure{
		Columns: columns,
		Indexes: indexes,
	}, nil
}

// GetTableData returns paginated table data
func GetTableData(dbName, tableName string, page, pageSize int, sortBy, sortDir string, host string, port int, user, password string) (*models.TableDataResponse, error) {
	dbConn, err := database.ConnectToDatabase(dbName, host, port, user, password)
	if err != nil {
		return nil, err
	}
	defer dbConn.Close()

	// Get total count
	var total int64
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM `%s`.`%s`", dbName, tableName)
	err = dbConn.QueryRow(countQuery).Scan(&total)
	if err != nil {
		return nil, err
	}

	// Build query
	offset := (page - 1) * pageSize
	query := fmt.Sprintf("SELECT * FROM `%s`.`%s`", dbName, tableName)
	
	if sortBy != "" {
		query += fmt.Sprintf(" ORDER BY `%s`", sortBy)
		if sortDir == "desc" || sortDir == "DESC" {
			query += " DESC"
		} else {
			query += " ASC"
		}
	}
	
	query += fmt.Sprintf(" LIMIT %d OFFSET %d", pageSize, offset)

	rows, err := dbConn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var resultRows [][]interface{}
	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range values {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			continue
		}

		// Convert []byte to string for JSON serialization
		row := make([]interface{}, len(values))
		for i, v := range values {
			if b, ok := v.([]byte); ok {
				row[i] = string(b)
			} else {
				row[i] = v
			}
		}
		resultRows = append(resultRows, row)
	}

	return &models.TableDataResponse{
		Columns:  columns,
		Rows:     resultRows,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}, nil
}

