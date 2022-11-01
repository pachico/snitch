package system_test

import (
	"pachico/snitch/internal/infrastructure/repository/system"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDNSResolutionReportReturnsProperReport(t *testing.T) {
	// Arrange
	repo := system.DNSResolutionReportRepository{}
	hostname := "google.com"
	// Act
	report, err := repo.GetDNSResolutionReport(hostname)
	// Assert
	assert.Nil(t, err)
	assert.NotEmpty(t, report.IPs)
	assert.Equal(t, hostname, report.Hostname)
	assert.NotEmpty(t, report.IPs[0].IP)
}

func TestGetDNSResolutionReportReturnsErrorIFFailed(t *testing.T) {
	// Arrange
	repo := system.DNSResolutionReportRepository{}
	hostname := "_invalid_hostname_"
	// Act
	_, err := repo.GetDNSResolutionReport(hostname)
	// Assert
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "error resolving _invalid_hostname_:")
}
