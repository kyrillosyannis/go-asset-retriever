package database

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB is a global variable to hold the database connection
var DB *gorm.DB

// ConnectDatabase initializes the database connection
func ConnectDatabase() (*gorm.DB, error) {
	// Default values if environment variables are not set
	host := "localhost"
	port := "5434"
	user := "postgres"
	password := "postgres"
	dbname := "asset_db"
	sslmode := "disable"

	// PostgreSQL DSN (modify for other databases)
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)

	// Open the connection with GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Enable SQL query logging
	})
	fmt.Println("db is:")
	fmt.Println(db)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Set database connection pool settings
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(25)            // Max open connections
	sqlDB.SetMaxIdleConns(10)            // Max idle connections
	sqlDB.SetConnMaxLifetime(5 * time.Minute) // Connection max lifetime

	DB = db
	fmt.Println("✅ Successfully connected to the database!")
	return db, nil
}
