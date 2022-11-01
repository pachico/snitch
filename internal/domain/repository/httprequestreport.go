package repository

import "pachico/snitch/internal/domain"

type HTTPRequestReportRepositoryInterface interface {
	GetHTTPRequestReport(hostname string) (domain.HTTPRequestReport, error)
}
