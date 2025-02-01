package database

// DNSProvider type defines the valid DNS providers
type DNSProvider string

const (
	NameCheap DNSProvider = "NameCheap"
	GoDaddy   DNSProvider = "GoDaddy"
	Google    DNSProvider = "Google"
)

type DNSRecord struct {
	BaseModel
	DNSProvider DNSProvider `gorm:"type:varchar(50);not null" json:"dnsProvider"`
	Domain      string      `gorm:"size:255;not null" json:"domain"`
	RecordType  string      `gorm:"size:50;not null" json:"recordType"`
	RecordName  string      `gorm:"size:255;not null" json:"recordName"`
	IPAddress   string      `gorm:"size:45;not null" json:"ipAddress"`
}

type DNSLog struct {
	BaseModel
	DNSRecordID string `gorm:"type:varchar(36);not null" json:"dnsRecordId"`
	OldValue    string `gorm:"type:varchar(255)" json:"oldValue"`
	NewValue    string `gorm:"type:varchar(255)" json:"newValue"`
}
