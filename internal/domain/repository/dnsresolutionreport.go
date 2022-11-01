package repository

import "pachico/snitch/internal/domain"

type DNSResolutionReportRepositoryInterface interface {
	GetDNSResolutionReport(host string) (domain.DNSResolutionReport, error)
}
