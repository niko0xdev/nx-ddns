package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/niko0xdev/nx-ddns/internal/config"
	"github.com/niko0xdev/nx-ddns/internal/database"
	"github.com/niko0xdev/nx-ddns/internal/repository"
	"github.com/niko0xdev/nx-ddns/internal/utils"
	"github.com/niko0xdev/nx-ddns/pkg/ddns"
)

func main() {
	godotenv.Load()

	cfg := config.LoadConfig()

	database.InitDBConnection(cfg)

	err := database.Ping()
	if err != nil {
		panic("Failed to connect to the database")
	}

	// init db repository
	repo := repository.NewDNSRecordRepository(database.DB)

	// get active dns records
	records, err := repo.GetDNSRecords()
	if err != nil {
		panic(err)
	}

	// get current public ip
	ip, err := utils.GetPublicIP()
	if err != nil {
		panic(err)
	}

	// update dns records
	for _, record := range records {
		if record.IPAddress != ip {
			// update dns record
			_, err := ddns.UpdateDNSRecord(repo, &record, ip, cfg)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
