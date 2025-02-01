package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/niko0xdev/nx-ddns/internal/config"
)

var DB *gorm.DB

// InitDBConnection initializes the database connection using PostgreSQL and the provided config.
func InitDBConnection(cfg *config.Config) {
	// Use the correct DSN format for PostgreSQL.
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUsername, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	log.Println("Database connection successfully established")
}

// Ping checks the database connection and returns any error encountered.
func Ping() error {
	// Get the underlying *sql.DB instance from GORM
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}

	// Ping the database to verify the connection
	return sqlDB.Ping()
}
