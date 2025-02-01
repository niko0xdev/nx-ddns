package dto

import (
	"time"

	"github.com/niko0xdev/nx-ddns/internal/database"
	"github.com/niko0xdev/nx-ddns/internal/utils"
)

type DNSRecordRequest struct {
	DNSProvider database.DNSProvider `json:"dnsProvider"`
	Domain      string               `json:"domain"`
	RecordType  string               `json:"recordType"`
	RecordName  string               `json:"recordName"`
	IPAddress   string               `json:"ipAddress"`
}

type DNSRecord struct {
	ID          string `json:"id"`
	DNSProvider string `json:"dnsProvider"`
	Domain      string `json:"domain"`
	RecordType  string `json:"recordType"`
	RecordName  string `json:"recordName"`
	IPAddress   string `json:"ipAddress"`
	LastUpdated string `json:"lastUpdated"`
}

type DNSLog struct {
	ID        string `json:"id"`
	Timestamp string `json:"timestamp"`
	OldValue  string `json:"oldValue"`
	NewValue  string `json:"newValue"`
}

func DNSRecordToDTO(record database.DNSRecord) DNSRecord {
	return DNSRecord{
		ID:          record.ID,
		DNSProvider: string(record.DNSProvider),
		Domain:      record.Domain,
		RecordType:  record.RecordType,
		RecordName:  record.RecordName,
		IPAddress:   record.IPAddress,
		LastUpdated: record.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func DNSLogToDTO(log database.DNSLog) DNSLog {
	return DNSLog{
		ID:        log.ID,
		Timestamp: log.CreatedAt.Format("2006-01-02 15:04:05"),
		OldValue:  log.OldValue,
		NewValue:  log.NewValue,
	}
}

func DNSRecordFromDTO(record DNSRecordRequest) database.DNSRecord {
	return database.DNSRecord{
		BaseModel: database.BaseModel{
			ID:        utils.GenerateID().String(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		DNSProvider: database.DNSProvider(record.DNSProvider),
		Domain:      record.Domain,
		RecordType:  record.RecordType,
		RecordName:  record.RecordName,
		IPAddress:   record.IPAddress,
	}
}
