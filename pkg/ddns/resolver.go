package ddns

import (
	"fmt"

	"github.com/niko0xdev/nx-ddns/internal/config"
	"github.com/niko0xdev/nx-ddns/internal/database"
	"github.com/niko0xdev/nx-ddns/internal/repository"
)

func UpdateDNSRecord(repo *repository.DNSRecordRepository, record *database.DNSRecord, newIP string, cfg *config.Config) (*database.DNSRecord, error) {
	var updateErr error

	switch record.DNSProvider {
	case database.GoDaddy:
		updateErr = UpdateGoDaddyDNSRecord(
			cfg.GoDaddyAPIKey,
			record.Domain,
			record.RecordName,
			newIP,
		)
	case database.Google:
		updateErr = UpdateGoogleDNSRecord(
			cfg.GoogleDomainAPIKey,
			record.Domain,
			newIP,
		)
	case database.NameCheap:
		updateErr = UpdateNameCheapDNSRecord(
			cfg.NameCheapAPIKey,
			record.Domain,
			record.RecordName,
			newIP,
		)
	default:
		return nil, fmt.Errorf("unsupported DNS provider: %s", record.DNSProvider)
	}

	if updateErr != nil {
		return nil, updateErr
	}

	record.IPAddress = newIP
	if _, err := repo.UpdateDNSRecord(record.ID, record); err != nil {
		return nil, fmt.Errorf("failed to update DNS record in repository: %w", err)
	}

	fmt.Printf(" -> Updated DNS record for %s to %s\n", record.RecordName, record.IPAddress)

	return record, nil
}
