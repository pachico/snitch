package repository

import "pachico/snitch/internal/domain"

type FSReportRepositoryInterface interface {
	GetFSReport() (domain.FSReport, error)
}
