package mysql

import (
	"database/sql"
	"fmt"
	"main/core/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	// Conn is a global variable that holds the database connection.
	Conn *gorm.DB

	// DB is a global variable that holds the database connection.
	DB *sql.DB
)

// Connect is a helper function that connects to the database.
//
// Returns an error if there is an issue connecting to the database.
func Connect() error {
	var err error

	// Create a dsn string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBConfig.Username, config.DBConfig.Password, config.DBConfig.Hostname, config.DBConfig.Port, config.DBConfig.DatabaseName)

	// Open a connection to the databasea
	Conn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// Check if there is an error connecting to the database
	if err != nil {
		return err
	}

	// Get the database instance
	DB, err = Conn.DB()

	return err
}
