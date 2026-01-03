package database

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db     *sql.DB
	dbOnce sync.Once
)

// Config holds database configuration
type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	MaxConns int
	MaxIdle  int
}

// Init initializes the database connection pool
func Init(config Config) error {
	var err error
	
	dbOnce.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&charset=utf8mb4",
			config.User,
			config.Password,
			config.Host,
			config.Port,
			config.Database,
		)

		db, err = sql.Open("mysql", dsn)
		if err != nil {
			return
		}

		// Set connection pool settings
		if config.MaxConns > 0 {
			db.SetMaxOpenConns(config.MaxConns)
		} else {
			db.SetMaxOpenConns(25)
		}

		if config.MaxIdle > 0 {
			db.SetMaxIdleConns(config.MaxIdle)
		} else {
			db.SetMaxIdleConns(5)
		}

		db.SetConnMaxLifetime(5 * time.Minute)

		// Test connection
		err = db.Ping()
	})

	return err
}

// GetDB returns the database connection
func GetDB() *sql.DB {
	return db
}

// Close closes the database connection
func Close() error {
	if db != nil {
		return db.Close()
	}
	return nil
}

// ConnectToDatabase connects to a specific database using credentials from context
func ConnectToDatabase(dbName string, host string, port int, user, password string) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&charset=utf8mb4",
		user,
		password,
		host,
		port,
		dbName,
	)

	return sql.Open("mysql", dsn)
}

var dbConfig *Config

// SetConfig stores the database configuration
func SetConfig(config Config) {
	dbConfig = &config
}

// GetConfig returns the stored database configuration
func GetConfig() Config {
	if dbConfig == nil {
		return Config{
			Host:     "localhost",
			Port:     3306,
			User:     "root",
			Password: "",
			Database: "mysql",
			MaxConns: 25,
			MaxIdle:  5,
		}
	}
	return *dbConfig
}

