package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUsername         string
	DBPassword         string
	DBName             string
	DBHost             string
	DBPort             string
	AgentID            string
	GoDaddyAPIKey      string
	GoogleDomainAPIKey string
	NameCheapAPIKey    string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config := &Config{
		DBUsername:         os.Getenv("DB_USER"),
		DBPassword:         os.Getenv("DB_PASSWORD"),
		DBName:             os.Getenv("DB_NAME"),
		DBHost:             os.Getenv("DB_HOST"),
		DBPort:             os.Getenv("DB_PORT"),
		AgentID:            os.Getenv("AGENT_ID"),
		GoDaddyAPIKey:      os.Getenv("GO_DADDY_API_KEY"),
		GoogleDomainAPIKey: os.Getenv("GOOGLE_DOMAIN_API_KEY"),
		NameCheapAPIKey:    os.Getenv("NAMECHEAP_API_KEY"),
	}

	return config
}
