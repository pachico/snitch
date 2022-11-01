package domain

import (
	"encoding/json"
	"testing"
)

func TestDNSResolutionSerialization(t *testing.T) {
	dns := DNSResolution{
		IP:              "192.168.1.1",
		IsPrivate:       true,
		IsGlobalUnicast: false,
		IsLoopback:      false,
		IsUnspecified:   false,
	}
	bytes, err := json.Marshal(dns)
	if err != nil {
		t.Fatalf("Failed to serialize DNSResolution: %v", err)
	}
	jsonStr := string(bytes)
	expected := `{"ip":"192.168.1.1","is_private":true,"is_global_unicast":false,"is_loopback":false,"is_unspecified":false}`
	if jsonStr != expected {
		t.Errorf("Expected JSON '%s', got '%s'", expected, jsonStr)
	}
}

func TestDNSResolutionReportSerialization(t *testing.T) {
	report := DNSResolutionReport{
		Hostname: "example.com",
		IPs: []DNSResolution{
			{
				IP:              "192.168.1.1",
				IsPrivate:       true,
				IsGlobalUnicast: false,
				IsLoopback:      false,
				IsUnspecified:   false,
			},
		},
	}
	bytes, err := json.Marshal(report)
	if err != nil {
		t.Fatalf("Failed to serialize DNSResolutionReport: %v", err)
	}
	jsonStr := string(bytes)
	expected := `{"hostname":"example.com","ips":[{"ip":"192.168.1.1","is_private":true,"is_global_unicast":false,"is_loopback":false,"is_unspecified":false}]}`
	if jsonStr != expected {
		t.Errorf("Expected JSON '%s', got '%s'", expected, jsonStr)
	}
}
