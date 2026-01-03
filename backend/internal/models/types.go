package models

// Database represents a MySQL database
type Database struct {
	Name string `json:"name"`
}

// Table represents a MySQL table
type Table struct {
	Name    string `json:"name"`
	Engine  string `json:"engine"`
	Rows    int64  `json:"rows"`
	Size    string `json:"size"`
	Comment string `json:"comment"`
}

// Column represents a table column
type Column struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	Null         string `json:"null"`
	Key          string `json:"key"`
	Default      *string `json:"default"`
	Extra        string `json:"extra"`
	Comment      string `json:"comment"`
}

// TableStructure represents the structure of a table
type TableStructure struct {
	Columns []Column `json:"columns"`
	Indexes []Index  `json:"indexes"`
}

// Index represents a table index
type Index struct {
	Name    string   `json:"name"`
	Columns []string `json:"columns"`
	Unique  bool     `json:"unique"`
}

// QueryRequest represents a SQL query request
type QueryRequest struct {
	Database string `json:"database"`
	Query    string `json:"query"`
}

// QueryResponse represents a SQL query response
type QueryResponse struct {
	Columns []string        `json:"columns"`
	Rows    [][]interface{} `json:"rows"`
	Affected int64          `json:"affected"`
	Error   string          `json:"error,omitempty"`
}

// TableDataRequest represents a request for table data
type TableDataRequest struct {
	Database string `json:"database"`
	Table    string `json:"table"`
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
	SortBy   string `json:"sortBy,omitempty"`
	SortDir  string `json:"sortDir,omitempty"`
}

// TableDataResponse represents paginated table data
type TableDataResponse struct {
	Columns   []string        `json:"columns"`
	Rows      [][]interface{} `json:"rows"`
	Total     int64           `json:"total"`
	Page      int             `json:"page"`
	PageSize  int             `json:"pageSize"`
}

// LoginRequest represents a login request
type LoginRequest struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

// LoginResponse represents a login response
type LoginResponse struct {
	Token string `json:"token"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}

