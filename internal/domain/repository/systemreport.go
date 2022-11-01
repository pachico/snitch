package repository

import "pachico/snitch/internal/domain"

type SystemReportRepositoryInterface interface {
	GetSystemReport() (domain.SystemReport, error)
}
