package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/niko0xdev/nx-ddns/internal/config"
	"github.com/niko0xdev/nx-ddns/internal/database"
)

func main() {
	godotenv.Load()

	cfg := config.LoadConfig()

	database.InitDBConnection(cfg)

	err := database.Ping()
	if err != nil {
		panic("Failed to connect to the database")
	}

	if err := database.DB.AutoMigrate(&database.DNSRecord{}, &database.DNSLog{}); err != nil {
		panic("Failed to migrate database")
	}

	log.Println("Database migrated successfully")
}
