package domain

// DNSResolution contains information about a single DNS resolution entry.
type DNSResolution struct {
	IP              string `json:"ip"`
	IsPrivate       bool   `json:"is_private"`
	IsGlobalUnicast bool   `json:"is_global_unicast"`
	IsLoopback      bool   `json:"is_loopback"`
	IsUnspecified   bool   `json:"is_unspecified"`
}

// DNSResolutionReport holds the resolution details for a given hostname.
type DNSResolutionReport struct {
	Hostname string          `json:"hostname"`
	IPs      []DNSResolution `json:"ips"`
}
