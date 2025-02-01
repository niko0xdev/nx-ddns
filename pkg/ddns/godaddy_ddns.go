package ddns

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// DNSRecord represents the data structure for a DNS record update
type DNSRecord struct {
	Data string `json:"data"` // IP address (or any other data you wish to update)
	TTL  int    `json:"ttl"`  // Time To Live (TTL) in seconds
	Type string `json:"type"` // Type of record (e.g., A, CNAME, etc.)
	Name string `json:"name"` // Name of the record
}

// api key should be: <key>:<seceret>
func UpdateGoDaddyDNSRecord(apiKey, domain, recordName, ip string) error {
	url := fmt.Sprintf("%s/%s/records/A/%s", GODADDY_DNS_API, domain, recordName)

	// Prepare the DNS record JSON data to be updated
	record := DNSRecord{
		Data: ip,  // IP address to update
		TTL:  600, // TTL value,
	}

	// Marshal the struct into JSON
	recordData, err := json.Marshal([]DNSRecord{record}) // Wrap in an array as the API expects a list of records
	if err != nil {
		return fmt.Errorf("failed to marshal record data: %v", err)
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(recordData))
	if err != nil {
		return err
	}

	fmt.Printf("[] Update GoDaddy DNS for URL: %s\n...", url)
	fmt.Printf("[] Update GoDaddy DNS for RecordName: %s\n...", recordName)
	fmt.Printf("[] Update GoDaddy DNS for ApiKey: %s\n...", apiKey)

	req.Header.Set("Authorization", "sso-key "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to update DNS record: %v", resp.Status)
	}

	return nil
}
