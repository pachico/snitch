package system

import (
	"errors"
	"net"
	"pachico/snitch/internal/domain"
)

type DNSResolutionReportRepository struct {
}

func (r *DNSResolutionReportRepository) GetDNSResolutionReport(hostname string) (domain.DNSResolutionReport, error) {

	report := domain.DNSResolutionReport{
		Hostname: hostname,
	}

	ips, err := net.LookupIP(hostname)
	if err != nil {
		return report, errors.New("error resolving " + hostname + ": " + err.Error())
	}

	for _, ip := range ips {
		dnsResolution := domain.DNSResolution{
			IP:              ip.String(),
			IsPrivate:       ip.IsPrivate(),
			IsGlobalUnicast: ip.IsGlobalUnicast(),
			IsLoopback:      ip.IsLoopback(),
			IsUnspecified:   ip.IsUnspecified(),
		}
		report.IPs = append(report.IPs, dnsResolution)
	}

	return report, nil
}
