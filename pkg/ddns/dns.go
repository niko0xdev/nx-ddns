package ddns

import "time"

const (
	NAMECHEAP_DNS_API_HOST = "https://dynamicdns.park-your-domain.com/update"
	GOOGLE_DNS_API_HOST    = "domains.google.com/nic/update"
	GODADDY_DNS_API        = "https://api.godaddy.com/v1/domains"
	httpTimeout            = 30 * time.Second
)
