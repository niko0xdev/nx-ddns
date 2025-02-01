package ddns

import (
	"fmt"
	"net/http"
)

// api key should be: <username>:<password>
func UpdateGoogleDNSRecord(apiKey, hostname, ip string) error {
	url := fmt.Sprintf("https://%s@%s?hostname=%s&myip=%s",
		apiKey, GOOGLE_DNS_API_HOST, hostname, ip)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	return nil
}
