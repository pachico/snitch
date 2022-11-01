package domain

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcalusic/sysinfo"
)

func TestNewSystemReport(t *testing.T) {
	mockSysInfo := sysinfo.SysInfo{
		Meta: sysinfo.Meta{
			Version: "1.0",
		},
		Node:    sysinfo.Node{Hostname: "testnode"},
		OS:      sysinfo.OS{Name: "Linux"},
		Kernel:  sysinfo.Kernel{Version: "5.x"},
		Product: sysinfo.Product{Name: "Server"},
		Board:   sysinfo.Board{Name: "Mainboard"},
		Chassis: sysinfo.Chassis{Type: 123},
		BIOS:    sysinfo.BIOS{Vendor: "BIOSVendor"},
		CPU:     sysinfo.CPU{Model: "Intel"},
		Memory:  sysinfo.Memory{Type: "16GB"},
		Storage: []sysinfo.StorageDevice{{Name: "/dev/sda"}},
		Network: []sysinfo.NetworkDevice{{Name: "eth0"}},
	}

	report := NewSystemReport(mockSysInfo)

	reportJSON, _ := json.Marshal(report)

	expectedReportJSON := `{"sysinfo":{"version":"1.0","timestamp":"0001-01-01T00:00:00Z"},"node":{"hostname":"testnode"},"os":{"name":"Linux"},"kernel":{"version":"5.x"},"product":{"name":"Server","uuid":"00000000-0000-0000-0000-000000000000"},"board":{"name":"Mainboard"},"chassis":{"type":123},"bios":{"vendor":"BIOSVendor"},"cpu":{"model":"Intel"},"memory":{"type":"16GB"},"storage":[{"name":"/dev/sda"}],"network":[{"name":"eth0"}]}`

	assert.Equal(t, expectedReportJSON, string(reportJSON))

}
