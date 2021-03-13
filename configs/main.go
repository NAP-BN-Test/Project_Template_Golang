package configs

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// Replace with your own connection parameters
var server = "localhost"
var port = 1433
var user = "sa"
var password = "1234"
var database = "REGISTER_STUDY_MANAGER"

func DBConn() (*gorm.DB, error) {
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", user, password, server, port, database)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	sqlDB, err := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err != nil {
		log.Fatal("Error creating connection pool: " + err.Error())
	}
	// Close the database connection pool after program executes
	// defer db.Close()
	return db, err
}
