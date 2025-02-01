package repository

import (
	"fmt"
	"time"

	"github.com/niko0xdev/nx-ddns/internal/database"
	"github.com/niko0xdev/nx-ddns/internal/utils"

	"gorm.io/gorm"
)

// DNSRecordRepository defines the methods for interacting with DNS records
type DNSRecordRepository struct {
	DB *gorm.DB
}

// NewDNSRecordRepository creates a new instance of DNSRecordRepository
func NewDNSRecordRepository(db *gorm.DB) *DNSRecordRepository {
	return &DNSRecordRepository{DB: db}
}

// CreateDNSRecord inserts a new DNS record into the database
func (repo *DNSRecordRepository) CreateDNSRecord(dnsRecord *database.DNSRecord) (*database.DNSRecord, error) {
	if err := repo.DB.Create(dnsRecord).Error; err != nil {
		return nil, fmt.Errorf("error creating DNS record: %w", err)
	}
	return dnsRecord, nil
}

// GetDNSRecords fetches all DNS records
func (repo *DNSRecordRepository) GetDNSRecords() ([]database.DNSRecord, error) {
	var dnsRecords []database.DNSRecord
	if err := repo.DB.Find(&dnsRecords).Error; err != nil {
		return nil, fmt.Errorf("error fetching DNS records: %w", err)
	}

	return dnsRecords, nil
}

// GetDNSRecordByID fetches a DNS record by its ID
func (repo *DNSRecordRepository) GetDNSRecordByID(id string) (*database.DNSRecord, error) {
	var dnsRecord database.DNSRecord
	if err := repo.DB.First(&dnsRecord, id).Error; err != nil {
		return nil, fmt.Errorf("error fetching DNS record: %w", err)
	}
	return &dnsRecord, nil
}

// UpdateDNSRecord updates an existing DNS record
func (repo *DNSRecordRepository) UpdateDNSRecord(id string, dnsRecord *database.DNSRecord) (*database.DNSRecord, error) {
	// get record first
	record, err := repo.GetDNSRecordByID(id)
	if err != nil {
		return nil, err
	}

	// add new DNS log
	dnsLog := &database.DNSLog{
		BaseModel: database.BaseModel{
			ID:        utils.GenerateID().String(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		DNSRecordID: record.ID,
		OldValue:    record.IPAddress,
		NewValue:    dnsRecord.IPAddress,
	}

	if err := repo.DB.Save(record).Error; err != nil {
		return nil, fmt.Errorf("error updating DNS record: %w", err)
	}

	if err := repo.DB.Save(dnsLog).Error; err != nil {
		return nil, fmt.Errorf("error add DNS record change log: %w", err)
	}

	return dnsRecord, nil
}

// DeleteDNSRecord deletes a DNS record by its ID
func (repo *DNSRecordRepository) DeleteDNSRecord(id string) error {
	if err := repo.DB.Delete(&database.DNSRecord{}, id).Error; err != nil {
		return fmt.Errorf("error deleting DNS record: %w", err)
	}
	return nil
}

// CreateDNSLog creates a DNS change log entry
func (repo *DNSRecordRepository) CreateDNSLog(dnsLog *database.DNSLog) (*database.DNSLog, error) {
	if err := repo.DB.Create(dnsLog).Error; err != nil {
		return nil, fmt.Errorf("error creating DNS log: %w", err)
	}
	return dnsLog, nil
}

// GetDNSLogs fetches logs related to DNS record changes
func (repo *DNSRecordRepository) GetDNSLogs(dnsRecordID string) ([]database.DNSLog, error) {
	var logs []database.DNSLog
	if err := repo.DB.Where("dns_record_id = ?", dnsRecordID).Find(&logs).Error; err != nil {
		return nil, fmt.Errorf("error fetching DNS logs: %w", err)
	}
	return logs, nil
}
